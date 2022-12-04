package orders

import (
	"net/http"

	"github.com/qor5/admin/presets"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, mux *http.ServeMux, b *presets.Builder) {
	err := db.AutoMigrate(Product{})
	mux.Handle("/menu", Menu(db))
	b.Model(&Product{})

	if err != nil {
		panic(err)
	}
}
