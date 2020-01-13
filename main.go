package main

import (
	"github.com/favclip/ucon/v3"
)

func main() {
	ucon.Orthodox()

	Setup()

	ucon.ListenAndServe(":80")
}
