package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		// Prepare HTML response with CSS styling
		html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Fisio Raupp Calendar</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f5f5f5;
					margin: 0;
					padding: 20px;
				}
				.container {
					background-color: #fff;
					border-radius: 8px;
					box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
					padding: 20px;
					max-width: 600px;
					margin: 0 auto;
				}
				h1 {
					color: #4285f4;
				}
				p {
					margin-top: 16px;
					margin-bottom: 16px;
					word-wrap: break-word;
				}
				strong {
					font-weight: bold;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>OAuth Response</h1>
				<p><strong>Code:</strong> %s</p>
			</div>
		</body>
		</html>
		`, code)

		// Write HTML response
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, html)
	})

	fmt.Println("Server is running on http://localhost/")
	http.ListenAndServe(":8080", nil)
}