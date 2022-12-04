package admin

import (
	"net/http"

	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB, mux *http.ServeMux) *presets.Builder {
	b := initializeProject(db)
	mux.Handle("/", b)

	return b
}

func initializeProject(db *gorm.DB) (b *presets.Builder) {
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

	return
}
