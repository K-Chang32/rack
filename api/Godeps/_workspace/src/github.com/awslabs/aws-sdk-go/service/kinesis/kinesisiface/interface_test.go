// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package kinesisiface_test

import (
	"testing"

	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/service/kinesis"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*kinesisiface.KinesisAPI)(nil), kinesis.New(nil))
}
