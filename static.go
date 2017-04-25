package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var dir string
	var e error
	port := flag.String("port", "4000", "Port to Serve Directory On")
	path := flag.String("path", "", "Path to Serve.")

	flag.Parse()

	if *path == "" {
		dir, e = os.Getwd()
		if e != nil {
			fmt.Printf("Exiting Could not Get Working Dir Error: %s \n", e)
			os.Exit(1)
		}
	} else {
		dir = *path
	}

	fmt.Printf("Starting Static Server at port: %s for directory: %s \n", *port, dir)
	err := http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))

	if err != nil {
		fmt.Println("Exiting Could Not Get Server Started:", err)
		os.Exit(1)
	}
}
