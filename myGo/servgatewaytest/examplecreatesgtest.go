package main

import (
	"context"
	"fmt"
	"os"

	"github.com/oracle/oci-go-sdk/v45/common"
	"github.com/oracle/oci-go-sdk/v45/core"
	"github.com/oracle/oci-go-sdk/v45/example/helpers"
)

func main() {
	displayName := "OCI-GOSDK-CreateServiceGateway-Example" // displayName for created VCN and ServiceGateway
	compartmentID := os.Getenv("OCI_COMPARTMENT_ID")        // OCI_COMPARTMENT_ID env variable must be defined
        vcnId := "ocid1.vcn.oc1.iad.amaaaaaasg6f64qazfst4kxlqvijif6girxdogpgnfmbgupkoiguq36ryyeq"

	// initialize VirtualNetworkClient
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	ctx := context.Background()

	// create VCN
	//createVcnRequest := core.CreateVcnRequest{
	//	RequestMetadata: helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	//}
	//createVcnRequest.CompartmentId = common.String(compartmentID)
	//createVcnRequest.DisplayName = common.String(displayName)
	//createVcnRequest.CidrBlock = common.String("10.0.0.0/16")
	//createVcnResponse, err := client.CreateVcn(ctx, createVcnRequest)
	//helpers.FatalIfError(err)

	// create ServiceGateway
	createServiceGatewayRequest := core.CreateServiceGatewayRequest{
		RequestMetadata: helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}
	createServiceGatewayRequest.CompartmentId = common.String(compartmentID)
	createServiceGatewayRequest.DisplayName = common.String(displayName)
	//createServiceGatewayRequest.VcnId = createVcnResponse.Id
	createServiceGatewayRequest.VcnId = common.String(vcnId)
	createServiceGatewayRequest.Services = []core.ServiceIdRequestDetails{}
	_, err = client.CreateServiceGateway(ctx, createServiceGatewayRequest)
	helpers.FatalIfError(err)

	fmt.Println("ServiceGateway created")
	// Output:
	// ServiceGateway created
        fmt.Println("Associating a route table")
        fmt.Println("List of all route tables")
        
}
