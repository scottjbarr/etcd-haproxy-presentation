package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-etcd/etcd"
)

type profile struct {
	Name    string   `json:"name"`
	Hobbies []string `json:"hobbies"`
}

var (
	hostname string
	port     int
	etcdPath = "profile_service"
	client   *etcd.Client
)

func init() {
	var err error
	hostname, err = os.Hostname()

	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to get hostname : %v", err))
	}
}

func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func sendRegistrationRequest() {
	client.Set(
		fmt.Sprintf("%v/services/%v:%v", etcdPath, hostname, port),
		fmt.Sprintf("%v:%v", hostname, port),
		60)
}

func register() {
	client = etcd.NewClient([]string{"http://127.0.0.1:4001"})

	// sleep briefly to let the server bind to a port in the range
	time.Sleep(time.Duration(2) * time.Second)
	log.Printf("Listening on %v", port)

	for {
		sendRegistrationRequest()
		time.Sleep(time.Duration(30) * time.Second)
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", indexHandler)

	go func() {
		register()
	}()

	// the range of ports we will try to bind to
	minPort := 8100
	maxPort := 8102

	// loop over the range until we get a port
	for port = minPort; port <= maxPort; port++ {
		http.ListenAndServe(fmt.Sprintf(":%v", port),
			logger(http.DefaultServeMux))
	}

	// we failed to bind to a port
	log.Printf("Failed to bind to port. Exiting")
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	message := profile{
		Name:    "Scott Barr",
		Hobbies: []string{"golang guy", "dad", "husband"},
	}

	js, err := json.Marshal(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
