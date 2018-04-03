package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
)

func SendWelcomeMail() error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "Welcome Mail"
	m.From = ""
	m.To = []string{}
	err := m.AddBody(r.HTML("welcome_mail.html"), render.Data{})
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
