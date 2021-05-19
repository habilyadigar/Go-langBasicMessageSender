package main

import (
	"context"
	"fmt"
	"html/template"
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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/user", user)
	http.ListenAndServe(":8080", nil)
}

func index(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "index.gohtml", nil)
}

func user(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Redirect(response, request, "/", http.StatusSeeOther)
		return
	}
	databaseURL := "mongodb://127.0.0.1:27017"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.Fatal("ERROR : " + err.Error())
	}
	fmt.Println("Connected to MongoDB!")
	fname := request.FormValue("name")
	fmail := request.FormValue("mail")
	fmessage := request.FormValue("message")

	collection := client.Database("go-lang-app").Collection("users")
	anUser := User{
		fname,
		fmail,
		fmessage,
	}
	insertResult, err := collection.InsertOne(context.TODO(), anUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted an User With This Id: ", insertResult.InsertedID)
	tpl.ExecuteTemplate(response, "user.gohtml", anUser)
}
