#!/bin/sh

#curl -X POST -d '{"token":"ABCD"}' https://mfsbz4gcnkehslkz7jarfx3vaq.apigateway.us-ashburn-1.oci.customer-oci.com/tokens/helloapi
curl -X GET -d '{"token":"70:6c:e3:b4:49:4e:06:67:fa:ac:0c:74:9d:c5:91:a9"}' https://mfsbz4gcnkehslkz7jarfx3vaq.apigateway.us-ashburn-1.oci.customer-oci.com/tokens/helloapi|jq .
echo ""
