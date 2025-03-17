package server

import "github.com/cpyun/gyopls-core/examples/internal/config"

func Exectue() error {
	config.Setup()
	return nil
}
