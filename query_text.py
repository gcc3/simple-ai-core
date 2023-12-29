import os
import sys
import json

from dotenv import load_dotenv, find_dotenv

from langchain.document_loaders import TextLoader
from langchain.indexes import VectorstoreIndexCreator
from langchain.llms import OpenAI
from langchain.chat_models import ChatOpenAI


def process_query(query):
    
    # load environment variables
    load_dotenv(find_dotenv())
    model_name = os.environ["MODEL_NAME"]
    temperature = os.environ["TEMPERATURE"]
    custom_data_only = os.environ["CUSTOM_DATA_ONLY"] == 'true'
    
    # create LLM
    model = ChatOpenAI(model_name=model_name, temperature=temperature, streaming=False)

    # load documents
    textLoader = TextLoader('./data.txt', autodetect_encoding=True)

    # create index from loader
    index = VectorstoreIndexCreator().from_loaders([textLoader])

    PROMPT = """ 
    {question}
    """

    # with LLM, with outside data
    # Note, if without LLM, the output is from custom data only
    if custom_data_only:
        result = index.query(PROMPT.format(question=query))
    else:
        result = index.query(PROMPT.format(question=query), llm=model)
    
    return json.dumps({
        "result": result
    })


if __name__ == '__main__':
    if len(sys.argv) > 1:
        print(process_query(sys.argv[1]))
    else:
        print('No input')