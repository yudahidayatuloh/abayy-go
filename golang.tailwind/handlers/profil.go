package handlers

import (
	"html/template"
	"net/http"

	"golang.tailwind/models"
	"golang.tailwind/utils"
)

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	// bikin struct data khusus profil
	page := struct {
		Title  string
		Navbar template.HTML
		Footer template.HTML
		Profil models.DataProfil
	}{
		Title:  "Profil Sekolah",
		Navbar: template.HTML(utils.GetNavbar()),
		Footer:        template.HTML(utils.GetFooter()),
		Profil: models.IsiProfil, // ambil dari models/profil.go
	}

	// parsing template
	tmpl, err := template.ParseFiles("templates/profil.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// eksekusi template dengan data page
	if err := tmpl.Execute(w, page); err != nil {
		http.Error(w, "Render error: "+err.Error(), http.StatusInternalServerError)
	}
}
