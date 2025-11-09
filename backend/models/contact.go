package models

import (
	"errors"
	"regexp"
	"strings"
)

// ContactRequest representa los datos del formulario de contacto
type ContactRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

// ContactResponse representa la respuesta de la API
type ContactResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Validate valida los datos del formulario de contacto
func (c *ContactRequest) Validate() error {
	// Validar nombres
	if strings.TrimSpace(c.FirstName) == "" {
		return errors.New("el nombre es requerido")
	}
	if len(c.FirstName) > 100 {
		return errors.New("el nombre es demasiado largo")
	}

	// Validar apellidos
	if strings.TrimSpace(c.LastName) == "" {
		return errors.New("los apellidos son requeridos")
	}
	if len(c.LastName) > 100 {
		return errors.New("los apellidos son demasiado largos")
	}

	// Validar email
	if strings.TrimSpace(c.Email) == "" {
		return errors.New("el correo electrónico es requerido")
	}
	if !isValidEmail(c.Email) {
		return errors.New("el correo electrónico no es válido")
	}

	// Validar asunto
	if strings.TrimSpace(c.Subject) == "" {
		return errors.New("el asunto es requerido")
	}
	if len(c.Subject) > 200 {
		return errors.New("el asunto es demasiado largo")
	}

	// Validar mensaje
	if strings.TrimSpace(c.Message) == "" {
		return errors.New("el mensaje es requerido")
	}
	if len(c.Message) < 10 {
		return errors.New("el mensaje es demasiado corto (mínimo 10 caracteres)")
	}
	if len(c.Message) > 5000 {
		return errors.New("el mensaje es demasiado largo")
	}

	return nil
}

// isValidEmail valida el formato del correo electrónico
func isValidEmail(email string) bool {
	// Expresión regular para validar emails
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// GetFullName retorna el nombre completo
func (c *ContactRequest) GetFullName() string {
	return strings.TrimSpace(c.FirstName + " " + c.LastName)
}
