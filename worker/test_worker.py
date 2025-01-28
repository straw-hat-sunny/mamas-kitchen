from worker import Worker 
import logging
import json

logging.basicConfig(level=logging.INFO)

class TestWorker(Worker):
    def process(self,msg):
        logging.info(msg)
        if 'text' not in msg:
            logging.error(f"Message does not contain 'text' property\n {msg}")
            return

        filename = msg['text']
        logging.info(f"Processing recipe: {filename}")

        processed_msg = json.dumps({"status":"complete"})
        return processed_msg


tw = TestWorker("blob-transcribed", 1, "blob-translated")
tw.run()
