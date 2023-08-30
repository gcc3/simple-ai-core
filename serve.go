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
	fmt.Fprintf(w, "Hi, this is simple ai core server.")
}

// default query handler
func handleQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Querying with database...")

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

// database query handler
func handleQueryDb(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Querying with database...")

	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Input query parameter is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Input: ", input, "\n")
	cmd := exec.Command("python", "query_db.py", input)
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

// browsing plugin query handler
func handleQueryBrowsing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Querying with plugin...")

	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Input query parameter is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Input: ", input, "\n")
	cmd := exec.Command("python", "query_browsing.py", input)
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

// text query handler
func handleQueryText(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Querying with plugin...")

	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Input query parameter is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Input: ", input, "\n")
	cmd := exec.Command("python", "query_text.py", input)
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
	r.HandleFunc("/api/query", handleQuery).Methods("GET")
    r.HandleFunc("/api/query_db", handleQueryDb).Methods("GET")
    r.HandleFunc("/api/query_plugin", handleQueryBrowsing).Methods("GET")
    r.HandleFunc("/api/query_text", handleQueryText).Methods("GET")

	endpoint := os.Getenv("END_POINT")
	fmt.Println("Server started on " + endpoint + "\n")
	http.ListenAndServe(endpoint, r)  // for windows use 127.0.0.1:8080
}
