package simprest

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Router struct {
	r *mux.Router
}

type RWriter http.ResponseWriter

func NewRouter() *Router {
	r := Router{mux.NewRouter()}
	return &r
}

func (r *Router) Get(url string, handle func(RWriter, map[string]string)) {
	r.r.HandleFunc(url, func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		handle(w, vars)
	})
}

func (r *Router) Listen(port int) {
	println("Listening on :" + strconv.Itoa(port))
	_ = http.ListenAndServe(":"+strconv.Itoa(port), r.r)
}
