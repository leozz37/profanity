package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
	"time"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Xingamento struct {
	Value string `json:"xingamento"`
}

func NewRandomXingamento() Xingamento {
	return Xingamento{
		Value: fmt.Sprintf("%s %s",
			getRandomFromFile("data/first.txt"),
			getRandomFromFile("data/last.txt"),
		),
	}
}

func getRandomFromFile(file string) string {
	rand.Seed(time.Now().UnixNano())
	all, _ := ioutil.ReadFile(file)
	list := bytes.Split(all, []byte("\n"))
	return string(list[rand.Intn(len(list))])
}

func GetWord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(NewRandomXingamento())
}

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")
	router.HandleFunc("/", GetWord).Methods("GET")
	log.Fatal(http.ListenAndServe(":" + port, router))
}