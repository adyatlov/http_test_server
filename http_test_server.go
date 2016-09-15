package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var delay int
	buff := bytes.NewBufferString("")
	buff.WriteString("New request: " + r.URL.String() + "\n")
	// Print env vars
	envVars := os.Environ()
	sort.Strings(envVars)
	buff.WriteString("\n")
	buff.WriteString("Env Vars\n")
	buff.WriteString("======\n")
	for _, envVar := range envVars {
		buff.WriteString(envVar + "\n")
	}
	buff.WriteString("======\n")
	buff.WriteString("\n")
	// Delay
	delay, err := strconv.Atoi(r.URL.Query().Get("delay"))
	if err != nil {
		buff.WriteString("Delay is wrong or not specified.\n")
	}
	buff.WriteString("Delay: " + strconv.Itoa(delay) + "\n")
	buff.WriteString("\n")
	// Headeer
	var keys []string
	for key := range r.Header {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	buff.WriteString("Header\n")
	buff.WriteString("======\n")
	buff.WriteString(fmt.Sprintf("%v: %v\n", "Host", r.Host))
	for _, key := range keys {
		buff.WriteString(fmt.Sprintf("%v: %v\n", key, r.Header[key]))
	}
	buff.WriteString("======\n")
	log.Println(buff)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Fprint(w, buff)
}

func main() {
	http.HandleFunc("/", handler)
	if len(os.Args) != 2 {
		log.Fatalf("Usage: sleeps_server <port>\n")
	}
	_, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		log.Fatalln("Port should be a number.")
	}
	log.Printf("Usage example: curl http://localhost:%v/?delay=3", os.Args[1])
	log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}
