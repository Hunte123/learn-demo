package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"websocket/socket"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found 11", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	http.ServeFile(w, r, pwd+"/home.html")
}

func main() {
	flag.Parse()
	hub := socket.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWs(hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
