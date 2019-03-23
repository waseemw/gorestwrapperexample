package main

import (
	"./simprest"
	"fmt"
)

func main() {
	r := simprest.NewRouter()
	r.Get("/{var}", func(w simprest.RWriter, r map[string]string) {
		_, _ = fmt.Fprint(w, r["var"])
	})
	r.Listen(80)
}
