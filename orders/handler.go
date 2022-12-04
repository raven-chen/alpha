package orders

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Menu(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prods := []Product{}
		if err := db.Find(&prods).Error; err != nil {
			fmt.Fprintln(w, "<div style='color:red;'>This is an error occurred</div>")
			return
		}

		content := "<table> <tr><td>Name</td><td>Price</td></tr>"
		for _, prod := range prods {
			content += fmt.Sprintf("<tr><td>%s</td><td>%v</td></tr>", prod.Name, prod.Price)
		}
		content += "</table>"

		fmt.Fprintln(w, content)
	})
}
