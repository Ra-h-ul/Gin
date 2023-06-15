package main

import (
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// creating struct
type Student struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"LastName"`
	Roll_no   int    `json:"roll_no"`
	Class     int    `json:"class"`
	Phone_No  int    `json:"Phone_No"`
}

// declaring struct values
var Students = []Student{
	{FirstName: "Rahul", LastName: "Pandita", Roll_no: 17, Class: 12, Phone_No: 78897},
	{FirstName: "john ", LastName: "Doe", Roll_no: 18, Class: 12, Phone_No: 97887},
	{FirstName: "Db", LastName: "Cooper", Roll_no: 19, Class: 11, Phone_No: 9419},
	{FirstName: "Tony", LastName: "Montana", Roll_no: 21, Class: 11, Phone_No: 2250},
	{FirstName: "papa", LastName: "Franku", Roll_no: 23, Class: 12, Phone_No: 9682},
	{FirstName: "Mah", LastName: "Dank", Roll_no: 32, Class: 12, Phone_No: 1234},
}

// GET request for students
func get_student_data(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, Students)
}

func main() {
	router := gin.Default()
	router.GET("/students", get_student_data)
	router.Run("localhost:8080")
}
