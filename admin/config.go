package admin

import (
	"net/http"

	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	"github.com/raven-chen/alpha/models"
	h "github.com/theplant/htmlgo"
)

func Initialize() *http.ServeMux {
	b := initializeProject()
	mux := SetupRouter(b)

	return mux
}

func initializeProject() (b *presets.Builder) {
	db := ConnectDB()

	// Initialize the builder of QOR5
	b = presets.New()

	// Setup the project name, ORM and Homepage
	b.URIPrefix("/admin").
		BrandTitle("alpha").
		DataOperator(gorm2op.DataOperator(db)).
		HomePageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
			r.Body = vuetify.VContainer(
				h.H1("Home"),
				h.P().Text("Change your home page here"))
			return
		})

	// Register Post into the builder
	// Use m to customize the model, Or config more models here.
	m := b.Model(&models.Post{})
	_ = m

	b.Model(&models.Product{})

	return
}
