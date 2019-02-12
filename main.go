package main

import (
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/ChrisTheShark/golang-mongo-api/repository"

	"github.com/ChrisTheShark/golang-mongo-api/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	ur := repository.NewUserRepository(getSession())
	uc := controllers.NewUserController(ur)

	r.GET("/users", uc.GetUsers)
	r.POST("/users", uc.AddUser)
	r.GET("/users/:id", uc.GetUserByID)
	r.DELETE("/users/:id", uc.DeleteUser)
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		log.Fatalf("unable to connect to mongo due to: %v", err)
	}
	return session
}
