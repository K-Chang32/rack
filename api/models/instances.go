package models

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"

	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/ec2"
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/ecs"
	"github.com/convox/rack/Godeps/_workspace/src/golang.org/x/crypto/ssh"
	"github.com/convox/rack/client"
)

type Instance client.Instance

type InstanceResource struct {
	Total int `json:"total"`
	Free  int `json:"free"`
	Used  int `json:"used"`
}

func (ir InstanceResource) PercentUsed() float64 {
	return float64(ir.Used) / float64(ir.Total)
}

func InstanceKeyroll() error {
	keyname := fmt.Sprintf("%s-keypair-%d", os.Getenv("RACK"), (rand.Intn(8999) + 1000))
	keypair, err := EC2().CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName: &keyname,
	})

	if err != nil {
		return err
	}

	env, err := GetRackSettings()

	if err != nil {
		return err
	}

	//TODO: have to update the CF template params
	env["InstancePEM"] = *keypair.KeyMaterial
	err = PutRackSettings(env)

	if err != nil {
		return err
	}

	app, err := GetApp(os.Getenv("RACK"))
	if err != nil {
		return err
	}

	err = app.UpdateParams(map[string]string{
		"Key": keyname,
	})
	if err != nil {
		return err
	}

	return nil
}

func InstanceSSH(id, command string, height, width int, rw io.ReadWriter) error {

	instanceIds := []*string{&id}
	ec2Res, err := EC2().DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{Name: aws.String("instance-id"), Values: instanceIds},
		},
	})

	if err != nil {
		return err
	}

	instance := ec2Res.Reservations[0].Instances[0]

	env, err := GetRackSettings()
	if err != nil {
		return err
	}

	signer, err := ssh.ParsePrivateKey([]byte(env["InstancePEM"]))
	if err != nil {
		return err
	}
	config := &ssh.ClientConfig{
		User: "ec2-user",
		Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
	}
	conn, err := ssh.Dial("tcp", *instance.PublicIpAddress+":22", config)
	if err != nil {
		return err
	}
	defer conn.Close()
	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stdout = rw
	session.Stdin = rw
	session.Stderr = rw

	if command != "" {
		err = session.Run(command)
		if err != nil {
			return err
		}
	} else {
		modes := ssh.TerminalModes{
			ssh.ECHOCTL:       0,
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		}
		// Request pseudo terminal
		if err := session.RequestPty("xterm", width, height, modes); err != nil {
			return err
		}
		// Start remote shell
		if err := session.Shell(); err != nil {
			return err
		}

		session.Wait()
	}

	return nil
}

func (s *System) GetInstances() ([]*Instance, error) {
	res, err := ECS().ListContainerInstances(
		&ecs.ListContainerInstancesInput{
			Cluster: aws.String(os.Getenv("CLUSTER")),
		},
	)

	if err != nil {
		return nil, err
	}

	ecsRes, err := ECS().DescribeContainerInstances(
		&ecs.DescribeContainerInstancesInput{
			Cluster:            aws.String(os.Getenv("CLUSTER")),
			ContainerInstances: res.ContainerInstanceArns,
		},
	)

	if err != nil {
		return nil, err
	}

	var instanceIds []*string
	for _, i := range ecsRes.ContainerInstances {
		instanceIds = append(instanceIds, i.Ec2InstanceId)
	}

	ec2Res, err := EC2().DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{Name: aws.String("instance-id"), Values: instanceIds},
		},
	})

	if err != nil {
		return nil, err
	}

	ec2Instances := make(map[string]*ec2.Instance)
	for _, r := range ec2Res.Reservations {
		for _, i := range r.Instances {
			ec2Instances[*i.InstanceId] = i
		}
	}

	var instances []*Instance
	for _, i := range ecsRes.ContainerInstances {
		// figure out the CPU and memory metrics
		var cpu, memory InstanceResource

		for _, r := range i.RegisteredResources {
			switch *r.Name {
			case "CPU":
				cpu.Total = int(*r.IntegerValue)
			case "MEMORY":
				memory.Total = int(*r.IntegerValue)
			}
		}

		for _, r := range i.RemainingResources {
			switch *r.Name {
			case "CPU":
				cpu.Free = int(*r.IntegerValue)
				cpu.Used = cpu.Total - cpu.Free
			case "MEMORY":
				memory.Free = int(*r.IntegerValue)
				memory.Used = memory.Total - memory.Free
			}
		}

		// find the matching Instance from the EC2 response
		ec2Instance := ec2Instances[*i.Ec2InstanceId]

		// build up the struct
		instance := &Instance{
			Agent:     *i.AgentConnected,
			Cpu:       truncate(cpu.PercentUsed(), 4),
			Memory:    truncate(memory.PercentUsed(), 4),
			Id:        *i.Ec2InstanceId,
			Ip:        *ec2Instance.PublicIpAddress,
			Processes: int(*i.RunningTasksCount),
			Status:    strings.ToLower(*i.Status),
		}

		instances = append(instances, instance)
	}

	return instances, nil
}
