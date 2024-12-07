package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Student struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

type Request struct {
	StudentID string `json:"student_id"`
}

var redisClient *redis.Client

func initRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetStudentInfo(ctx context.Context, studentId string) (*Student, error) {
	stu := &Student{}

	data, err := redisClient.HGetAll(ctx, studentId).Result()
	if err != nil {
		return nil, err
	}

	for field, value := range data {
		switch field {
		case "Name":
			stu.Name = value
		case "Age":
			stu.Age, err = strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
		case "Height":
			stu.Height, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
		}
	}

	return stu, nil
}

func GetName(c *gin.Context) {
	studentID := c.Query("student_id")
	if studentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
		return
	}

	stu, err := GetStudentInfo(c, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch student info", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": stu.Name})
}

func GetAge(c *gin.Context) {
	studentID := c.PostForm("student_id")
	if studentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
		return
	}

	stu, err := GetStudentInfo(c, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch student info", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"age": stu.Age})

}

func GetHeight(c *gin.Context) {
	var req Request

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stu, err := GetStudentInfo(c, req.StudentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch student info", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"height": stu.Height})

}

func main() {
	initRedisClient()

	r := gin.Default()
	r.GET("/get_name", GetName)
	r.POST("/get_age", GetAge)
	r.POST("/get_height", GetHeight)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

//redis-cli
//HMSET 1 Name "John Doe" Age "20" Height "175.5"
//go run main.go
