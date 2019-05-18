package grifts

import (
	"golang_tutorial_projects/GameOfBuffalo/site/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
