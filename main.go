package main

import (
	"flag"
	"log"
	"telegramBot/clients/telegram"
)

func main() {
	tgClient := telegram.New(mustTgHost(), mustToken())
	log.Println(tgClient)

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

func mustTgHost() string {
	var tgHost *string = flag.String(
		"tg-bot-host",
		"",
		"telegram bots host",
	)
	flag.Parse()

	if *tgHost == "" {
		log.Fatal("Tg host is not specified")
	}

	return *tgHost
}
