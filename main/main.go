package main


import (
	"log"

	"github.com/wkmx/operation"
)

func main() {
	var (
		err error
	)
	if err = operation.InitApiServer(); err != nil {
		goto ERR
	}

ERR:
	if err != nil {
		log.Println(err)
	}

}