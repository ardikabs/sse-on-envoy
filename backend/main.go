package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/alexandrevicenzi/go-sse"
)

func main() {
	opt := &sse.Options{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "https://localhost:8080",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Methods":     "GET, OPTIONS",
			"Access-Control-Allow-Headers":     "Keep-Alive,X-Requested-With,Cache-Control,Content-Type,Last-Event-ID",
		},
	}
	// Create the server.
	s := sse.NewServer(opt)
	defer s.Shutdown()

	// Register with /events endpoint
	http.Handle("/events", s)

	// Dispatch messages
	id := 1
	go func() {
		for {
			joke := getDadJoke()
			s.SendMessage("/events", sse.NewMessage(strconv.Itoa(id), joke, "DAD_JOKE"))
			id++
			time.Sleep(10 * time.Second)
		}
	}()

	http.ListenAndServe(":8080", nil)
}

func getDadJoke() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Add("Accept", "text/plain")
	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	joke := string(body)
	if err != nil {
		fmt.Println(err)
	}
	return joke
}
