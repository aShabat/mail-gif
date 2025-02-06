package main

import (
	"fmt"
	"image/gif"
	"net/http"
	"strconv"

	rand "math/rand/v2"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"mail-gif/cmd/server"
	"mail-gif/cmd/server/models"
	"mail-gif/cmd/web/html"
)

type Application struct {
	gifs models.GifStoreInterface
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	app := &Application{gifs: models.GifStoreInit()}
	router.Get("/", app.home)
	router.Get("/{index}", app.homeWith)
	router.Post("/send_page", app.postPage)
	router.Post("/gif/{index}", app.postGif)
	fmt.Println("setup done")

	if err := http.ListenAndServe(":4000", router); err != nil {
		panic(err)
	}
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	html.Home().Render(w)
}

func (app *Application) homeWith(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(chi.URLParam(r, "index"))
	html.HomeWith(index, false).Render(w)
}

func (app *Application) postPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["page-type"]) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	page_type := r.Form["page-type"][0]
	page := r.Form["page"][0]
	if page_type == "link" {
		index := rand.Int()
		for {
			if _, ok := app.gifs.Get(index); ok {
				index = rand.Int()
			} else {
				break
			}
		}
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()
			page := server.OpenUrlPage(page)
			gif := server.PageGif(page)
			app.gifs.Add(gif, index)
		}()

		fmt.Println("sent", fmt.Sprintf("/%d", index))
		http.Redirect(w, r, fmt.Sprintf("/%d", index), http.StatusSeeOther)
		return
	} else if page_type == "html" {
		index := rand.Int()
		for {
			if _, ok := app.gifs.Get(index); ok {
				index = rand.Int()
			} else {
				break
			}
		}
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()
			page := server.OpenHtmlPage(page)
			gif := server.PageGif(page)
			app.gifs.Add(gif, index)
		}()

		http.Redirect(w, r, fmt.Sprintf("/%d", index), http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) postGif(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(chi.URLParam(r, "index"))
	if g, ok := app.gifs.Get(index); ok {
		w.Header().Set("Content-Disposition", "attachment; filename=example.gif;")
		gif.EncodeAll(w, g)
	} else {
		html.HomeWith(index, true).Render(w)
	}
}
