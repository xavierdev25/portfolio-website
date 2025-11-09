package main

import (
	"log"
	"net/http"
	"os"

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

	// Configurar el multiplexor HTTP
	mux := http.NewServeMux()

	// Rutas
	mux.HandleFunc("/api/health", handlers.HealthCheckHandler)
	mux.HandleFunc("/api/contact", middleware.MethodMiddleware("POST", handlers.ContactHandler))

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   getAllowedOrigins(),
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Aplicar CORS al handler
	handler := c.Handler(mux)

	// Iniciar el servidor
	log.Printf("🚀 Servidor iniciado en http://localhost:%s", port)
	log.Printf("✅ Health check disponible en http://localhost:%s/api/health", port)
	log.Printf("📧 Endpoint de contacto en http://localhost:%s/api/contact", port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
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
