package main

import (
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	soundFile := "ching.wav"
	if len(os.Args) > 1 {
		soundFile = os.Args[1]
	}

	f, err := os.Open(soundFile)
	if err != nil {
		log.Fatal("Error opening file")
	}
	s, format, _ := wav.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	http.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		_, err := res.Write([]byte("OK"))
		if err != nil {
			log.Printf("error writing response to %s: %v", req.RemoteAddr, err)
		}
	})

	http.HandleFunc("/ching", func(res http.ResponseWriter, req *http.Request) {
		log.Printf("received ching from %s", req.RemoteAddr)
		speaker.Play(s)
		err := s.Seek(0)
		if err != nil {
			log.Printf("error resetting streamer: %v", err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusOK)
	})

	log.Println("Listening on Port 8080 for some chings 💰")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
