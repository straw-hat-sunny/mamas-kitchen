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
            logging.info(recipe)
            return recipe

       


        '''
        curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
  "model": "phi-4",
  "messages": [{"role": "system", "content": "You convert recipes into json objects following the schema: dish_name: name of the dish. dish_type: if the dish is a appetizer, side dish, entree, or dessert. ingredients: a list of objects {item, measurement, quantity}. steps: list of steps/methods"}, {"role":"user", "content":  ":Dollies, Raita, accompaniment or side dish, ingredients, one teaspoon mustard seeds, apinch of aspatoria, half teaspoon urad dal, bunch of curry leaves, three dried chilies,a tablespoon, one tablespoon oil, ten pods garlic chopped, two cans of fire roasted tomatoes,four-bott, half teaspoon turmeric, one teaspoon salt, one teaspoon sugar, quarter cup cilantroleaves, method. In hot oil add mustard seeds, aspatoria or hing, urad dal, curry leaves, redchilies. When this splatters add lots of the chopped garlic, stir for three minutes, thenadd two cans of fire roasted tomatoes, turmeric, salt, sugar and let this cook down. It needsto cook for 10 minutes, then add cilantro leaves, turn off the stove and cool the dish.Once cooled add this to a bowl of whipped yogurt. You can also substitute the red chilieswith chopped green chilies.[BLANK_AUDIO"}],
  "temperature": 0.7,
  "response_format":"json_object"
}'
        
        
        '''