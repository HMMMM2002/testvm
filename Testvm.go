package test

import(
	"testing"

	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestVMCreation(t*testing.T){
	terraformOptions:=&terraform.Options{
TerraformDir:"../path/to/terraform/code",
	}

	defer terraform.Destory(t,terraformOptions)
	terraform.InitAndApply(t,terraformOptions)

	projectID:=gcp.GetGoogleProjectIDFromEnvVar(t)
	region:="us-central1"
	instanceName:="my-instance"

	instance ,err:=gcp.GetComputeInstance(t,projectID,region,instanceName)
	assert.NoError(t,err)
	assert.NoNil(t,instance)
	assert.True(t,instance.Running)
}