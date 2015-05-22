package ecsiface

import (
	"github.com/convox/kernel/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/service/ecs"
)

type ECSAPI interface {
	CreateCluster(*ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error)

	CreateService(*ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error)

	DeleteCluster(*ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error)

	DeleteService(*ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error)

	DeregisterContainerInstance(*ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error)

	DeregisterTaskDefinition(*ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error)

	DescribeClusters(*ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error)

	DescribeContainerInstances(*ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error)

	DescribeServices(*ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error)

	DescribeTaskDefinition(*ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error)

	DescribeTasks(*ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error)

	DiscoverPollEndpoint(*ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error)

	ListClusters(*ecs.ListClustersInput) (*ecs.ListClustersOutput, error)

	ListContainerInstances(*ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error)

	ListServices(*ecs.ListServicesInput) (*ecs.ListServicesOutput, error)

	ListTaskDefinitionFamilies(*ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error)

	ListTaskDefinitions(*ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error)

	ListTasks(*ecs.ListTasksInput) (*ecs.ListTasksOutput, error)

	RegisterContainerInstance(*ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error)

	RegisterTaskDefinition(*ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error)

	RunTask(*ecs.RunTaskInput) (*ecs.RunTaskOutput, error)

	StartTask(*ecs.StartTaskInput) (*ecs.StartTaskOutput, error)

	StopTask(*ecs.StopTaskInput) (*ecs.StopTaskOutput, error)

	SubmitContainerStateChange(*ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error)

	SubmitTaskStateChange(*ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error)

	UpdateService(*ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error)
}
