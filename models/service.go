package models

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/convox/kernel/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/aws"
	"github.com/convox/kernel/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/service/cloudformation"
	"github.com/convox/kernel/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/service/dynamodb"
	"github.com/convox/kernel/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/service/s3"
)

type Service struct {
	Name     string
	Password string
	Type     string
	Status   string
	URL      string

	App string

	Stack string

	Outputs    map[string]string
	Parameters map[string]string
	Tags       map[string]string
}

type Services []Service

func ListServices(app string) (Services, error) {
	a, err := GetApp(app)

	if err != nil {
		if strings.Index(err.Error(), "does not exist") != -1 {
			return Services{}, nil
		}

		return nil, err
	}

	req := &s3.ListObjectsInput{
		Bucket: aws.String(a.Outputs["Settings"]),
		Prefix: aws.String("service/"),
	}

	res, err := S3().ListObjects(req)

	services := make(Services, len(res.Contents))
	servicesByName := map[string]Service{}

	for i, s := range res.Contents {
		name := strings.TrimPrefix(*s.Key, "service/")
		svc, err := GetService(app, name)

		if err != nil {
			return nil, err
		}

		services[i] = *svc
		servicesByName[name] = *svc
	}

	release, err := a.LatestRelease()

	if err != nil {
		return nil, err
	}

	if release != nil {
		rss, err := release.Services()

		if err != nil {
			return nil, err
		}

		for _, rs := range rss {
			if _, ok := servicesByName[rs.Name]; !ok {
				services = append(services, rs)
			}
		}
	}

	return services, nil
}

func ListServiceStacks() (Services, error) {
	res, err := CloudFormation().DescribeStacks(&cloudformation.DescribeStacksInput{})

	if err != nil {
		return nil, err
	}

	services := make(Services, 0)

	for _, stack := range res.Stacks {
		tags := stackTags(stack)

		if tags["System"] == "convox" && tags["Type"] == "service" {
			services = append(services, *serviceFromStack(stack))
		}
	}

	return services, nil
}

func GetService(app, name string) (*Service, error) {
	a, err := GetApp(app)

	if err != nil {
		return nil, err
	}

	value, err := s3Get(a.Outputs["Settings"], fmt.Sprintf("service/%s", name))

	if err != nil {
		return nil, err
	}

	var service *Service

	err = json.Unmarshal([]byte(value), &service)

	if err != nil {
		return nil, err
	}

	return service, nil
}

func GetServiceFromName(name string) (*Service, error) {
	res, err := CloudFormation().DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(name)})

	if err != nil {
		return nil, err
	}

	return serviceFromStack(res.Stacks[0]), nil
}

func (s *Service) Create() error {
	formation, err := s.Formation()

	if err != nil {
		return err
	}

	params := map[string]string{
		"Password": s.Password,
	}

	if s.Type == "redis" {
		params["SSHKey"] = ""
	}

	tags := map[string]string{
		"System":  "convox",
		"Type":    "service",
		"Service": s.Type,
	}

	req := &cloudformation.CreateStackInput{
		StackName:    aws.String(s.Name),
		TemplateBody: aws.String(formation),
	}

	for key, value := range params {
		req.Parameters = append(req.Parameters, &cloudformation.Parameter{ParameterKey: aws.String(key), ParameterValue: aws.String(value)})
	}

	for key, value := range tags {
		req.Tags = append(req.Tags, &cloudformation.Tag{Key: aws.String(key), Value: aws.String(value)})
	}

	_, err = CloudFormation().CreateStack(req)

	return err
}

func (s *Service) Formation() (string, error) {
	data, err := exec.Command("docker", "run", "convox/service", s.Type).CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (s *Service) Save() error {
	app, err := GetApp(s.App)

	if err != nil {
		return err
	}

	data, err := json.Marshal(s)

	if err != nil {
		return err
	}

	return s3Put(app.Outputs["Settings"], fmt.Sprintf("service/%s", s.Name), data, false)
}

func (s *Service) ManagementUrl() string {
	region := os.Getenv("AWS_REGION")

	resources, err := ListResources(s.App)

	if err != nil {
		panic(err)
	}

	switch s.Type {
	case "convox/postgres":
		id := resources[fmt.Sprintf("%sDatabase", upperName(s.Name))].Id
		return fmt.Sprintf("https://console.aws.amazon.com/rds/home?region=%s#dbinstances:id=%s;sf=all", region, id)
	case "convox/redis":
		id := resources[fmt.Sprintf("%sInstances", upperName(s.Name))].Id
		return fmt.Sprintf("https://console.aws.amazon.com/ec2/autoscaling/home?region=%s#AutoScalingGroups:id=%s;view=details", region, id)
	default:
		return ""
	}
}

func servicesTable(app string) string {
	return fmt.Sprintf("%s-services", app)
}

func serviceFromItem(item map[string]*dynamodb.AttributeValue) *Service {
	return &Service{
		Name: coalesce(item["name"], ""),
		Type: coalesce(item["type"], ""),
		App:  coalesce(item["app"], ""),
	}
}

func serviceFromStack(stack *cloudformation.Stack) *Service {
	outputs := stackOutputs(stack)
	parameters := stackParameters(stack)
	tags := stackTags(stack)

	url := fmt.Sprintf("redis://u:%s@%s:%s/%s", outputs["EnvRedisPassword"], outputs["Port6379TcpAddr"], outputs["Port6379TcpPort"], outputs["EnvRedisDatabase"])

	if tags["Service"] == "postgres" {
		url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", outputs["EnvPostgresUsername"], outputs["EnvPostgresPassword"], outputs["Port5432TcpAddr"], outputs["Port5432TcpPort"], outputs["EnvPostgresDatabase"])
	}

	return &Service{
		Name:       cs(stack.StackName, "<unknown>"),
		Status:     humanStatus(*stack.StackStatus),
		Outputs:    outputs,
		Parameters: parameters,
		Tags:       tags,
		URL:        url,
	}
}
