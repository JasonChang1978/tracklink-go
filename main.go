package main

/*
社交工程-可監控連結設計方案
*/

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/track", trackHandler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func trackHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	// 記錄到 logs.csv
	file, err := os.OpenFile("logs.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to write log", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	now := time.Now().Format("2006-01-02 15:04:05")
	writer.Write([]string{id, now})

	// 顯示訊息
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<p>您已完成帳號確認</p>")
}
