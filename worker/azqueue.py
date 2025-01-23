

from worker import QueueService, Worker
from azure.storage.blob import BlobServiceClient

import logging

logging.basicConfig(level=logging.INFO)
logging.info("Transcriber Worker is running...")


# Define the connection string and queue name
connection_string = "AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;DefaultEndpointsProtocol=http;BlobEndpoint=http://azureite:10000/devstoreaccount1;QueueEndpoint=http://azureite:10001/devstoreaccount1;TableEndpoint=http://azureite:10002/devstoreaccount1;"

# Create a BlobServiceClient
blob_service_client = BlobServiceClient.from_connection_string(conn_str=connection_string)

# Get the blob client
blob_container_client = blob_service_client.get_container_client("audio-files")

queue_service = QueueService()

worker = Worker(queue_service, "audio-files", 1, "transform")









