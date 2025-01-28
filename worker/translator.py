from abc import ABC, abstractmethod
import logging
import requests


class Translator(ABC):
    @abstractmethod
    def run(self, context:str) -> dict:
        pass

class LocalTranslator(Translator):
    def run(self, context:str) -> dict:
        logging.info("Local Translator Running...")
        response = requests.post(
            "http://open-ai:8080//v1/chat/completions",
            data = {
                "model":"gpt-4o",
                "messages": [
                    {
                        "role": "system",
                        "content": [
                            {
                                "type": "text",
                                "text": "You convert recipes into json objects following the schema:\ndish_name: name of the dish\ndish_type: if the dish is a appetizer, side dish, entree, or dessert\ningredients: a list of objects {item, measurement, quantity}\nsteps: list of steps/methods"
                            }
                        ]
                    },
                    {
                        "role": "user",
                        "content": [
                            {
                                "type": "text",
                                "text": {context}
                            }
                        ]
                    }
                ],
                "response_format":{
                    "type": "json_object"
                },
                "temperature":1,
                "max_tokens":2048,
                "top_p":1,
                "frequency_penalty":0,
                "presence_penalty":0
            }
        )
        if response.status_code == 200:
            logging.info("Translation successful")
            recipe = response.json()["choices"][0]["message_content"]
            logging.info(recipe)
            return recipe

        