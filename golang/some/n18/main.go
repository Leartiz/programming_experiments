package main

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     1,
		Compress:   false,
	})

	log.SetPrefix("app ")

	ipsum := "ipsum"
	for i := 0; i < 100000; i++ {
		log.Printf("Lorem %s", ipsum)
	}
}
