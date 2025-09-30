package handlers

import (
	"html/template"
	"net/http"

	"golang.tailwind/models"
	"golang.tailwind/utils"
)

func EkstrakulikulerHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		Title          string
		Navbar         template.HTML
		Footer         template.HTML
		Ekstrakulikuler []models.DataEskul
	}{
		Title:          "Ekstrakulikuler",
		Navbar:         template.HTML(utils.GetNavbar()),
		Footer:        template.HTML(utils.GetFooter()),
		Ekstrakulikuler: models.IsiEskul, // ini slice []DataEskul
	}

	tmpl, err := template.ParseFiles("templates/ekstrakulikuler.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, page); err != nil {
		http.Error(w, "Render error: "+err.Error(), http.StatusInternalServerError)
	}
}
