package main

import (
	"encoding/json"
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"os"
)

type Response struct {
	Message string `json:"result"`
	Error   string `json:"error,omitempty"`
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Input query parameter is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Input: ", input, "\n")
	cmd := exec.Command("python", "query_db.py", input)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")  // avoid encoding error

	// debug python
	debug := false
	if (debug) {
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

	response.Message = string(output)
	sendJSONResponse(w, response, http.StatusOK)
}

func sendJSONResponse(w http.ResponseWriter, resp Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/query", handleQuery)
	fmt.Println("Server started on :8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
