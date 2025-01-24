from worker import QueueService, Worker
from transcriber import Transcriber
import logging

logging.basicConfig(level=logging.INFO)

class TranscriberWorker(Worker):
    def __init__(self, transcriber:Transcriber, queue_service,recieve_queue_name, max_msg, send_queue_name):
        super().__init__(queue_service,recieve_queue_name,max_msg,send_queue_name)
        self.transcriber = transcriber

    def process(self,msg):
        pass