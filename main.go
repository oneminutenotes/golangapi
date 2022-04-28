package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Marks   int    `json:"marks"`
}

func main() {
	http.HandleFunc("/", handler) //each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	s := &Student{Id: 1, Name: "Ajinkya", Address: "Pune", Marks: 100}
	fmt.Fprintf(w, "Student = %v\n", myConv(s))
}

func myConv(s *Student) string {
	b, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}
	return string(b)
}
