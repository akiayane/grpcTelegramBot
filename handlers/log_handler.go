package handlers

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"net/smtp"
	"os"
)

type log_handler struct {

	functions_array []pattern
	bot *tgbotapi.BotAPI
	u *tgbotapi.UpdateConfig

}

func(h *log_handler) add_new_function(p pattern){
	h.functions_array = append(h.functions_array, p)
}

func(h *log_handler) call_function(update tgbotapi.Update,pattern_name string, param string){
	found := false
	for _, value := range h.functions_array {
		if value.pattern_name == pattern_name {
			call := value.func_to_call
			call(update,param)
			found = true
			break
		}
	}
	if !found{
		call:=h.functions_array[0].func_to_call
		call(update, "")
	}
}


func Log_handler_init(bot *tgbotapi.BotAPI, u *tgbotapi.UpdateConfig, channel *tgbotapi.UpdatesChannel) log_handler{
	functions_array := []pattern{}


	handler := log_handler{functions_array, bot, u}

	//adding functionalities
	//function for unknown input, always should be first
	handler.add_new_function(pattern{"/notfound", func(update tgbotapi.Update, param string) {

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I dont know such command")
		fmt.Println("start called")
		_, _ = bot.Send(msg)

	}})
	//register email
	handler.add_new_function(pattern{"/reg", func(update tgbotapi.Update, param string) {

		from := os.Getenv("GMAIL")
		password := os.Getenv("GMAIL_PASSWORD")
		to := []string{
			param,
		}
		smtpHost := os.Getenv("SMTP_HOST")
		smtpPort := os.Getenv("SMTP_POST")

		message := []byte("Test, hello "+param)
		auth := smtp.PlainAuth("", from, password, smtpHost)
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Email Sent Successfully!")

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "The email with confirmation link has been sent to given email: " + param + ". Please, chek spam folder if you will not find it!")
		_, _ = bot.Send(msg)

	}})

	return handler
}