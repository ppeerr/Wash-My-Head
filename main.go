package main

import (
	"flag"
	"log"
)

func main() {
	var token string = mustToken() //tgClient = telegram.New(token)
	log.Println(token)

	//fetcher = fetcher.New()
	//processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	var token *string = flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
