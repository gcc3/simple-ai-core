import os
import sys
import json

from dotenv import load_dotenv, find_dotenv

from langchain_openai import OpenAI, ChatOpenAI
from langchain import SQLDatabaseChain
from langchain_community.utilities.sql_database import SQLDatabase


def process_query(query):
    
    # load environment variables
    load_dotenv(find_dotenv())
    openai_api_key = os.environ["OPENAI_API_KEY"]
    model_name = os.environ["MODEL_NAME"]
    temperature = os.environ["TEMPERATURE"]
    db_username = os.environ["DB_USERNAME"]
    db_password = os.environ["DB_PASSWORD"]
    db_host = os.environ["DB_HOST"]
    db_port = os.environ["DB_PORT"]
    db_database = os.environ["DB_DATABASE"]
    use_verbose = os.environ["USE_VERBOSE"] == "true"

    # create database connection
    db_uri = f"mysql+pymysql://{db_username}:{db_password}@{db_host}:{db_port}/{db_database}"
    db = SQLDatabase.from_uri(db_uri)

    # create LLM
    model = ChatOpenAI(openai_api_key=openai_api_key, model_name=model_name, temperature=temperature)

    # create database chain
    db_chain = SQLDatabaseChain(llm=model, database=db, verbose=use_verbose)

    PROMPT = """ 
    Given an input question, first create a syntactically correct MySQL query to run,  
    then look at the results of the query and return the answer.  
    The question: {question}  
    Try different word while quering data.
    Example:
    User input word is 東京, try 東京 and 東京都 both or together.
    Limit output for 20 items.
    """

    result = db_chain.run(PROMPT.format(question=query))
    return json.dumps({
        "result": result
    })


if __name__ == '__main__':
    if len(sys.argv) > 1:
        print(process_query(sys.argv[1]))
    else:
        print('No input')