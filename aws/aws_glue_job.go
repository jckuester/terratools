// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListGlueJob(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Glueconn.GetJobsRequest(&glue.GetJobsInput{})

	var result []terraform.Resource

	p := glue.NewGetJobsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Jobs {

			result = append(result, terraform.Resource{
				Type:      "aws_glue_job",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
