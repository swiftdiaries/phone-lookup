package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/swiftdiaries/phone-lookup/search/query"
	"github.com/swiftdiaries/phone-lookup/search/store"
	"github.com/swiftdiaries/phone-lookup/search/util"
)

var (
	port        = os.Getenv("PORT")
	redisServer = os.Getenv("REDIS_URL")
)

func main() {
	redisPassword := flag.String("redis-password", "", "Specify the redis server password")

	flag.Parse()

	if redisServer != "" && redisPassword != nil {
		store.Pool = store.NewPool(redisServer, *redisPassword)
	}

	router := mux.NewRouter()
	router.HandleFunc("/phonenumber/{phonenumber}/username/{username}", GetPhoneNumberEndPoint)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

//GetPhoneNumberEndPoint calls the function to check if phone number
//and username exists and returns the boolean
func GetPhoneNumberEndPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phonenumber := vars["phonenumber"]
	username := vars["username"]

	person := &util.Person{
		Name:        username,
		Phonenumber: phonenumber,
		Address:     "",
	}

	person = query.CheckAndFetch(person)
	if person == nil {
		fmt.Fprintf(w, "nil")
	}
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("Error in Marshalling JSON: %s", err)
	}

	fmt.Println(string(data))
	fmt.Fprintf(w, "%s", data)

}
