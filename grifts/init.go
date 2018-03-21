package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/obiknows/bigyeti/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
