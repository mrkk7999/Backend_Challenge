package http

import (
	"backend_challenge/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRouter(controller *controller.Controller) http.Handler {
	var (
		router = mux.NewRouter()
		// subRouter = router.PathPrefix("/loc/api/v1").Subrouter()
	)
	// Logging middleware
	// router.Use(func(next http.Handler) http.Handler {
	// 	return middleware.LoggingMiddleware(next, log)
	// })

	// API Endpoints
	router.HandleFunc("/echo", controller.Echo).Methods("POST")
	router.HandleFunc("/invert", controller.Invert).Methods("POST")
	router.HandleFunc("/flatten", controller.Flatten).Methods("POST")
	router.HandleFunc("/sum", controller.Sum).Methods("POST")
	router.HandleFunc("/multiply", controller.Multiply).Methods("POST")
	return router
}
