import logging 
import os
from azure.storage.queue import QueueServiceClient
from azure.storage.blob import BlobServiceClient

logger = logging.getLogger("azure.core.pipeline.policies.http_logging_policy")
logger.setLevel(logging.WARNING)


# QueueService connects to the Azure Queue Storage and returns a queue client based on the queue name
class QueueService:
    def __init__(self):
        connection_string = "AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;DefaultEndpointsProtocol=http;BlobEndpoint=http://azureite:10000/devstoreaccount1;QueueEndpoint=http://azureite:10001/devstoreaccount1;TableEndpoint=http://azureite:10002/devstoreaccount1;"

        env = os.getenv("ENV")
        if env is None or env == "dev":
            self.queue_service_client = QueueServiceClient.from_connection_string(conn_str=connection_string)

    def get_queue_client(self, queue_name):
        return self.queue_service_client.get_queue_client(queue_name)


# BlobService connects to the Azure Blob Store and returns a blob client based on a container name
class BlobService:
    def __init__(self):
        connection_string = "AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;DefaultEndpointsProtocol=http;BlobEndpoint=http://azureite:10000/devstoreaccount1;QueueEndpoint=http://azureite:10001/devstoreaccount1;TableEndpoint=http://azureite:10002/devstoreaccount1;"

        env = os.getenv("ENV")
        if env is None or env == "dev":
            self.blob_service_client = BlobServiceClient.from_connection_string(conn_str=connection_string)

    def get_container_client(self, container_name):
        return self.blob_service_client.get_container_client(container_name)