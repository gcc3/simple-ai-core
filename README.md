
Simple AI Core
==============


Provide database connection and ChatGPT plugin support with LangChain.  


Dependencies
------------

LangChain https://python.langchain.com/docs/  
GoLang https://go.dev/doc/  
Python 3 https://docs.python.org/3/  


Features
--------

1. For database query  
LangChain will generate SQL query from text and magically query the data.  

2. For plugin query  
Query as same as ChatGPT plugins.  
Refer: https://python.langchain.com/docs/integrations/tools/chatgpt_plugins  


Execute
-------

Tested with Golang 1.21 and python3.11.2  

* Docker  
`docker compose up --build -d`  

* Manually  
`go run server.go`  


.env
----

`END_POINT`  
The endpoint of the go server.  

`OPENAI_API_KEY`  
Get from OpenAI 

`MODEL_NAME`  
`TEMPERATURE`  
Refer OpenAI API document.  

`DB_USERNAME`  
`DB_PASSWORD`  
`DB_HOST`  
`DB_PORT`  
`DB_DATABASE`  
DB connection settings

`DEBUG`  
To debug the Python code from the API response.  

`USE_VERBOSE`  
The LangChain response verbose details.  

`DEFAULT_QUERY_ENGINE`  
The default query engine can be `text`, `db`, or `browsing`.  


Interface
---------

GET `/query?input=query_text`  
The Simple AI Chat will.  

GET `/query_db?input=query_text`  
For the database query.  

GET `/query_browsing?input=query_text`  
For the web browsing query.  

GET `/query_text?input=query_text`  
For the text file query.  
