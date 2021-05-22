package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name    string `bson:"name,omitempty" json:"name,omitempty"`
	Mail    string `bson:"mail,omitempty" json:"mail,omitempty"`
	Message string `bson:"message,omitempty" json:"message,omitempty"`
}

func main() {

	http.HandleFunc("/user", user)
	http.ListenAndServe(":8080", nil)

}

func user(response http.ResponseWriter, request *http.Request) {

	setupCorsResponse(&response, request)
	if (*request).Method == "OPTIONS" {
		return
	}

	if request.Method != "POST" {
		http.Redirect(response, request, "/user", http.StatusSeeOther)
		return
	}

	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	databaseURL := "mongodb://127.0.0.1:27017"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.Fatal("ERROR : " + err.Error())
	}
	//fmt.Println(p)
	fmt.Println("Connected to MongoDB!")

	name := user.Name
	mail := user.Mail
	message := user.Message

	collection := client.Database("go-lang-app").Collection("users")

	userData := User{
		name,
		mail,
		message,
	}

	insertResult, err := collection.InsertOne(context.TODO(), userData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted an User With This Id: ", insertResult.InsertedID)
	response.WriteHeader(200)
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

}
