package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

var (
	addr            = flag.String("addr", ":8080", "listen address")
	bufferSize      = flag.Int("buffer-size-kb", 256, "write buffer size in KB")
	defaultDuration = flag.Duration("default-duration", 20*time.Second, "response is generated for this duration")
)

var buffer []byte

func main() {
	flag.Parse()
	buffer = make([]byte, *bufferSize)

	http.HandleFunc("/", puke)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}

func puke(w http.ResponseWriter, r *http.Request) {
	var err error
	var duration = *defaultDuration
	durationStr := r.FormValue("duration")
	if durationStr != "" {
		duration, err = time.ParseDuration(durationStr)
		if err != nil {
			http.Error(w, "invalid duration param", 400)
			return
		}
	}
	var startTime = time.Now()
	for {
		_, err := w.Write(buffer)
		if err != nil {
			return
		}
		if time.Since(startTime) > duration {
			return
		}
	}
}
