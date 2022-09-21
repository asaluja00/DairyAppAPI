package main

import (
	"errors"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type entry struct {
	ID      string `json "id"`
	Date    string `json "date"`
	Title   string `json "title"`
	Content string `json "content"`
	Time    string `json "time"`
}
type user struct {
	ID           string `json "id"`
	SecretCode   string `json secretcode"`
	Name         string `json "name"`
	Email        string `json "email"`
	DOB          string `json "DOB"`
	dairyentries entry  `json "dairyentries"`
}

var entries = []entry{}
var users = []user{}

func login(id string) (*user, error) {
	for i, t := range users {
		if t.SecretCode == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("Wrong Code")

}
func getuser(context *gin.Context) {
	id := context.Param("secretcode")
	user, err := login(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Not Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, user)
}
func register(context *gin.Context) {
	var newuser user
	if err := context.BindJSON(&newuser); err != nil {
		return
	}
	min := 1
	max := 100
	var id = (rand.Intn(max-min) + min)
	var temp = (rand.Intn(999999) + 100000)
	newuser.SecretCode = strconv.Itoa(temp)
	newuser.ID = strconv.Itoa(id)
	users = append(users, newuser)
	context.IndentedJSON(http.StatusCreated, newuser)

}
func addEntry(context *gin.Context) {
	var newEntry entry
	if err := context.BindJSON(&newEntry); err != nil {
		return
	}
	min := 1
	max := 100
	var id = (rand.Intn(max-min) + min)
	newEntry.ID = strconv.Itoa(id)
	entries = append(entries, newEntry)
	context.IndentedJSON(http.StatusCreated, newEntry)
}
func deleteentry(context *gin.Context) {
	date := context.Param("date")
	err := false
	for i, t := range entries {
		if t.Date == date {
			entries[i] = entry{}
			err = true
		}
	}
	if !err {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Not Found"})
		return
	}
}
func updateEntry(context *gin.Context) {

}
func showEntry(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, entries)
}

func main() {
	router := gin.Default()
	router.GET("/showEntry", showEntry)
	router.GET("/Login/:secretcode", getuser)
	router.POST("/register", register)
	router.POST("/addEntry", addEntry)
	router.POST("/deleteentry/:date", deleteentry)

	router.Run("localhost:9090")
}
