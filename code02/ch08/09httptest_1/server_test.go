package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGet(t *testing.T) {
	//テストを実行するためのマルチプレクサを生成
	mux := http.NewServeMux()
	//テスト対象のハンドラをマルチプレクサに登録
	mux.HandleFunc("/post/", handleRequest)

	//返されたHTTPレスポンスを取得するためのレコーダを生成
	writer := httptest.NewRecorder()
	//テストしたいハンドラに対してリクエストを生成
	request, _ := http.NewRequest("GET", "/post/1", nil)
	//テスト対象のハンドラにリクエストを送信
	mux.ServeHTTP(writer, request)

	//レスポンスコードが200であることを確認
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
