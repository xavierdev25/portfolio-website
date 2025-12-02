package models

import (
	"errors"
	"regexp"
	"strings"

	"github.com/xavierdev25/portfolio-backend/utils"
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

// Validate valida y sanitiza los datos del formulario de contacto
func (c *ContactRequest) Validate() error {
	// Sanitizar y validar nombres
	sanitized, valid := utils.ValidateAndSanitize(c.FirstName, 100)
	if !valid || sanitized == "" {
		return errors.New("el nombre es requerido y debe tener máximo 100 caracteres")
	}
	c.FirstName = sanitized

	// Sanitizar y validar apellidos
	sanitized, valid = utils.ValidateAndSanitize(c.LastName, 100)
	if !valid || sanitized == "" {
		return errors.New("los apellidos son requeridos y deben tener máximo 100 caracteres")
	}
	c.LastName = sanitized

	// Sanitizar y validar email
	c.Email = utils.SanitizeEmail(c.Email)
	if c.Email == "" {
		return errors.New("el correo electrónico es requerido")
	}
	if !isValidEmail(c.Email) {
		return errors.New("el correo electrónico no es válido")
	}

	// Sanitizar y validar asunto
	sanitized, valid = utils.ValidateAndSanitize(c.Subject, 200)
	if !valid || sanitized == "" {
		return errors.New("el asunto es requerido y debe tener máximo 200 caracteres")
	}
	c.Subject = sanitized

	// Sanitizar y validar mensaje
	c.Message = utils.StripHTML(c.Message)
	c.Message = strings.TrimSpace(c.Message)
	
	if c.Message == "" {
		return errors.New("el mensaje es requerido")
	}
	if len(c.Message) < 10 {
		return errors.New("el mensaje es demasiado corto (mínimo 10 caracteres)")
	}
	if len(c.Message) > 5000 {
		return errors.New("el mensaje es demasiado largo (máximo 5000 caracteres)")
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
