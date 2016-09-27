package main

import (
	"flag"
	"net/http"
	"strconv"
	"time"
)

var (
	addr = flag.String("addr", ":8080", "listen address")
	key  = flag.String("key", "", "secret key")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", puke)
	http.ListenAndServe(*addr, nil)
}

func puke(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("key") != *key {
		http.Error(w, "invalid key", 401)
		return
	}
	bufferSizeStr := r.FormValue("bufferSize")
	if bufferSizeStr == "" {
		bufferSizeStr = "256000"
	}
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	if err != nil {
		http.Error(w, "invalid bufferSize param", 400)
		return
	}
	durationStr := r.FormValue("duration")
	if durationStr == "" {
		durationStr = "20s"
	}
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		http.Error(w, "invalid duration param", 400)
		return
	}
	var buffer = make([]byte, bufferSize)
	var startTime = time.Now()
	for {
		_, err := w.Write(buffer)
		if err != nil {
			return
		}
		if time.Now().Sub(startTime) > duration {
			return
		}
	}
}
