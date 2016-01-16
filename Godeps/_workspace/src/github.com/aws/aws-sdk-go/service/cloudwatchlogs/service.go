// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatchlogs

import (
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/request"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/internal/signer/v4"
)

// This is the Amazon CloudWatch Logs API Reference. Amazon CloudWatch Logs
// enables you to monitor, store, and access your system, application, and custom
// log files. This guide provides detailed information about Amazon CloudWatch
// Logs actions, data types, parameters, and errors. For detailed information
// about Amazon CloudWatch Logs features and their associated API calls, go
// to the Amazon CloudWatch Developer Guide (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide).
//
// Use the following links to get started using the Amazon CloudWatch Logs
// API Reference:
//
//   Actions (http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_Operations.html):
// An alphabetical list of all Amazon CloudWatch Logs actions.  Data Types (http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_Types.html):
// An alphabetical list of all Amazon CloudWatch Logs data types.  Common Parameters
// (http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/CommonParameters.html):
// Parameters that all Query actions can use.  Common Errors (http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/CommonErrors.html):
// Client and server errors that all actions can return.  Regions and Endpoints
// (http://docs.aws.amazon.com/general/latest/gr/index.html?rande.html): Itemized
// regions and endpoints for all AWS products.  In addition to using the Amazon
// CloudWatch Logs API, you can also use the following SDKs and third-party
// libraries to access Amazon CloudWatch Logs programmatically.
//
//  AWS SDK for Java Documentation (http://aws.amazon.com/documentation/sdkforjava/)
// AWS SDK for .NET Documentation (http://aws.amazon.com/documentation/sdkfornet/)
// AWS SDK for PHP Documentation (http://aws.amazon.com/documentation/sdkforphp/)
// AWS SDK for Ruby Documentation (http://aws.amazon.com/documentation/sdkforruby/)
//  Developers in the AWS developer community also provide their own libraries,
// which you can find at the following AWS developer centers:
//
//  AWS Java Developer Center (http://aws.amazon.com/java/) AWS PHP Developer
// Center (http://aws.amazon.com/php/) AWS Python Developer Center (http://aws.amazon.com/python/)
// AWS Ruby Developer Center (http://aws.amazon.com/ruby/) AWS Windows and .NET
// Developer Center (http://aws.amazon.com/net/)
type CloudWatchLogs struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new CloudWatchLogs client.
func New(config *aws.Config) *CloudWatchLogs {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "logs",
			APIVersion:   "2014-03-28",
			JSONVersion:  "1.1",
			TargetPrefix: "Logs_20140328",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &CloudWatchLogs{service}
}

// newRequest creates a new request for a CloudWatchLogs operation and runs any
// custom request initialization.
func (c *CloudWatchLogs) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
