package main
import("fmt";
       "log";
       "html";
       "net/http")

func main() {
  //http.Handle("/foo", fooHandler)
  
  http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })
  
  log.Fatal(http.ListenAndServe(":8080", nil))
}
