package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Starting server")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	rep := DynamoRepo{DB: dynamodb.New(sess)}

	controller := new(Controller)
	controller.DB = rep

	http.HandleFunc("/mutant/", controller.VerifyDNA)
	http.HandleFunc("/stats", controller.GetStats)
	log.Fatal(http.ListenAndServe(":8000", nil))

}