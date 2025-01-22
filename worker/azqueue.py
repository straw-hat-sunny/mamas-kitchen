
import time
from azure.storage.queue import QueueServiceClient
from azure.storage.blob import BlobServiceClient
import json
import logging

logging.basicConfig(level=logging.INFO)
logging.info("Transcriber Worker is running...")


# Define the connection string and queue name
connection_string = "AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;DefaultEndpointsProtocol=http;BlobEndpoint=http://127.0.0.1:10000/devstoreaccount1;QueueEndpoint=http://127.0.0.1:10001/devstoreaccount1;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;"
queue_name = "audio-files"

# Create a QueueServiceClient
queue_service_client = QueueServiceClient.from_connection_string(conn_str=connection_string)

# Create a BlobServiceClient
blob_service_client = BlobServiceClient.from_connection_string(conn_str=connection_string)

# Get the queue client
queue_client = queue_service_client.get_queue_client(queue_name)

# Get the blob client
blob_container_client = blob_service_client.get_container_client("audio-files")

# Dequeue a message


while True:
    print("Waiting for messages...")
    messages = queue_client.receive_messages(max_messages=1)
    for message in messages:
        decoded_message = json.loads(message.content)
        file_name = decoded_message['FileName'] 
        print(f"Dequeued message: {file_name}")
        blob_client = blob_container_client.get_blob_client(file_name)
        # Download Page Blob
        with open('dump/'+file_name, "wb") as my_blob:
            download_stream = blob_client.download_blob()
            my_blob.write(download_stream.readall())
        # Delete the message from the queue
        queue_client.delete_message(message)

        print("Message dequeued and deleted successfully.")
    time.sleep(5)

