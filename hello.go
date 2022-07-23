package main

import (
	"fmt"
	"http"
)

func main1() {
	a := 1
	b := 2
	fmt.Println("!... Hello World ...!")
	fmt.Println(1 + 5)
	fmt.Println(a + b)
	c := a + b
	fmt.Println(c)
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	http.ServeFile(res, req, "/Users/surajvitekar/HTML/index.html")

	switch req.Method {
	case "GET":
		http.ServeFile(res, req, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(res, "Post from website! r.PostFrom = %v\n", req.PostForm)
		name := req.FormValue("fname")
		address := req.FormValue("address")
		fmt.Fprintf(res, "Name = %s\n", name)
		fmt.Fprintf(res, "Address = %s\n", address)
	default:
		fmt.Fprintf(res, "Sorry, only GET and POST methods are supported.")
	}
}

func Handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hi"))
}

func Handler2(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hi suraj"))
}

func home(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	http.ServeFile(res, req, "/Users/surajvitekar/plain.html")
}
