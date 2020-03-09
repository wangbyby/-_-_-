package middleware

import (
	"net/http"
	"net/http/httptest"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

//1
type SingleHost struct {
	Handler   http.Handler
	AllowHost string
}

func (this *SingleHost) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Host == this.AllowHost {
		this.Handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

//2.
func SingleHostFunc(handler http.Handler, allowhost string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		println(r.Host)
		if r.Host == allowhost {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	}

	return http.HandlerFunc(fn)
}

//3.
type AppendMidware struct {
	handler http.Handler
}

func (this *AppendMidware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.handler.ServeHTTP(w, r)
	w.Write([]byte("midware done sth"))

}

//4.
type ModifiedMidware struct {
	handler http.Handler
}

func (this *ModifiedMidware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := httptest.NewRecorder()
	this.handler.ServeHTTP(rec, r)

	for k, v := range rec.Header() {
		w.Header()[k] = v

	}
	w.Header().Set("go-web", "vip")
	w.WriteHeader(418)
	w.Write([]byte("hey this is midware"))
	w.Write(rec.Body.Bytes())
}

// func main(){
// 	single := Singlepost{
// 		Handler: http.HandleFunc(myHandler),
// 		AllowHost: "hello.com"
// 	}
// 	sin := SingleHostFunc(http.HandlerFunc(MyHandler), "localhost:8088")
// 	http.ListenAndServe(":8088", sin)

// 	mid := &AppendMidware{http.HandlerFunc(MyHandler)}
// 	http.ListenAndServe(":8088", mid)

// 	mmid := &ModifiedMidware{http.HandlerFunc(MyHandler)}
// 	http.ListenAndServe(":8088", mmid)
// }
