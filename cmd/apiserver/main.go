package main

import (
	"github.com/BedivereZero/nekodns/pkg/apiserver"
)

func main() {
	s, _ := apiserver.Default()
	s.Router.Run()
}
