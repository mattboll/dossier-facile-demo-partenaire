package grifts

import (
	"dossier_facile_demo_partenaire/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
