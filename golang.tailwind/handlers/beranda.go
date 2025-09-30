package handlers

import (
	"html/template"
	"net/http"

	"golang.tailwind/models"
	"golang.tailwind/utils"
)

// Data jurusan disiapkan di sini
var isiJurusan = []models.DataJurusan{
	{"PPLG", "Pengembangan Perangkat Lunak dan Gim", "https://web-smkn4.vercel.app/_nuxt/pplg.DhbRSnK2.png", ""},
	{"TJKT", "Teknik Jaringan Komputer dan Telekomunikasi", "https://web-smkn4.vercel.app/_nuxt/tjkt.XeO8V4_I.png", ""},
	{"TBSM", "Teknik Sepeda Motor", "https://web-smkn4.vercel.app/_nuxt/tbsm.5YZNeyvV.png", ""},
	{"DKV", "Desain Komunikasi Visual", "https://web-smkn4.vercel.app/_nuxt/dkv.Btdc-HP2.png", ""},
	{"TOI", "Teknik Otomasi Industri", "https://web-smkn4.vercel.app/_nuxt/toi.BFD6ZBmB.png", ""},
}

// Handler untuk halaman Beranda
func Beranda(w http.ResponseWriter, r *http.Request) {
	// Tambahkan warna ke setiap jurusan
	var jurusanWithColors []models.DataJurusan
	for _, j := range isiJurusan {
		j.Warna = utils.GetWarna(j.NamaJ) // ambil warna dari utils
		jurusanWithColors = append(jurusanWithColors, j)
	}

	// Data yang dikirim ke template
	page := struct {
		Title         string
		Jurusan       []models.DataJurusan
		JumlahJurusan int
		JumlahEskul   int
		JumlahSiswa   int
		Navbar        template.HTML
		Footer        template.HTML
	}{
		Title:         "Beranda",
		Jurusan:       jurusanWithColors,
		JumlahJurusan: len(jurusanWithColors),
		JumlahEskul:   len(models.IsiEskul),
		JumlahSiswa:   len(models.IsiSiswa),
		Navbar:        template.HTML(utils.GetNavbar()), // Navbar dipanggil dari utils
		Footer:        template.HTML(utils.GetFooter()), // Navbar dipanggil dari utils
	}

	// Render halaman dengan template
	renderTemplate(w, "beranda.html", page)
}

// Fungsi helper untuk render template
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Render error: "+err.Error(), http.StatusInternalServerError)
	}
}
