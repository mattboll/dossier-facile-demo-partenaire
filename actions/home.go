package actions

import (
	"dossier_facile_demo_partenaire/models"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

func CallbackHandler(c buffalo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	e := models.Event{
		Content: string(b),
	}
	err = models.DB.Save(&e)
	if err != nil {
		log.Panic(err)
		return err
	}
	return c.Render(http.StatusOK, r.JSON(""))
}
