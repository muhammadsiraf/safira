package main

import (
	"log"
	"net/http"
	"os"

	"github.com/muhammadfaris/safira/handlers"
)

func main() {
	l := log.New(os.Stdout, "safira-api", log.LstdFlags)
	newHello := handlers.NewHello(l)
	newBye := handlers.NewBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", newHello)
	sm.Handle("/goodbye", newBye)

	http.ListenAndServe(":9090", sm)
}
