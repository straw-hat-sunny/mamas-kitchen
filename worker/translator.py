from abc import ABC, abstractmethod
import logging
import requests
import json

class Translator(ABC):
    @abstractmethod
    def run(self, context:str) -> dict:
        pass

class LocalTranslator(Translator):
    def run(self, context:str) -> dict:
        logging.info("Local Translator Running...")

        response = requests.post(
            "http://local-ai:8080//v1/chat/completions",
            json = {
                "model": "phi-4",
                "messages": [
                    {
                        "role": "system",
                        "content":  "You convert recipes into json objects following the schema: dish_name: name of the dish. dish_type: if the dish is a appetizer, side dish, entree, or dessert. ingredients: a list of objects {item, measurement, quantity}. steps: list of steps/methods"
                    },
                    {
                        "role": "user",
                        "content": context
                    }
                ],
                "temperature":1,
            },
            headers={"Content-Type": "application/json"}
        )
        if response.status_code == 200:
            logging.info("Translation successful")
            recipe = response.json()["choices"][0]["message"]["content"]
            cleaned_json_string = recipe.replace('`',"").replace('json',"").strip()

            data = json.loads(cleaned_json_string)
            return data

    