from worker import Worker
from transcriber import Transcriber, LocalTranscriber
from az import BlobService
import logging
import json

logging.basicConfig(level=logging.INFO)

class TranscriberWorker(Worker):
    def __init__(self, transcriber:Transcriber,recieve_queue_name, max_msg, send_queue_name):
        super().__init__(recieve_queue_name,max_msg,send_queue_name)
        self.transcriber = transcriber
        self.container_service = BlobService().get_container_client("audio-files")

    def process(self,msg):
        if 'FileName' not in msg:
            # add error handling
            logging.error("no file name present")
            return
        file_name = msg['FileName']

        try:
            blob = self.container_service.download_blob(file_name).readall()
        except Exception as e:
            logging.info("container service failed to download blob")
            logging.error(e)
            return

        try:
            transcribed_text = self.transcriber.run(file_name, blob)
            logging.info(transcribed_text)
            msg = {
                "text" : transcribed_text
            }
            msg_content = json.dumps(msg)
            return msg_content
        except Exception as e:
            logging.info("transcriber failed to process audio")
            logging.error(e)
            return

transcriber = LocalTranscriber()

worker = TranscriberWorker(transcriber,"blob-uploaded",1,"blob-transcribed")
worker.run()