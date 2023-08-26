package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func handleQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	data := map[string]string{
		"query": query,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}

	cmd := exec.Command("python", "query_db.py")
	cmd.Stdin = bytes.NewReader(jsonData)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		http.Error(w, "Query script execution failed", http.StatusInternalServerError)
		return
	}

	result := out.String()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"result": result,
	})
}

func main() {
	http.HandleFunc("/query", handleQuery)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}