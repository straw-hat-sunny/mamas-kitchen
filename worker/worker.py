
import time
import json
from abc import ABC, abstractmethod
from az import QueueService
import logging


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
            try:
                messages = self.recieve_queue_client.receive_messages(max_messages=self.max_msgs)
                for message in messages:
                    msg_content = json.loads(message.content)
                    processed_msg = self.process(msg_content)
                    # if there are no errors processing the message then we can delete the message from the queue and send message to the send_aueue
                    self.recieve_queue_client.delete_message(message)
                    self.send_queue_client.send_message(processed_msg)
            except Exception as e:
                 logging.info(e)
                    
            time.sleep(15)
    