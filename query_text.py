import os
import sys

from dotenv import load_dotenv, find_dotenv
from langchain.document_loaders import TextLoader
from langchain.indexes import VectorstoreIndexCreator
from langchain.llms import OpenAI
from langchain.chat_models import ChatOpenAI

# with LLM, with outside data
# Note, if without LLM, the output is from custom data only
def process_query(query):
    load_dotenv(find_dotenv())
    model_name = os.environ["MODEL_NAME"]
    temperature = os.environ["TEMPERATURE"]
    model = ChatOpenAI(model_name=model_name, temperature=temperature, streaming=False)

    # load documents
    textLoader = TextLoader('./data.txt', autodetect_encoding=True)

    # create index from loader
    index = VectorstoreIndexCreator().from_loaders([textLoader])

    PROMPT = """ 
    {question}
    """

    result = index.query(PROMPT.format(question=query), llm=model)
    return result

if __name__ == '__main__':
    if len(sys.argv) > 1:
        print(process_query(sys.argv[1]))
    else:
        print('No input')