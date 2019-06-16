package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	var port = flag.String("port", "3141", "The port which httpi should listen on")
	flag.Parse()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, (strconv.FormatFloat(math.Pi, 'f', 10, 64)))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:"+*port, nil))
}
