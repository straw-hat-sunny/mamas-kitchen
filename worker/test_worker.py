from worker import Worker 
import logging

logging.basicConfig(level=logging.INFO)

class TestWorker(Worker):
    def process(self,msg):
        if 'FileName' not in msg:
            logging.error(f"Message does not contain 'filename' property\n {msg}")
            return

        filename = msg['FileName']
        logging.info(f"Processing file: {filename}")



tw = TestWorker("audio-files", 1, "audio-files")
tw.run()
