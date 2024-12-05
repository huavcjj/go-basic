package main

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
)

type Student struct {
	Name   string
	Age    int
	Gender bool
}

type Class struct {
	Id       string
	Students []Student
}

var (
	s = Student{
		Name:   "John",
		Age:    25,
		Gender: true}
	c = Class{
		Id:       "C001",
		Students: []Student{s, s, s},
	}
)

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(c)

		var c2 Class
		json.Unmarshal(bytes, &c2)

	}
}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := sonic.Marshal(c)

		var c2 Class
		sonic.Unmarshal(bytes, &c2)

	}
}
