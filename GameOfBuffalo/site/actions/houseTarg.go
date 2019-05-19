package actions

import "github.com/gobuffalo/buffalo"

func HouseTargHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("houseTarg.html"))
}
