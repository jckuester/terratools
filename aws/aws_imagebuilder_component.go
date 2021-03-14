// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListImagebuilderComponent(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Imagebuilderconn.ListComponents(ctx, &imagebuilder.ListComponentsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.ComponentVersionList) > 0 {

		for _, r := range resp.ComponentVersionList {

			result = append(result, terraform.Resource{
				Type:      "aws_imagebuilder_component",
				ID:        *r.Arn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
