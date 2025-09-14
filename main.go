package main

import ("net/http"
	"fmt"
	"io/ioutil"
)







func main(){
	l := log.New(os.Stout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	s := &http.Server{
	addr: ":9090",
	Handler: sm,
	IdleTimeout: 120 * time.Second,
	ReadTimeout: 1 * time.Second,
	WriteTimeout: 1 * time.Second
	
	}

go func() {
	err := s.ListenAndServe()
	if err != nil {
		l.Fatal(err)
	}()

	sigChan := make (chan os.Signal)
	signal.Notify(sigChan, os.Inrerrupt)
	signal.Notify(sigChan, os.Kill)
	
	sig <- sigChan
	l.Println("Recived termiante, graceful shutdown", sig)


tc := context.WithTimeOut(context.Background(), 30 * time.Second)
s.Shutdown(tc)

}
