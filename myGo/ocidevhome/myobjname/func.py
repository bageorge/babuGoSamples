import io
import json
import logging
import oci

from fdk import response


def handler(ctx, data: io.BytesIO = None):
    name = "World"
    try:
        body = json.loads(data.getvalue())
        name = body.get("name")
    except (Exception, ValueError) as ex:
        logging.getLogger().info('error parsing json payload: ' + str(ex))

    logging.getLogger().info("Inside Python Hello World function")
    #object_storage_client=oci.object_storage.ObjectStorageClient(oci.config.from_file())
    #result=object_storage_client.get_namespace()
    #logging.getLogger().info("Current object storage namespace: {}".format(result.data))

    return response.Response(
        ctx, response_data=json.dumps(
            {"message": "Hello {0}".format(name)}),
        headers={"Content-Type": "application/json"}
    )
