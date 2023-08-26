
Simple AI Core
==============


Automatically connect to the database and execute SQL query.  


Dependencies
------------

GoLang https://go.dev/doc/  
python 3 https://docs.python.org/3/  

* Standard Version  
go1.21  
python3.11.2  


Execute
-------

* Docker  
`docker compose up --build -d`

* Manually  
`go run server.go`  


Interface
---------

GET http://localhost:8080/query?q=some_query_text  
