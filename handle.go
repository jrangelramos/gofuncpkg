package function

import (
	"fmt"
	"function/hello"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, hello.Hello("Jeff"))
}
