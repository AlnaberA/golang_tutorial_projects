package actions

import "github.com/gobuffalo/buffalo"

func HouseStarkHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("houseStark.html"))
}
