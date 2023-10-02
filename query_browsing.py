import os
import sys

from dotenv import load_dotenv, find_dotenv

from langchain.chat_models import ChatOpenAI
from langchain.agents import load_tools, initialize_agent
from langchain.agents import AgentType
from langchain.tools import AIPluginTool

def process_query(query):
    
    # load environment variables
    load_dotenv(find_dotenv())
    openai_api_key = os.environ["OPENAI_API_KEY"]
    model_name = os.environ["MODEL_NAME"]
    temperature = os.environ["TEMPERATURE"]
    use_verbose = os.environ["USE_VERBOSE"] == "true"

    # Use VoxScript as browsing engine
    tool = AIPluginTool.from_plugin_url("https://voxscript.awt.icu/.well-known/ai-plugin.json")
    tools = load_tools(["requests_all"])
    tools += [tool]
    
    # create LLM
    llm = ChatOpenAI(openai_api_key=openai_api_key, model_name=model_name, temperature=temperature)
    
    # create agent chain
    agent_chain = initialize_agent(
        tools, llm, agent=AgentType.ZERO_SHOT_REACT_DESCRIPTION, verbose=use_verbose
    )

    PROMPT = """ 
    {question}
    """
    
    result = agent_chain.run(PROMPT.format(question=query))
    return result


if __name__ == '__main__':
    if len(sys.argv) > 1:
        print(process_query(sys.argv[1]))
    else:
        print('No input')