// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDxTransitVirtualInterface(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Directconnectconn.DescribeVirtualInterfaces(ctx, &directconnect.DescribeVirtualInterfacesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.VirtualInterfaces) > 0 {

		for _, r := range resp.VirtualInterfaces {

			result = append(result, terraform.Resource{
				Type:      "aws_dx_transit_virtual_interface",
				ID:        *r.VirtualInterfaceId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
