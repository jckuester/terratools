// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListElasticacheReplicationGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Elasticacheconn.DescribeReplicationGroupsRequest(&elasticache.DescribeReplicationGroupsInput{})

	var result []terraform.Resource

	p := elasticache.NewDescribeReplicationGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ReplicationGroups {

			result = append(result, terraform.Resource{
				Type:      "aws_elasticache_replication_group",
				ID:        *r.ReplicationGroupId,
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
