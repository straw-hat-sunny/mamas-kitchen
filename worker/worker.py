from azure.storage.queue import QueueServiceClient
import time
import json

# QueueService connects to the Azure Queue Storage and returns a queue client based on the queue name
class QueueService:
    def __init__(self, connection_string):
        self.connection_string = connection_string
        self.queue_service_client = QueueServiceClient.from_connection_string(conn_str=connection_string)

    def get_queue_client(self, queue_name):
        return self.queue_service_client.get_queue_client(queue_name)

# Worker listens to a queue and dequeues messages given a max number, performs a function, deletes the message, and sends a message to another queue
class Worker:
    def __init__(self, queue_service, recieve_queue_name, max_msgs, send_queue_name):
        self.recieve_queue_client = queue_service.get_queue_client(recieve_queue_name)
        self.send_queue_client = queue_service.get_queue_client(send_queue_name)
        self.max_msgs = max_msgs  
    
    def run(self, func:callable):
        while True:
            messages = self.recieve_queue_client.receive_messages(max_messages=self.max_msgs)
            for message in messages:
                msg = json.loads(message.content)
                func(msg)
                self.recieve_queue_client.delete_message(message)
            time.sleep(5)