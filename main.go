package main

/*
社交工程-可監控連結設計方案
*/

import (
	"bytes"
	"encoding/json"
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

	// 準備要發送的資料
	postData := map[string]string{
		"id":        id,
		"timestamp": now,
	}
	jsonData, err := json.Marshal(postData)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Google Apps Script 的網路應用程式 URL (請替換為您的實際 URL)
	appsScriptURL := "https://script.google.com/macros/s/AKfycbxcJw0LmJtWjbEEs-UAs5ns9na_vdDsygC_E-nB3vHtbgZfkV6EOKKGqXNRu5_K3DSgnA/exec"

	// 發送 HTTP POST 請求
	resp, err := http.Post(appsScriptURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request to Apps Script:", err)
		http.Error(w, "Error logging data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Data sent to Google Sheets. Status:", resp.Status)

	// 顯示訊息
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>您已完成帳號確認</p>")
}
