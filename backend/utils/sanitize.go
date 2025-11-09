package utils

import (
	"html"
	"strings"
)

// SanitizeInput sanitiza el input del usuario para prevenir XSS
func SanitizeInput(input string) string {
	// Escapar caracteres HTML especiales
	sanitized := html.EscapeString(input)
	
	// Eliminar espacios múltiples
	sanitized = strings.Join(strings.Fields(sanitized), " ")
	
	return strings.TrimSpace(sanitized)
}

// SanitizeEmail sanitiza y normaliza un email
func SanitizeEmail(email string) string {
	// Convertir a minúsculas y eliminar espacios
	email = strings.ToLower(strings.TrimSpace(email))
	
	// Escapar caracteres HTML
	return html.EscapeString(email)
}

// StripHTML elimina todas las etiquetas HTML de un string
func StripHTML(input string) string {
	// Lista de caracteres/strings peligrosos
	dangerous := []string{
		"<script", "</script>",
		"<iframe", "</iframe>",
		"javascript:",
		"onerror=", "onload=", "onclick=",
		"<object", "</object>",
		"<embed", "</embed>",
	}

	sanitized := strings.ToLower(input)
	
	// Verificar si contiene elementos peligrosos
	for _, danger := range dangerous {
		if strings.Contains(sanitized, danger) {
			// Si contiene elementos peligrosos, escapar todo el HTML
			return html.EscapeString(input)
		}
	}

	return input
}

// ValidateAndSanitize valida y sanitiza un campo de texto
func ValidateAndSanitize(input string, maxLength int) (string, bool) {
	// Eliminar espacios al inicio y final
	input = strings.TrimSpace(input)
	
	// Verificar longitud
	if len(input) == 0 || len(input) > maxLength {
		return "", false
	}
	
	// Sanitizar
	sanitized := SanitizeInput(input)
	
	return sanitized, true
}
