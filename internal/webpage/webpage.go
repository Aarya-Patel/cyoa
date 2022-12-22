package webpage

import (
	"html/template"
	"net/http"

	"github.com/Aarya-Patel/cyoa/model"
)

func CreateWebPageHandlerForArc(t *template.Template, arc model.Arc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, arc)
		if err != nil {
			panic(err)
		}
	}

}

func CreateWebPagesForStory(story model.Story) http.Handler {
	mux := http.NewServeMux()
	t, err := template.ParseFiles("./template/template.html")
	if err != nil {
		panic(err)
	}

	for arcKey, arc := range story.Arcs {
		serveFunc := CreateWebPageHandlerForArc(t, arc)

		path := "/" + arcKey
		if arcKey == "intro" {
			path = "/"
		}
		mux.HandleFunc(path, serveFunc)
	}

	return mux
}
