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

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

}

var lead Lead 

