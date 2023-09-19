package main

import (
	"encoding/json"
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"result"`
	Error   string `json:"error,omitempty"`
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, os.Getenv("NODE"))
}

// default query handler
func handleQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Querying engine: " + os.Getenv("DEFAULT_QUERY_ENGINE") + "...")  // text, db, browsing

	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Input query parameter is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Input: ", input, "\n")
	cmd := exec.Command("python", "query.py", input)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")  // avoid encoding error

	// debug python
	if (os.Getenv("DEBUG") == "true") {
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			http.Error(w, "Failed to run python script: " + err.Error() + "\n\n" + stderr.String(), http.StatusInternalServerError)
			return
		}
	    fmt.Fprintf(w, "No error")
		return
	}

	output, err := cmd.Output()
	response := Response{}
	if err != nil {
		http.Error(w, "Error: " + err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Output: ", string(output))
	response.Message = string(output)
	sendJSONResponse(w, response, http.StatusOK)
}

func sendJSONResponse(w http.ResponseWriter, resp Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// routes
	r := mux.NewRouter()
	r.HandleFunc("/", infoHandler).Methods("GET")
	r.HandleFunc("/query", handleQuery).Methods("GET")  // query?input=...

	endpoint := os.Getenv("PORT")
	fmt.Println("Server started on port " + endpoint + "\n")
	http.ListenAndServe(":" + endpoint, r)  // for windows use 127.0.0.1:port
}
