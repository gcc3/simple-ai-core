import os
import sys

from langchain import OpenAI, SQLDatabase, SQLDatabaseChain
from langchain.chat_models import ChatOpenAI

from dotenv import load_dotenv, find_dotenv
load_dotenv(find_dotenv())
model_name = os.environ["MODEL_NAME"]
temperature = os.environ["TEMPERATURE"]

# db setup
db_username = os.environ["DB_USERNAME"]
db_password = os.environ["DB_PASSWORD"]
db_host = os.environ["DB_HOST"]
db_port = os.environ["DB_PORT"]
db_database = os.environ["DB_DATABASE"]

db_uri = f"mysql+pymysql://{db_username}:{db_password}@{db_host}:{db_port}/{db_database}"
db = SQLDatabase.from_uri(db_uri)

llm = ChatOpenAI()

# verbose=True for debugging
use_verbose = os.environ["USE_VERBOSE"] == "true"
db_chain = SQLDatabaseChain(llm=llm, database=db, verbose=use_verbose)

PROMPT = """ 
Given an input question, first create a syntactically correct MySQL query to run,  
then look at the results of the query and return the answer.  
The question: {question}
"""

def process_query(query):
    result = db_chain.run(PROMPT.format(question=query))
    return result

if __name__ == '__main__':
    if len(sys.argv) > 1:
        print(process_query(sys.argv[1]))
    else:
        print('No input')