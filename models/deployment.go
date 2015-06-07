package models

import (
	"os"
	"time"

	"github.com/convox/kernel/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/aws"
	"github.com/convox/kernel/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/service/ecs"
)

type Deployment struct {
	Status  string
	Release string

	Desired int64
	Pending int64
	Running int64

	Created time.Time
}

type Deployments []Deployment

func ListDeployments(app string) (Deployments, error) {
	res, err := ECS().DescribeServices(&ecs.DescribeServicesInput{
		Cluster:  aws.String(os.Getenv("CLUSTER")),
		Services: []*string{aws.String(app)},
	})

	if err != nil {
		return nil, err
	}

	// no service yet, so no deployments
	if len(res.Services) != 1 {
		return Deployments{}, nil
	}

	service := res.Services[0]

	deployments := make(Deployments, len(service.Deployments))

	for i, d := range service.Deployments {
		deployments[i] = Deployment{
			Status:  *d.Status,
			Desired: *d.DesiredCount,
			Pending: *d.PendingCount,
			Running: *d.RunningCount,
			Created: *d.CreatedAt,
		}

		tres, err := ECS().DescribeTaskDefinition(&ecs.DescribeTaskDefinitionInput{
			TaskDefinition: d.TaskDefinition,
		})

		if err != nil {
			return nil, err
		}

		if len(tres.TaskDefinition.ContainerDefinitions) > 0 {
			for _, kp := range tres.TaskDefinition.ContainerDefinitions[0].Environment {
				if *kp.Name == "RELEASE" {
					deployments[i].Release = *kp.Value
					break
				}
			}
		}
	}

	return deployments, nil
}
