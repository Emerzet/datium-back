package handlers
import (
	"log"
	"fmt"
	"net/http"
)

type Hello struct {

	l *log.Logger

 

}

func NewHello(l *log.Logger) *Hello {
return &Hello{l}

}



func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {


http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){
	log.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Elo %s", d)
	if err !=nil {
		http.Error(w, "snisz", http.StatusBadRequest)
	return
}

}
