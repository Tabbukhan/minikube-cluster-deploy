package test

import (
    "testing"

    "github.com/gruntwork-io/terratest/modules/aws"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraformAWSExample(t *testing.T) {
    t.Parallel()

    terraformOptions := &terraform.Options{
        TerraformDir: "../terraform",
    }

    defer terraform.Destroy(t, terraformOptions)

    terraform.InitAndApply(t, terraformOptions)

    asgName := terraform.Output(t, terraformOptions, "example-autoscaling-group")
    awsRegion := "us-west-2"

    instances := aws.GetInstancesForAsg(t, asgName, awsRegion)

    assert.Equal(t, 1, len(instances))
}
