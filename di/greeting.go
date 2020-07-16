package di

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
