package main

import (
	"context"
	"fmt"
	student_service "grpc/grpc/idl/my_proto"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpc(t *testing.T) {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc.NewClient error: %v", err)
	}
	defer conn.Close()

	client := student_service.NewStudentServiceClient(conn)
	resp, err := client.GetStudentInfo(context.Background(), &student_service.Request{StudentId: "1"})
	if err != nil {
		t.Fatalf("client.GetStudentInfo error: %v", err)
	}
	fmt.Printf("name: %s, age: %d, height: %f\n", resp.Name, resp.Age, resp.Height)
}

// go test -v ./grpc -run TestGrpc -count=1
