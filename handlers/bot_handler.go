package handlers

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"strings"
)

type bot_handler struct {

	command_pool []command_pattern
	log_handler
	commands_handler

}

type command_pattern struct{

	command_name string
	optional_param string //"" if nil
	required_handler string

}

func(h *bot_handler) add_command_pattern(command_pattern command_pattern){
	h.command_pool = append(h.command_pool, command_pattern)
}

func(h *bot_handler) separate_param(str string) (string,string) {
	if strings.Contains(str, " "){
		space_index := strings.Index(str," ")
		command := str[0:space_index]
		param := str[space_index:]
		return strings.Trim(command," "), strings.Trim(param," ")

	}
	return strings.Trim(str," "), ""
}


func(h *bot_handler) call_function(command_name string, command_param string, update tgbotapi.Update){
	found := false
	for _, value := range h.command_pool {
		if value.command_name == command_name {
			fmt.Println(value.required_handler+"|"+value.command_name)
			switch value.required_handler {
			case "commands_handler":
				h.commands_handler.call_function(update, command_name, command_param)
				found = true
				break
			case "log_handler":
				h.log_handler.call_function(update, command_name, command_param)
				found = true
				break
			default:
				fmt.Println("Calling unexpected handler, check command pool")
				found = true
				break
			}
		}
	}
	if !found{
		h.call_function("/notfound","", update)
		//fmt.Println("fin")
	}
}

func Bot_handler_init(commands_handler commands_handler,log_handler log_handler) bot_handler{

	var command_pool []command_pattern

	bot_handler := bot_handler{command_pool,log_handler , commands_handler}

	bot_handler.add_command_pattern(command_pattern{"/notfound", "","log_handlers"})
	bot_handler.add_command_pattern(command_pattern{"/start", "","commands_handler"})
	bot_handler.add_command_pattern(command_pattern{"/gen", "","commands_handler"})
	bot_handler.add_command_pattern(command_pattern{"/main", "","commands_handler"})
	bot_handler.add_command_pattern(command_pattern{"/reg", "","log_handler"})

	return bot_handler
}

func ListenAndServe(handler bot_handler, channel tgbotapi.UpdatesChannel){


	fmt.Println("Bot is running...")
	for update := range channel {
		if update.Message == nil {
			continue
		}

		command, param:= handler.separate_param(update.Message.Text)
		handler.call_function(command, param, update)

	}
	/*
	fmt.Println(handler.command_pool)*/
}




