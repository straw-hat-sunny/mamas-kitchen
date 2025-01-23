from abc import ABC, abstractmethod
import logging

logging.basicConfig(level=logging.INFO)

class Transcriber(ABC):
    @abstractmethod
    def run(self, data:bytes):
        pass


class LocalTranscriber(Transcriber):
    def __init__(self):
        pass

    def run(self, data:bytes ):
        logging.info("Local Transcriber Ran")



class OpenAITranscriber(Transcriber):
    def __init__(self):
        pass
    
    def run(self, data:bytes):
        logging.info("OpenAI Transcriber")



