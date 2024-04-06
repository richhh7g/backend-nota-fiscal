package main

import (
	"fmt"

	"github.com/richhh7g/mm-api-nfe/pkg/environment"
)

func init() {
	err := environment.NewEnvLoader(nil)
	if err != nil {
		panic(err)
	}

}

func main() {
	fmt.Println("Hello, World! MM")
}
