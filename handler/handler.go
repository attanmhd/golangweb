package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Sedang terjadi kesalahan , mohon bersabar", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "GolangWeb",
		"content": "I'm learning golangweb",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sedang terjadi kesalahan , mohon bersabar", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya belajar golang"))
}

func WorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good Morning World"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Sedang terjadi kesalahan , mohon bersabar", http.StatusInternalServerError)
		return
	}

	dataProduct := []entity.Product{
		{ID: 1, Name: "Shirt", Price: 120000, Stock: 12},
		{ID: 2, Name: "Short", Price: 160000, Stock: 4},
		{ID: 3, Name: "Shoes", Price: 180000, Stock: 3},
	}

	err = tmpl.Execute(w, dataProduct)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sedang terjadi kesalahan , mohon bersabar", http.StatusInternalServerError)
		return
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Sedang terjadi kesalahan , mohon bersabar", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Sedang terjadi kesalahan , mohon bersabar", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Erro is happening, keep calm", http.StatusBadRequest)

}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Sedang terjadi kesalahan , mohon bersabar", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		// message := r.Form.Get("message")

		w.Write([]byte(name))
		// w.Write([]byte(message))

		return

	}
	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)

}
