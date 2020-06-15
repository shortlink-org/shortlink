package smtp

import (
	"net/smtp"

	"github.com/spf13/viper"
)

type Bot struct {
	from string
	pass string
	to   string
	host string
	addr string
}

func (b *Bot) Init() error {
	// Set configuration
	b.setConfig()

	return nil
}

func (b *Bot) Send(message string) error {
	msg := `To: "Some User" <vitya.login@yandex.ru>
From: "Other User" <vitya.login@yandex.ru>
Subject: Add new link!!

` + message

	auth := smtp.PlainAuth("", b.from, b.pass, b.host)

	err := smtp.SendMail(
		b.addr,
		auth,
		b.from,
		[]string{b.to},
		[]byte(msg),
	)
	if err != nil {
		return err
	}

	return nil
}

// setConfig - set configuration
func (b *Bot) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("BOT_SMTP_FROM", "example@site.com")
	viper.SetDefault("BOT_SMTP_PASS", "YOUR_PASSWORD")
	viper.SetDefault("BOT_SMTP_TO", "EMAIL_USER")
	viper.SetDefault("BOT_SMTP_HOST", "smtp.gmail.com")
	viper.SetDefault("BOT_SMTP_ADDR", "smtp.gmail.com:587")

	b.from = viper.GetString("BOT_SMTP_FROM")
	b.pass = viper.GetString("BOT_SMTP_PASS")
	b.to = viper.GetString("BOT_SMTP_TO")
	b.host = viper.GetString("BOT_SMTP_HOST")
	b.addr = viper.GetString("BOT_SMTP_ADDR")
}
