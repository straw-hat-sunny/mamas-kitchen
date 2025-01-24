from worker import QueueService, Worker 
import logging

logging.basicConfig(level=logging.INFO)

class TestWorker(Worker):
    def process(self,msg):
        logging.info(f"Processing message: {msg}")


        
connection_string = "AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;DefaultEndpointsProtocol=http;BlobEndpoint=http://azureite:10000/devstoreaccount1;QueueEndpoint=http://azureite:10001/devstoreaccount1;TableEndpoint=http://azureite:10002/devstoreaccount1;"
queue_service = QueueService(connection_string)

tw = TestWorker(queue_service,"audio-files", 1, "audio-files")
tw.run()
