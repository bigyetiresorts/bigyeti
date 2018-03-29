package actions

import "github.com/gobuffalo/buffalo"

// TestemailIndex default implementation.
func TestemailIndex(c buffalo.Context) error {
	c.Set("title", []string{"Big Yeti Promo Email"})
	return c.Render(200, r.HTML("testemail/index.html"))
}
