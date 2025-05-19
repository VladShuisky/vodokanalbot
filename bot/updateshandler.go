package bot

import (
	"github.com/VladShuisky/vodokanalbot/parsing"
	"github.com/VladShuisky/vodokanalbot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(update tgbotapi.Update) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	switch update.Message.Command() {
	case "start":
		message := GetStartMessage()
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, message)
		msg.ReplyToMessageID = update.Message.MessageID
	case "get_last_info":
		htmlFromVodokanal := parsing.GetHtmlDataFromVodokanal()
		targetTexts := parsing.ExtractText(htmlFromVodokanal)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, targetTexts[1])
		msg.ReplyToMessageID = update.Message.MessageID
	case "date":
		all_command_with_text := update.Message.Text
		dateStr := utils.TrimTelegramCommand(all_command_with_text)
		htmlFromVodokanal := parsing.GetHtmlDataFromVodokanal()
		targetTexts := parsing.ExtractText(htmlFromVodokanal)
		current_data_info, err := parsing.GetContentByDate(dateStr, targetTexts[1:])
		var content string = ""
		if err != nil {
			content = "Введите дату правильно"
		} else {
			content = utils.JoinWithParagraphs(current_data_info)
		}

		msg = tgbotapi.NewMessage(update.Message.Chat.ID, content)
	case "db_healthcheck":
		check := utils.CheckDbConnect()
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, check)
		msg.ReplyToMessageID = update.Message.MessageID

	default:
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
	}
	return msg
}

	// case "dev":
	// 	msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 	msg.ReplyToMessageID = update.Message.MessageID
	// 	fmt.Println("-----dev------")
	// 	spew.Dump(update.Message.Chat)
	// 	fmt.Println("--------------")
	// case "dev_orm":
	// 	msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 	db := database.GetDb()
	// 	fmt.Println("-----dev gorm------")
	// 	spew.Dump(db)
	// 	fmt.Println("--------------")
	// case "dev_orm_create_user":
	// 	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "db row created!")
	// 	db := database.GetDb()
	// 	recipient := database.TelegramRecipient{
	// 		TelegramChatId: update.Message.Chat.ID,
	// 		Data: database.JSONB{
	// 			"foo": "bar",
	// 		},
	// 	}
	// 	db.Create(&recipient)
	// case "dev_orm_delete_user":
	// 	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "db row deleted!")
	// 	db := database.GetDb()
	// 	var recipient database.TelegramRecipient
	// 	db.Unscoped().Where("telegram_chat_id = ?", update.Message.Chat.ID).Delete(&recipient)
	// 	// db.Delete(&recipient)