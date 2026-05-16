package middleware

import "net/http"

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		responseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		responseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if request.Method == http.MethodOptions {
			responseWriter.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(responseWriter, request)
	})
}
