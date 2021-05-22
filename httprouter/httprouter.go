package httprouter

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type SpecialRouter struct {
	router mux.Router
}

type SpecialRoute struct {
	route mux.Route
	err   error
}

func (r *SpecialRoute) addRegexpMatcher(tpl string) error {
	defaultPattern := "[^/]+"
	resources := strings.Split(tpl, "/")
}

func (r *SpecialRoute) Path(tpl string) *SpecialRoute {
	r.err = r.addRegexpMatcher(tpl)
	return r
}

func (r *SpecialRouter) SpecialHandleFunc(path string, f func(http.ResponseWriter,
	*http.Request)) *mux.Route {
	return r.router.NewRoute().Path(path).HandlerFunc(f)
}
