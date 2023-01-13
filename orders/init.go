package orders

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/qor5/admin/presets"
	"gorm.io/gorm"
)

//go:embed views/*
var viewFiles embed.FS
var templates map[string]*template.Template

func Init(db *gorm.DB, mux *http.ServeMux, b *presets.Builder) {
	if err := db.AutoMigrate(Product{}); err != nil {
		panic(err)
	}
	mux.Handle("/menu", Menu(db))
	b.Model(&Product{})

	if err := loadTemplates(); err != nil {
		panic(err)
	}
}

func loadTemplates() error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	tmplFiles, err := fs.ReadDir(viewFiles, "views")
	if err != nil {
		return err
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, err := template.ParseFS(viewFiles, "views/"+tmpl.Name())
		if err != nil {
			return err
		}

		templates[tmpl.Name()] = pt
	}

	return nil
}
