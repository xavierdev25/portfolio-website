package middleware

import (
	"net/http"
)

// MethodMiddleware restringe las solicitudes a un método HTTP específico
func MethodMiddleware(allowedMethod string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"success": false, "message": "Método no permitido"}`))
			return
		}
		next(w, r)
	}
}
