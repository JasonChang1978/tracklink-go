package main

/*
社交工程-可監控連結設計方案
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ClickEvent struct {
	ID string `json:"id"`
}

func main() {
	http.HandleFunc("/track", trackHandler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func logToGoogleSheets(id string) {
	url := "https://script.google.com/macros/s/AKfycbxZqyfNdh38wjquMFxd2yBQhJ4_63XY-ZcZC1GhbiVTLFvUvoF-1okg8ssVsEKd3QJm9w/exec" // 換成你的網址

	payload := map[string]string{"id": id}
	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error logging to Google Sheets:", err)
		return
	}
	defer resp.Body.Close()
}

func trackHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	if id == "12345" {
		fmt.Println("ID: 12345 is test data, ignore")
	} else {
		logToGoogleSheets(id)
	}

	// 顯示訊息
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>您已完成社交工程演練!!</p><p style=\"color: red;\">請保持沉默，勿告知其他同仁</p>")
}
