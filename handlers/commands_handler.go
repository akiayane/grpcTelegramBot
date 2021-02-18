package handlers

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type commands_handler struct {

	functions_array []pattern
	bot *tgbotapi.BotAPI
	u *tgbotapi.UpdateConfig

}
type pattern struct {

	pattern_name string
	func_to_call func(update tgbotapi.Update, param string)

}

func(h *commands_handler) add_new_function(p pattern){
	h.functions_array = append(h.functions_array, p)
}

func(h *commands_handler) call_function(update tgbotapi.Update,pattern_name string, param string){
	found := false
	for _, value := range h.functions_array {
		if value.pattern_name == pattern_name {
			call := value.func_to_call
			call(update, param)
			found = true
			break
		}
	}
	if !found{
		call:=h.functions_array[0].func_to_call
		call(update, "")
	}
}


func Commands_handler_init(bot *tgbotapi.BotAPI, u *tgbotapi.UpdateConfig, channel *tgbotapi.UpdatesChannel) commands_handler{
	var functions_array []pattern

	handler := commands_handler{functions_array, bot, u}

	//adding functionalities
	//function for unknown input, always should be first
	handler.add_new_function(pattern{"/notfound", func(update tgbotapi.Update, param string) {

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I dont know such command")
		fmt.Println("start called")
		_, _ = bot.Send(msg)

	}})

	//start
	handler.add_new_function(pattern{"/start", func(update tgbotapi.Update, param string) {

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+param)
		fmt.Println("start called")
		_, _ = bot.Send(msg)

	}})



	//gen
	handler.add_new_function(pattern{"/gen", func(update tgbotapi.Update, param string) {

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+param)
		fmt.Println("gen called")
		_, _ = bot.Send(msg)

		/*
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter name plz:")
		fmt.Println("gen called")
		_, _ = bot.Send(msg)


		for update := range *channel {
			if update.Message == nil {
				continue
			}
			if update.Message.Text == "/main"{
				handler.call_function(update,"/start", "")
				break
			}
			name := update.Message.Text
			val := rand.Intn(3)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, name+" "+lohs[val])
			_, _ = bot.Send(msg)

		}*/
	}})

	return handler
}



