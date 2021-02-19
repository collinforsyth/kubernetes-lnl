package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// We need to update configuration outside of our binary
	address := os.Getenv("FOO_ADDRESS")

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	client := &http.Client{}

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				req, err := http.NewRequest(http.MethodGet, address, nil)
				if err != nil {
					log.Fatal(err)
				}
				resp, err := client.Do(req)
				if err != nil {
					log.Print(err)
					break
				}
				bytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Print(err)
				}
				log.Print(string(bytes))
			}
		}
	}()

	http.ListenAndServe(":8080", r)

	// is never hit but we aren't implementing graceful shutdown are we?
	done <- true
}
