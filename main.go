package main

/*
社交工程-可監控連結設計方案
*/

import (
	"fmt"
	"net/http"
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

	now := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("ID: %s, Timestamp: %s", id, now)
	fmt.Println(logEntry) // 輸出到標準輸出

	// 顯示訊息
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>您已完成帳號確認</p>")
}
