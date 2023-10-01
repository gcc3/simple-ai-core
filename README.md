
Simple AI Node
==============


Node of distributed AI links.  
Provide database connection, ChatGPT plugin support. etc.  


Dependencies
------------

LangChain https://python.langchain.com/docs/  
GoLang https://go.dev/doc/  


Features
--------

1. For database query  
LangChain will generate SQL query from text and magically query the data.  

2. For plugin query  
Query as same as ChatGPT plugins.  
Refer: https://python.langchain.com/docs/integrations/tools/chatgpt_plugins  


Supported Data Sources
----------------------

| Data Type                        | Supported |
|----------------------------------|-----------|
| Relational database              | Yes       |
| Text (.txt)                      | Yes       |
| PDF (.pdf)                       | No        |
| EPUB (.epub)                     | No        |
| CSV (.csv)                       | No        |
| Unstructured                     | No        |
| Web browsing                     | No        |
| Image (.jpg, .png, etc.)         | No        |
| Whisper audio (.wav, .mp3, etc.) | No        |

Please [join](https://github.com/gcc3) the development.  


Execute
-------

Tested with Golang 1.21 and Python 3.11.2  

* Docker  
`docker compose up --build -d`  

* Manually  
`go run server.go`  


Interface
---------

GET `/query?input=query_text`  


.env
----

`PORT`  
The port of the go server.  

`NODE`  
The node name.  

`ID`  
The node ID, use a number.  

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
