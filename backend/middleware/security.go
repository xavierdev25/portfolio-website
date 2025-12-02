package middleware

import "net/http"

// SecurityHeadersMiddleware añade headers de seguridad HTTP a todas las respuestas
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevenir que el navegador interprete el contenido como otro tipo MIME
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Prevenir clickjacking
		w.Header().Set("X-Frame-Options", "DENY")

		// Habilitar protección XSS del navegador
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		// Controlar qué información se envía en el header Referer
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		// Content Security Policy - restringir fuentes de contenido
		w.Header().Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")

		// Prevenir que la API sea cacheada
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")

		// Strict Transport Security (HTTPS)
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		// Deshabilitar detección de tipo de contenido
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		next.ServeHTTP(w, r)
	})
}
