package main

import (
	"log"

	"gihtub.com/wkmx/operator/config"
)

func main() {
	var (
		err error
	)

	if err = config.InitApiServer(); err != nil {
		goto ERR
	}

ERR:
	if err != nil {
		log.Println(err)
	}

}
