package main

import "os"

func main() {
	path := "../input/yourJson.json"
	if verify(path) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
