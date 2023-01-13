package orders

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

const menuTmpl = "menu.tmpl"

func Menu(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prods := []Product{}
		if err := db.Find(&prods).Error; err != nil {
			fmt.Fprintln(w, "<div style='color:red;'>There is an error occurred</div>")
			return
		}
		tmpl, ok := templates[menuTmpl]
		if !ok {
			fmt.Fprintf(w, "<div style='color:red;'>Cannot find template: %s</div>", menuTmpl)
			return
		}

		data := map[string]interface{}{
			"Products": prods,
		}

		if err := tmpl.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<div style='color:red;'>There is an error occurred: %s</div>", err.Error())
		}
	})
}
