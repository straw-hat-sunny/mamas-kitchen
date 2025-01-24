from abc import ABC, abstractmethod
import logging
import requests

logging.basicConfig(level=logging.INFO)

class Transcriber(ABC):
    @abstractmethod
    def run(self, data:bytes):
        pass


class LocalTranscriber(Transcriber):
    def __init__(self):
        pass

    def run(self, data:bytes ):
        logging.info("Local Transcriber Running")
        
        response = requests.post(
            "http://localhost:8080/v1/audio/transcriptions",
            headers={"Content-Type": "multipart/form-data"},
            files={"file": data},
            data={"model": "whisper-1"}
        )

        if response.status_code == 200:
            logging.info("Transcription successful")
            logging.info(response.json())
        else:
            logging.error(f"Transcription failed with status code {response.status_code}")



class OpenAITranscriber(Transcriber):
    def __init__(self):
        pass
    
    def run(self, data:bytes):
        logging.info("OpenAI Transcriber")



