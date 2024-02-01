package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "Method not allowed", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello from the server")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	// parsing the form
	if err:=r.ParseForm(); err!=nil{
		fmt.Fprintf(w, "ParseForm() error %v\n", err)
		return
	}

	name:=r.FormValue("name")
	address:=r.FormValue("address")

	fmt.Fprintf(w, "Details submitted\n")
	fmt.Fprintf(w, "Hi %s\n", name)
	fmt.Fprintf(w, "Happy to see you all the way from %s\n", address)

}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting Server at http://localhost:8080")

	if err:=http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal(err)
	}
}
