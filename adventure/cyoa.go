package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func ParseJSON(file *os.File) (Story, error) {
	jsonBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var story Story
	err = json.Unmarshal(jsonBytes, &story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

// Funtional options
type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFn(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

var defaultTpl = template.Must(template.ParseFiles("web/layout.go.html"))

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, defaultTpl, defaultPathFn}
	for _, opt := range opts {
		// Pass reference to the handler so that can be modified
		opt(&h)
	}
	return h
}

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

const (
	PageNotFound        = "Chapter not found."
	InternalServerError = "Something went wrong..."
)

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)
	if chapter, ok := h.s[path]; ok {
		if err := h.t.Execute(w, chapter); err != nil {
			log.Printf("%v", err)
			http.Error(w, InternalServerError, http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, PageNotFound, http.StatusNotFound)

}

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
