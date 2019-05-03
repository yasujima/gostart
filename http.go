package main

import (
	"encoding/json"
	"log"
	"os"
	"fmt"
	"net/http"
	"strconv"
	"io/ioutil"
)

type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func IndexHandler(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "hello http world")
}

func PersonHandler(w http.ResponseWriter,
	r *http.Request) {
	defer r.Body.Close()

	if r.Method == "POST" {
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusCreated)
	} else if r.Method == "GET" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
		filename := fmt.Sprintf("%d.txt", id)
		message , err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(message))
	}
}




func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndServe(":8000", nil)
}
