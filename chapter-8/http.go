package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type", "text/html",
	)
	_, err := io.WriteString(
		res,
		`<DOCTYPE html>
			<html>
				<head>
					<title>Hello, World</title>
				</head>
				<body>
					Hello, World!
				</body>
			</html>`,
	)
	if err != nil {
		fmt.Println(err)
		return

	}
}

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
