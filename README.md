
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

| Data Type                           | Engine | Supported |
|-------------------------------------|--------|-----------|
| Text (.txt)                         | text   | Yes       |
| Relational database                 | db     | Yes       |
| Plugin (.well-known/ai-plugin.js)   | plugin | Yes       |
| PDF (.pdf)                          | pdf    | No        |
| EPUB (.epub)                        | epub   | No        |
| CSV (.csv)                          | csv    | No        |
| Unstructured                        | ust    | No        |
| Web browsing                        | url    | No        |
| Image (.jpg, .png, etc.)            | img    | No        |
| Whisper audio (.wav, .mp3, etc.)    | audio  | No        |

Welcom to [join](https://github.com/gcc3) the development.  


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


Response
--------

```json
{
    "result": "Sample response text."
}
```


.env
----

PORT  
The port of the go server.  

PYTHON_PATH  
Specify the python executable path.  
Get from `which python`.  

NODE  
The node name.  

ID  
The node ID, use a number, usual I use unix time.  

OPENAI_API_KEY  
Get from OpenAI 

MODEL_NAME  
TEMPERATURE  
Refer OpenAI API document.  

DEFAULT_QUERY_ENGINE  
Engine name refer "Supported Data Sources".  

CUSTOM_DATA_ONLY  
Custom data only for text query.  
Outside the custom data AI will say "I don't know"  

DEBUG  
To debug the Python code from the API response.  

USE_VERBOSE  
The LangChain response verbose details.  

DB_USERNAME  
DB_PASSWORD  
DB_HOST  
DB_PORT  
DB_DATABASE  
For relational database node.  
DB connection settings.

PLUGIN_URL  
For plugin node.
.well-known/ai-plugin.json file URL.  
Refer OpenAI plagin document.  
