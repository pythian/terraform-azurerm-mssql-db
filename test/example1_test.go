package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/example1",
	}
	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndApply(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	outputValue := terraform.Output(t, terraformOptions, "output_name")
	assert.NotNil(outputValue)
	assert.Equal("output_value", outputValue)
}
