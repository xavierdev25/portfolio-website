package middleware

import (
	"net/http"
	"sync"
	"time"
)

// RateLimiter almacena información sobre las peticiones por IP
type RateLimiter struct {
	visitors map[string]*Visitor
	mu       sync.RWMutex
	limit    int           // Número máximo de peticiones
	window   time.Duration // Ventana de tiempo
}

// Visitor representa un visitante con sus peticiones
type Visitor struct {
	limiter  *rate
	lastSeen time.Time
}

type rate struct {
	tokens   int
	lastTime time.Time
}

// NewRateLimiter crea un nuevo rate limiter
func NewRateLimiter(requestsPerWindow int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*Visitor),
		limit:    requestsPerWindow,
		window:   window,
	}

	// Limpiar visitantes inactivos cada 5 minutos
	go rl.cleanupVisitors()

	return rl
}

// Allow verifica si una IP puede hacer una petición
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	visitor, exists := rl.visitors[ip]
	if !exists {
		visitor = &Visitor{
			limiter: &rate{
				tokens:   rl.limit,
				lastTime: time.Now(),
			},
			lastSeen: time.Now(),
		}
		rl.visitors[ip] = visitor
	}

	visitor.lastSeen = time.Now()

	// Calcular tokens disponibles basado en el tiempo transcurrido
	now := time.Now()
	elapsed := now.Sub(visitor.limiter.lastTime)
	
	// Si ha pasado la ventana completa, resetear tokens
	if elapsed >= rl.window {
		visitor.limiter.tokens = rl.limit
		visitor.limiter.lastTime = now
	}

	// Verificar si hay tokens disponibles
	if visitor.limiter.tokens > 0 {
		visitor.limiter.tokens--
		return true
	}

	return false
}

// cleanupVisitors elimina visitantes inactivos
func (rl *RateLimiter) cleanupVisitors() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		for ip, visitor := range rl.visitors {
			if time.Since(visitor.lastSeen) > 30*time.Minute {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// RateLimitMiddleware es el middleware de rate limiting
func RateLimitMiddleware(rl *RateLimiter) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ip := GetClientIP(r)

			if !rl.Allow(ip) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte(`{"success": false, "message": "Demasiadas peticiones. Por favor, intenta más tarde."}`))
				return
			}

			next(w, r)
		}
	}
}

// GetClientIP obtiene la IP real del cliente (exportada para uso en otros paquetes)
func GetClientIP(r *http.Request) string {
	// Intentar obtener la IP desde headers de proxy
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// Tomar la primera IP si hay múltiples
		for i := 0; i < len(ip); i++ {
			if ip[i] == ',' {
				return ip[:i]
			}
		}
		return ip
	}

	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	// Si no hay headers de proxy, usar RemoteAddr
	for i := len(r.RemoteAddr) - 1; i >= 0; i-- {
		if r.RemoteAddr[i] == ':' {
			return r.RemoteAddr[:i]
		}
	}

	return r.RemoteAddr
}
