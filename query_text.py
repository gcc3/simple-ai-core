import os
import sys
import json

from dotenv import load_dotenv, find_dotenv
from langchain_community.document_loaders import TextLoader
from langchain_openai import ChatOpenAI
from langchain_openai.embeddings import OpenAIEmbeddings
from langchain.indexes import VectorstoreIndexCreator


def process_query(query):
    
    # load environment variables
    load_dotenv(find_dotenv())
    model_name = os.environ["MODEL_NAME"]
    temperature = os.environ["TEMPERATURE"]
    
    # create LLM
    llm = ChatOpenAI(model_name=model_name, temperature=temperature, streaming=False)

    # load documents
    loader = TextLoader('./data.txt', autodetect_encoding=True)

    # create index from loader
    index = VectorstoreIndexCreator(embedding=OpenAIEmbeddings()).from_loaders([loader])

    PROMPT = """ 
    {question}
    """

    # Query llm with data
    result = index.query(PROMPT.format(question=query), llm=llm)
    return json.dumps({
        "result": result
    })


if __name__ == '__main__':
    if len(sys.argv) > 1:
        print(process_query(sys.argv[1]))
    else:
        print('No input')
