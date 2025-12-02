package utils

import (
	"log"
	"time"
)

// SecurityLogger registra eventos de seguridad sospechosos
type SecurityLogger struct {
	enabled bool
}

// NewSecurityLogger crea un nuevo logger de seguridad
func NewSecurityLogger(enabled bool) *SecurityLogger {
	return &SecurityLogger{enabled: enabled}
}

// LogSuspiciousActivity registra actividad sospechosa
func (sl *SecurityLogger) LogSuspiciousActivity(ip, reason, details string) {
	if !sl.enabled {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("🚨 [SECURITY] %s | IP: %s | Razón: %s | Detalles: %s", 
		timestamp, ip, reason, details)
}

// LogRateLimitExceeded registra cuando se excede el rate limit
func (sl *SecurityLogger) LogRateLimitExceeded(ip string) {
	sl.LogSuspiciousActivity(ip, "Rate limit excedido", "Demasiadas peticiones en corto tiempo")
}

// LogInvalidInput registra inputs inválidos
func (sl *SecurityLogger) LogInvalidInput(ip, field, value string) {
	sl.LogSuspiciousActivity(ip, "Input inválido", 
		"Campo: "+field+" | Valor sospechoso detectado")
}

// LogXSSAttempt registra intentos de XSS
func (sl *SecurityLogger) LogXSSAttempt(ip, field string) {
	sl.LogSuspiciousActivity(ip, "Posible intento XSS", 
		"Campo: "+field+" | Contenido HTML/Script detectado")
}

// LogEmailSendError registra errores al enviar emails
func (sl *SecurityLogger) LogEmailSendError(ip, email, reason string) {
	if !sl.enabled {
		return
	}
	
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("⚠️ [EMAIL ERROR] %s | IP: %s | Email: %s | Razón: %s", 
		timestamp, ip, email, reason)
}

// LogSuccessfulContact registra contactos exitosos
func (sl *SecurityLogger) LogSuccessfulContact(ip, email, name string) {
	if !sl.enabled {
		return
	}
	
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("✅ [CONTACT] %s | IP: %s | De: %s (%s)", 
		timestamp, ip, name, email)
}
