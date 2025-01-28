from worker import Worker 
import logging
import json
from translator import Translator, LocalTranslator

logging.basicConfig(level=logging.INFO)

class TranslatorWorker(Worker):
    def __init__(self, translator:Translator,recieve_queue_name, max_msg, send_queue_name):
        super().__init__(recieve_queue_name,max_msg,send_queue_name)
        self.translator = translator

    def process(self,msg):
        if 'text' not in msg:
            logging.error(f"Message does not contain 'text' property\n {msg}")
            return

        recipe_text = msg['text']
      
        logging.info(f"Processing recipe:")
        try:
            recipe_object = translator.run(recipe_text)
            logging.info(recipe_object)
            processed_msg = json.dumps(recipe_object)
            return processed_msg
        except Exception as e:
            logging.info("translator failed")
            logging.error(e)
            return

translator = LocalTranslator()
tw = TranslatorWorker(translator, "blob-transcribed", 1, "blob-translated")
tw.run()



