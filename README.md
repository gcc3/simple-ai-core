
Simple AI Node
==============


Example node of distributed AI links (Node AI).  
Refer [simple-ai-chat](https://github.com/gcc3/simple-ai-chat).  


Dependencies
------------

GoLang https://go.dev/doc/  
LangChain https://python.langchain.com/docs/  


Features
--------

1. Text query  
Query for `data.txt` file.  

2. Database query  
LangChain will generate SQL query from text and magically query the data.  

3. Plugin query  
Query as same as ChatGPT plugins.  
Refer: https://python.langchain.com/docs/integrations/tools/chatgpt_plugins  


Setup
-----

Install python dependencies.  
`pip install -r requirements.txt`

Sometimes need to update the packages:  
`pip install --upgrade openai`  
`pip install --upgrade langchain`   
`pip install --upgrade -r requirements.txt`  


Execute
-------

Tested with Golang 1.21 and Python 3.11.2  

* Docker  
`docker compose up --build -d`  

* Manually  
`go run serve.go`  


API
---

GET `/generate?input=input_text`  


Response
--------

```json
{
    "result": "Sample response text."
}
```

Or

```json
{
    "result": {
        "text": "Sample response text.",
        "image": "some_imeage_url",
    }
}
```


Query Engines
-------------

1. text  
Simple text query base on langchain.  
If use query engine `text`, Create `data.txt` to include the data.  

* Supported Data Sources  

| Data Type                             | Engine | Supported   |
|---------------------------------------|--------|-------------|
| Text (.txt)                           | text   | Yes         |
| Relational database                   | db     | Yes         |
| Plugin (.well-known/ai-plugin.js)     | plugin | Yes         |
| PDF (.pdf)                            | pdf    | No          |
| Amazon Textract PDF OCR (.pdf, .jpeg) | ocr    | No          |
| EPUB (.epub)                          | epub   | No          |
| CSV (.csv)                            | csv    | No          |
| Unstructured                          | ust    | No          |
| Web browsing                          | url    | No          |
| Image (.jpg, .png, etc.)              | img    | No          |
| Whisper audio (.wav, .mp3, etc.)      | audio  | No          |


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

QUERY_ENGINE  
Engine name refer "Supported Data Sources".  

CUSTOM_DATA_ONLY  
Custom data only for text query.  
Outside the custom data AI will say "I don't know"  

DEBUG  
To debug the Python code from the API response.  

USE_VERBOSE  
The LangChain response verbose details.  

MYSQL_DB_USERNAME  
MYSQL_DB_PASSWORD  
MYSQL_DB_HOST  
MYSQL_DB_PORT  
MYSQL_DB_DATABASE  
For relational database node.  
DB connection settings.

PLUGIN_URL  
For plugin node.
.well-known/ai-plugin.json file URL.  
Refer OpenAI plagin document.  
