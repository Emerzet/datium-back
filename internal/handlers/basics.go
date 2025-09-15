package handlers


import "net/http"




func FormHandler (w http.Response, r *http.Request){
	if r.Method =! r.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
		
	}
{




