package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/xavierdev25/portfolio-backend/handlers"
	"github.com/xavierdev25/portfolio-backend/middleware"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: No se encontró el archivo .env")
	}

	// Obtener el puerto del entorno o usar el predeterminado
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Crear rate limiter: máximo 5 peticiones cada 15 minutos por IP
	rateLimiter := middleware.NewRateLimiter(5, 15*time.Minute)

	// Configurar el multiplexor HTTP
	mux := http.NewServeMux()

	// Rutas
	mux.HandleFunc("/api/health", handlers.HealthCheckHandler)

	// Aplicar rate limiting solo al endpoint de contacto
	contactHandler := middleware.RateLimitMiddleware(rateLimiter)(
		middleware.MethodMiddleware("POST", handlers.ContactHandler),
	)
	mux.HandleFunc("/api/contact", contactHandler)

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   getAllowedOrigins(),
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	// Aplicar headers de seguridad y CORS
	handler := middleware.SecurityHeadersMiddleware(c.Handler(mux))

	// Configurar el servidor con timeouts
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Iniciar el servidor en una goroutine
	go func() {
		log.Printf("🚀 Servidor iniciado en http://localhost:%s", port)
		log.Printf("✅ Health check disponible en http://localhost:%s/api/health", port)
		log.Printf("📧 Endpoint de contacto en http://localhost:%s/api/contact", port)
		log.Printf("🛡️  Rate limit: 5 peticiones cada 15 minutos por IP")
		log.Printf("⏱️  Timeouts configurados: Read 10s | Write 10s | Idle 120s")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar el servidor: %v", err)
		}
	}()

	// Esperar señal de interrupción para apagado graceful
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("🛑 Apagando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error en el apagado del servidor: %v", err)
	}

	log.Println("✅ Servidor apagado correctamente")
}

// getAllowedOrigins obtiene los orígenes permitidos desde las variables de entorno
func getAllowedOrigins() []string {
	origins := os.Getenv("ALLOWED_ORIGINS")
	if origins == "" {
		return []string{"http://localhost:4321"}
	}

	// Separar por comas si hay múltiples orígenes
	var originsList []string
	for i, j := 0, 0; j <= len(origins); j++ {
		if j == len(origins) || origins[j] == ',' {
			if j > i {
				originsList = append(originsList, origins[i:j])
			}
			i = j + 1
		}
	}

	return originsList
}
