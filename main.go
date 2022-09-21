package main

import "log"

func main() {
	_, err := NewRedisClient()
	if err != nil {
		log.Panic(err)
	}
}
