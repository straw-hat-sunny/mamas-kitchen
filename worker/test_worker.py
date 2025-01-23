from worker import QueueService, Worker 
import logging

logging.basicConfig(level=logging.INFO)

class TestWorker:
    def __init__(self):
        connection_string = "AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;DefaultEndpointsProtocol=http;BlobEndpoint=http://azureite:10000/devstoreaccount1;QueueEndpoint=http://azureite:10001/devstoreaccount1;TableEndpoint=http://azureite:10002/devstoreaccount1;"
        self.queue_service = QueueService(connection_string)
        self.worker = Worker(self.queue_service, "newTask", 1, "finishTask")

    def doSomething():
        logging.info("This does something")

    def process(self):
        self.worker.run(self.doSomething)



tw = TestWorker()
tw.process()