package actions

import "github.com/gobuffalo/buffalo"

func HouseLanisterHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("houseLanister.html"))
}
