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

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, os.Getenv("NODE"))
}

type Response struct {
	Message string `json:"result"`
	Error   string `json:"error,omitempty"`
}

// default query handler
func generateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Querying engine: " + os.Getenv("DEFAULT_QUERY_ENGINE") + "...")  // text, db, browsing

	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Input query parameter is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Input: ", input, "\n")
	cmd := exec.Command(os.Getenv("PYTHON_PATH"), "query.py", input)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")  // avoid encoding error

	// Debug python
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
	if err != nil {
		http.Error(w, "Error: " + err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Raw Output: ", string(output))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(output)); err != nil {
        // Handle the error if there's one
        http.Error(w, "Internal Server Error while writing response", http.StatusInternalServerError)
    }
	
	// response := Response{}
	// if err := json.Unmarshal(output, &response); err != nil {
	// 	http.Error(w, "Error unmarshaling output: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// resultString := strings.TrimSpace(response.Message)

	// var result interface{}
	// if err := json.Unmarshal([]byte(resultString), &result); err == nil {
	// 	fmt.Println("Format: JSON")

	// 	// If the inner JSON is successfully parsed, marshal it back to a JSON string
	// 	resultJSON, err := json.Marshal(result)
	// 	if err != nil {
	// 		http.Error(w, "Error marshaling result: "+err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	response.Message = string(resultJSON)
	// } else {
	// 	fmt.Println("Format: text")

	// 	// If the inner JSON is not successfully parsed, just use the original string
	// 	response.Message = "\"" + response.Message + "\""
	// }

	// sendJSONResponse(w, response, http.StatusOK)
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
	r.HandleFunc("/generate", generateHandler).Methods("GET")

	port := os.Getenv("PORT")
	fmt.Println("Server started on port " + port + "\n")
	http.ListenAndServe(":" + port, r)  // for windows use 127.0.0.1:port
}
