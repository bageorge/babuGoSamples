// package example
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/oracle/oci-go-sdk/v45/common"
	"github.com/oracle/oci-go-sdk/v45/core"
	"github.com/oracle/oci-go-sdk/v45/example/helpers"
)

//func ExampleCreateVcn() {
func main() {
	displayName := "OCI-GOSDK-CreateVcn-Example"
	compartmentID := os.Getenv("OCI_COMPARTMENT_ID") // OCI_COMPARTMENT_ID env variable must be defined

	// initialize VirtualNetworkClient
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	ctx := context.Background()

	// create VCN
	createVcnRequest := core.CreateVcnRequest{
		RequestMetadata: helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}
	createVcnRequest.CompartmentId = common.String(compartmentID)
	createVcnRequest.DisplayName = common.String(displayName)
	createVcnRequest.CidrBlock = common.String("10.0.0.0/16")
	_, err = client.CreateVcn(ctx, createVcnRequest)
	helpers.FatalIfError(err)

	fmt.Println("VCN created")

	// Output:
	// VCN created
}
