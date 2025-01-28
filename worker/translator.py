from abc import ABC, abstractmethod
import logging


class Translator(ABC):
    @abstractmethod
    def run(self, context:str) -> dict:
        pass

class LocalTranslator(Translator):
    def run(self, context:str) -> str:
        logging.info("Local Translator Running...")
        