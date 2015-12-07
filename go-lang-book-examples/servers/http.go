package main

import ("net/http" ; "io")

func helloHtmlResponse() string {
    return `
<DOCTYPE html>
<html>
  <head>
      <title>Hello World</title>
  </head>
  <body>
      Hello World!
  </body>
</html>`
}

func nameHtmlResponse() string {
    return `
<DOCTYPE html>
<html>
  <head>
      <title>Hello World</title>
  </head>
  <body>
      Xunda!
  </body>
</html>`
}

func hello(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "text/html")
    io.WriteString(res, helloHtmlResponse(), )
}

func name(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "text/html")
    io.WriteString(res, nameHtmlResponse(), )
}

func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/name", name)
    http.ListenAndServe(":9000", nil)
}