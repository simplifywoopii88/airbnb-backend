package utils

import "log"

func HandleErr(err error) {
	log.Fatal(err)
}
