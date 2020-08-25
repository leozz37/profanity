package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"github.com/gorilla/mux"
	"time"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func NewRandomXingamento() string {
	return fmt.Sprintf("%s %s",
		getRandomFromFile("data/first.txt"),
		getRandomFromFile("data/last.txt"),)
}

func getRandomFromFile(file string) string {
	rand.Seed(time.Now().UnixNano())
	all, _ := ioutil.ReadFile(file)
	list := bytes.Split(all, []byte("\n"))
	return string(list[rand.Intn(len(list))])
}

func GetWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, NewRandomXingamento())
}

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")
	router.HandleFunc("/", GetWord).Methods("GET")
	log.Fatal(http.ListenAndServe(":" + port, router))
}