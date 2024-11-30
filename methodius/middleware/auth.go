package middleware

import (
	"fmt"
	"methodius/config"
	"methodius/handler"
	"methodius/process"

	tele "gopkg.in/telebot.v3"
)

func Auth(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		item, err := process.GetChatFromDynamo(fmt.Sprint(c.Chat().ID)) // TODO caching
		if err != nil {
			c.Send("Error occured. Please try again later.")
			return err
		}

		if item.IsAuthorized {
			return next(c)
		}

		if c.Message().Text == config.Conf.UserPassword {
			err = process.UploadChatToDynamo(item.ChatId, item.Username, item.FirstName,
				item.LastName, item.History, item.Language, true)
			if err != nil {
				c.Send("Error occured. Please try again later.")
				return err
			}

			c.Delete()

			handler.OnAuthorized(c)

			return nil
		}

		c.Send("You are not authorized to use this bot. Please enter the password.")
		return err
	}
}
