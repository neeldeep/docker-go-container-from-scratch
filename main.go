package main

import (
  "html/template"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("template/page.html")
  if err != nil {
    panic(err)
  }
  t.Execute(w, nil)
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9190", nil)
}