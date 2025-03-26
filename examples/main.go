package main

import (
	"os"

	"github.com/cpyun/gyopls-core/examples/cmd/server"
)

func main() {
	if err := server.Exectue(); err != nil {
		os.Exit(-1)
	}
}
