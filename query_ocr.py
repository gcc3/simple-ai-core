import os
import sys

from dotenv import load_dotenv, find_dotenv

from langchain.chat_models import ChatOpenAI
from langchain.chains.question_answering import load_qa_chain
from langchain.document_loaders import AmazonTextractPDFLoader
import boto3


def process_query(query):
    
    # load environment variables
    load_dotenv(find_dotenv())
    model_name = os.environ["MODEL_NAME"]
    temperature = os.environ["TEMPERATURE"]
    
    # create LLM
    model = ChatOpenAI(model_name=model_name, temperature=temperature, streaming=False)

    # load documents
    loader = AmazonTextractPDFLoader("./data.jpeg", region_name="us-east-1")
    documents = loader.load()

    # create index from loader
    chain = load_qa_chain(llm=model, chain_type="map_reduce")

    PROMPT = """ 
    {question}
    """

    result = chain.run(input_documents=documents, question=PROMPT.format(question=query))
    return result


if __name__ == '__main__':
    if len(sys.argv) > 1:
        print(process_query(sys.argv[1]))
    else:
        print('No input')