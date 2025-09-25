package handlers

import (
	"net/http"
)


type Lead struct {
	name string
	surname string
	phone string
	email string
	
	


}




func LeadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w,"Method Not Allowed", http.StatusMethodNotAllowed)
		return 
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

}

var lead Lead 

