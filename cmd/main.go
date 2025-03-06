package main

import (
	"backend_challenge/controller"
	"backend_challenge/implementation"
	httpTransport "backend_challenge/transportation/http"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	var (
		httpAddr = os.Getenv("HTTP_ADDR")
	)

	svc := implementation.New()
	ctrl := controller.New(svc)
	handler := httpTransport.SetUpRouter(ctrl)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		server := &http.Server{
			Addr:    httpAddr,
			Handler: handler,
		}
		errs <- server.ListenAndServe()
	}()

	log.Printf("server stopped: %v", <-errs)

}

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

// func main() {
// 	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
// 		file, _, err := r.FormFile("file")
// 		if err != nil {
// 			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
// 			return
// 		}
// 		defer file.Close()
// 		records, err := csv.NewReader(file).ReadAll()
// 		if err != nil {
// 			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
// 			return
// 		}
// 		var response string
// 		for _, row := range records {
// 			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
// 		}
// 		fmt.Fprint(w, response)
// 	})
// 	http.ListenAndServe(":8080", nil)
// }
