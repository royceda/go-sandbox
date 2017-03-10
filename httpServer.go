package main

import("net/http"
"io"
"fmt")



func controller(res http.ResponseWriter, req *http.Request){

  res.Header().Set("Content-Type", "text/html");
  io.WriteString(
    res,
  `<DOCTYPE html>
  <html>
  <head>
  <title>Hello World </title>
  </head>
  <body>
  <h1>Hello World</h1>
  <p>This page is sent with GoLang</p>
  </body>
  </html>`)
}


func main(){
  http.HandleFunc("/hello", controller); //to the route /hello we have the function page as controller
  fmt.Println("The magic port is 8888")
  http.ListenAndServe(":8888", nil);
}
