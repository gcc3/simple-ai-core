
Simple AI Core
==============


Provide database connection and ChatGPT plugin support with LangChain.  


Dependencies
------------

GoLang https://go.dev/doc/  
python 3 https://docs.python.org/3/  


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


Interface
---------

GET /query_db?q=some_query_text  
GET /query_plugin?q=some_query_text  
