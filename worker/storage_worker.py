import pymongo
import logging
from worker import Worker 


logging.basicConfig(level=logging.INFO)
# logging.info("Accessing Database")
# try:
#     client = pymongo.MongoClient("mongodb://root:example@mongo:27017/")
#     dbs = client.list_database_names()
#     logging.info(dbs)
#     logging.info(len(dbs))
# except Exception as e:
#     logging.error(e)

class StorageWorker(Worker):
   
   def __init__(self,recieve_queue_name, max_msg, send_queue_name):
        super().__init__(recieve_queue_name,max_msg,send_queue_name)
        self.client =  pymongo.MongoClient("mongodb://root:example@mongo:27017/")
        self.collection = self.client['local']['recipes']

   
   def process(self, msg):
      if 'recipe' not in msg:
         logging.error(f"Message doesn't the contain 'recipe' property\n{msg}")
         return
      recipe = msg['recipe']

      logging.info(recipe)
      result = self.collection.insert_one(recipe)
      logging.info(f"Inserted document with ID: {result.inserted_id}")
      data = {"status": f"completed: {result.inserted_id}"}
      return data

sw = StorageWorker("blob-translated",1,"storage-completed")

sw.run()