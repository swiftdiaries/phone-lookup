package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/swiftdiaries/phone-lookup/search/store"
	"github.com/swiftdiaries/phone-lookup/search/util"
)

const (
	htmlpath = "frontend"
	filename = "index.html"
)

var (
	respponsePerson *util.Person
	port            = os.Getenv("PORT_1")
	redisServer     = os.Getenv("REDIS_URL")
)

func main() {

	redisPassword := flag.String("redis-password", "", "Specify the redis server password")

	flag.Parse()

	if redisServer != "" && redisPassword != nil {
		store.Pool = store.NewPool(redisServer, *redisPassword)
	}

	fileServerIndex := http.FileServer(http.Dir("./frontend/index/"))
	http.Handle("/", fileServerIndex)
	http.HandleFunc("/result", output)
	fileServerResult := http.FileServer(http.Dir("./frontend/result/"))
	http.Handle("/display", fileServerResult)
	go open("http://localhost:" + port + "/")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func output(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form["username"], r.Form["phonenumber"])
		phonenumberURL := "http://localhost:4040/phonenumber/" + r.Form["phonenumber"][0] + "/username/" + r.Form["username"][0]
		response, err := http.Get(phonenumberURL)
		if err != nil {
			fmt.Printf("Error in http.get for response: %s", err)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("%s", err)
			}
			fmt.Printf("%s\n", string(contents))
			json.Unmarshal(contents, &respponsePerson)
			fmt.Printf("%s \n", respponsePerson.Address)
			http.Redirect(w, r, "http://localhost:"+port+"/result", http.StatusSeeOther)
		}
	} else {
		t, _ := template.ParseFiles("./frontend/result/result.html")
		fmt.Printf("%s \n", respponsePerson.Address)
		if respponsePerson.Address == "" {
			respponsePerson.Address = "Name Does Not Match Records"
		}
		t.Execute(w, respponsePerson)
	}
}

func open(url string) error {

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
		//args = []string{"-a", "'Google Chrome'"}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	fmt.Println(cmd, args)
	return exec.Command(cmd, args...).Start()

}
