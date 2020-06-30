package main

import (
	"bytes"
	"github.com/bots/primerBotTelegram/pokemon"
	"github.com/bots/primerBotTelegram/pokemon/templates"
	"log"
	"text/template"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func newTemplateMenAll(pokes pokemon.PokemonList) string{
	text,err:= template.New("templateListPokemon").Parse(templates.TemplateListPokemon)
	if err != nil{
		panic(err)
	}
	var textString bytes.Buffer
	error := text.Execute(&textString, pokes)
	if error != nil{
		panic(error)
	}
	return textString.String()
}

func newTemplateMenOne(poke pokemon.Pokemon) string{
	text,err:= template.New("templateOnePokemon").Parse(templates.TemplateOnePokemon)
	if err != nil{
		panic(err)
	}
	var textString bytes.Buffer
	error := text.Execute(&textString, poke)
	if error != nil{
		panic(error)
	}
	return textString.String()
}

func main() {
	bot, err := tgbotapi.NewBotAPI()
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,"")
			switch update.Message.Command() {
			case "pokemons":
				pokemons := pokemon.PokemonAll()
				msg.Text = newTemplateMenAll(pokemons)
			default:
				poke := pokemon.PokemonOne(update.Message.Command())
				msg.Text = newTemplateMenOne(poke)
			}
			bot.Send(msg)
		}

	}
}