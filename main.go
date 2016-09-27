package main

import (
	"flag"
	"net/http"
	"strconv"
	"time"
)

var (
	addr = flag.String("addr", ":8080", "listen address")
)

const (
	defaultBufferSize = "262144" // 256K
	defaultDuration   = "20s"
)

func main() {
	flag.Parse()

	http.HandleFunc("/", puke)
	http.ListenAndServe(*addr, nil)
}

func puke(w http.ResponseWriter, r *http.Request) {
	bufferSizeStr := r.FormValue("bufferSize")
	if bufferSizeStr == "" {
		bufferSizeStr = defaultBufferSize
	}
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	if err != nil {
		http.Error(w, "invalid bufferSize param", 400)
		return
	}
	durationStr := r.FormValue("duration")
	if durationStr == "" {
		durationStr = defaultDuration
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
