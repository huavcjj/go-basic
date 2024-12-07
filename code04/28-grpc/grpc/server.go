package main

import (
	"context"
	"errors"
	"fmt"
	student_service "grpc/grpc/idl/my_proto"
	"net"

	"google.golang.org/grpc"
)

type StudentServer struct {
	student_service.UnimplementedStudentServiceServer
}

func (StudentServer) GetStudentInfo(ctx context.Context, in *student_service.Request) (*student_service.Student, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in f", err)
		}
	}()
	if len(in.StudentId) == 0 {
		return nil, errors.New("student id is empty")
	}
	student := &student_service.Student{
		Name:   "golang",
		Age:    10,
		Height: 180.5,
	}
	return student, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	student_service.RegisterStudentServiceServer(server, &StudentServer{})
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
