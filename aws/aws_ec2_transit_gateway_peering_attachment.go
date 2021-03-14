// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEc2TransitGatewayPeeringAttachment(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := ec2.NewDescribeTransitGatewayPeeringAttachmentsPaginator(client.Ec2conn, &ec2.DescribeTransitGatewayPeeringAttachmentsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.TransitGatewayPeeringAttachments {

			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_ec2_transit_gateway_peering_attachment",
				ID:        *r.TransitGatewayAttachmentId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
