package main

import (
	"os"
	"time"

	"go-with-tests/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
