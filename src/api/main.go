package main
 
import (
	"io"
	"log"
	"net/http"
)
 
func main() {
	indexFunc := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello World!")
	}
 
	http.HandleFunc("/", indexFunc)
	log.Fatal(http.ListenAndServe(":8000", nil))
}