package main

import "fmt"
import "net/http"

func handleFunc(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"<h1>hello,这里是 goblog</h1>")
}
func main(){
	http.HandleFunc("/",handleFunc)
	http.ListenAndServe(":3000",nil)
}