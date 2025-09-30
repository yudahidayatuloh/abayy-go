package handlers

import (
	"html/template"
	"net/http"

	"golang.tailwind/models"
	"golang.tailwind/utils"
)

func GaleriHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		Title string
		Navbar template.HTML
		Footer template.HTML
		Galeri []models.DataGaleri
	}{
		Title: "Galeri",
		Navbar: template.HTML(utils.GetNavbar()),
		Footer: template.HTML(utils.GetFooter()),
		Galeri: models.IsiGaleri,
	}

	tmpl, err := template.ParseFiles("templates/galeri.html")
	if err != nil {
		http.Error(w, "Template error:" +err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, page); err != nil {
		http.Error(w, "Render error: "+err.Error(), http.StatusInternalServerError)
	}
}
