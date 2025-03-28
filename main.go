package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	m := http.NewServeMux()
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT default to 8080")
		port = "8080"
	}
	port = ":" + port

	m.HandleFunc("/", handlePage)

	srv := http.Server{
		Handler:      m,
		Addr:         port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("PORT env variable:", os.Getenv("PORT"))
	fmt.Println("server started on ", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a simple Go server. </p>
</body>
</html>
`
	w.Write([]byte(page))
}
