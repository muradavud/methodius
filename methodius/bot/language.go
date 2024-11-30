package bot

import (
	_ "embed"

	tele "gopkg.in/telebot.v3"
)

func SendLanguageInline(c tele.Context) {
	selector := &tele.ReplyMarkup{
		RemoveKeyboard: true,
	}

	btn1 := selector.Data("🇺🇸", "en-US")
	btn2 := selector.Data("🇷🇺", "ru-RU")
	btn3 := selector.Data("🇹🇷", "tr-TR")
	btn4 := selector.Data("🇩🇪", "de-DE")
	btn5 := selector.Data("🇪🇸", "es-ES")
	btn6 := selector.Data("🇮🇹", "it-IT")

	selector.Inline(
		selector.Row(btn1, btn2, btn3),
		selector.Row(btn4, btn5, btn6),
	)

	c.Send(KV["pick_language"], selector)
}
