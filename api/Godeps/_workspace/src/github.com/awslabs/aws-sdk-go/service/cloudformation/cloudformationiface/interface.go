// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudformationiface provides an interface that satisfies the cloudformation service.
package cloudformationiface

import (
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/service/cloudformation"
)

type CloudFormationAPI interface {
	CancelUpdateStack(*cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error)

	CreateStack(*cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)

	DeleteStack(*cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error)

	DescribeStackEvents(*cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error)

	DescribeStackResource(*cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error)

	DescribeStackResources(*cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error)

	DescribeStacks(*cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)

	EstimateTemplateCost(*cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error)

	GetStackPolicy(*cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error)

	GetTemplate(*cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error)

	GetTemplateSummary(*cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error)

	ListStackResources(*cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error)

	ListStacks(*cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error)

	SetStackPolicy(*cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error)

	SignalResource(*cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error)

	UpdateStack(*cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error)

	ValidateTemplate(*cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error)
}
