
import time
import json
from abc import ABC, abstractmethod
from az import QueueService



# Worker listens to a queue and dequeues messages given a max number, performs a function, deletes the message, and sends a message to another queue
class Worker(ABC):
    def __init__(self, recieve_queue_name, max_msgs, send_queue_name):
        queue_service = QueueService()
        self.recieve_queue_client = queue_service.get_queue_client(recieve_queue_name)
        self.send_queue_client = queue_service.get_queue_client(send_queue_name)

        self.max_msgs = max_msgs  
    @abstractmethod
    def process(self,msg):
        pass

    def run(self):
        while True:
            messages = self.recieve_queue_client.receive_messages(max_messages=self.max_msgs)
            for message in messages:
                msg = json.loads(message.content)
                self.process(msg)
                self.recieve_queue_client.delete_message(message)
            time.sleep(15)
    