package main

import (
	"fmt"
	"os"

	"github.com/luewell/framework/bootstrap"
)

func init() {
	bootstrap.App()
}

func main() {
	fmt.Println(os.Environ())
}
