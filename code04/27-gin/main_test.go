package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetName(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/get_name?student_id=1")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(bytes))

}

func TestGetAge(t *testing.T) {
	resp, err := http.PostForm("http://localhost:8080/get_age", url.Values{"student_id": []string{"1"}})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(bytes))
}

func TestGetHeight(t *testing.T) {

	client := &http.Client{}
	reader := strings.NewReader(`{"student_id": "1"}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/get_height", reader)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(bytes))
}

//go test -v ./ -run TestGetName -count=1
//go test -v ./ -run TestGetAge -count=1
//go test -v ./ -run TestGetHeight -count=1
