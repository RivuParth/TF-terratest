package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestTerraformS3Module(t *testing.T) {
	t.Parallel()

	// Setup the path to the Terraform code that will be tested
	terraformDir := "../"

	// Initialize and apply the Terraform configuration
	options := terraform.Options{
		// Set the path to the Terraform code
		TerraformDir: terraformDir,

		// Variables to pass to our Terraform configuration
		Vars: map[string]interface{}{
			"bucket_name": "test-bucket-for-terratest",
			"bucket_acl":  "private",
		},

		// Variables to override default retryable errors
		RetryableTerraformErrors: map[string]string{
			"Error calling CreateBucket": "BucketAlreadyExists",
		},

		// Whether or not to generate and show a Terraform plan
		NoColor: true,
	}

	// Clean up resources after testing
	defer terraform.Destroy(t, &options)

	// First, initialize Terraform
	terraform.InitAndApply(t, &options)

	// Run `terraform output` to get the value of an output variable
	bucketID := terraform.Output(t, &options, "bucket_id")

	// Verify if the bucket ID is not empty
	assert.NotEmpty(t, bucketID)

	// Optional: Print the output for debugging purposes
	fmt.Printf("Bucket ID: %s\n", bucketID)
}
