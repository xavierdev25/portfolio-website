package services

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/xavierdev25/portfolio-backend/models"
)

// EmailService maneja el envío de correos electrónicos
type EmailService struct {
	smtpHost      string
	smtpPort      string
	smtpEmail     string
	smtpPassword  string
	recipientEmail string
}

// NewEmailService crea una nueva instancia del servicio de email
func NewEmailService() *EmailService {
	return &EmailService{
		smtpHost:      os.Getenv("SMTP_HOST"),
		smtpPort:      os.Getenv("SMTP_PORT"),
		smtpEmail:     os.Getenv("SMTP_EMAIL"),
		smtpPassword:  os.Getenv("SMTP_PASSWORD"),
		recipientEmail: os.Getenv("RECIPIENT_EMAIL"),
	}
}

// SendContactEmail envía un correo electrónico con los datos del formulario de contacto
func (e *EmailService) SendContactEmail(contact *models.ContactRequest) error {
	// Validar configuración SMTP
	if err := e.validateConfig(); err != nil {
		return err
	}

	// Construir el mensaje
	subject := fmt.Sprintf("Nuevo mensaje de contacto: %s", contact.Subject)
	body := e.buildEmailBody(contact)

	// Preparar el mensaje en formato MIME
	message := e.buildMIMEMessage(subject, body)

	// Configurar autenticación
	auth := smtp.PlainAuth("", e.smtpEmail, e.smtpPassword, e.smtpHost)

	// Enviar el correo
	addr := fmt.Sprintf("%s:%s", e.smtpHost, e.smtpPort)
	to := []string{e.recipientEmail}

	err := smtp.SendMail(addr, auth, e.smtpEmail, to, []byte(message))
	if err != nil {
		return fmt.Errorf("error al enviar el correo: %w", err)
	}

	return nil
}

// validateConfig valida que todas las configuraciones necesarias estén presentes
func (e *EmailService) validateConfig() error {
	if e.smtpHost == "" {
		return fmt.Errorf("SMTP_HOST no está configurado")
	}
	if e.smtpPort == "" {
		return fmt.Errorf("SMTP_PORT no está configurado")
	}
	if e.smtpEmail == "" {
		return fmt.Errorf("SMTP_EMAIL no está configurado")
	}
	if e.smtpPassword == "" {
		return fmt.Errorf("SMTP_PASSWORD no está configurado")
	}
	if e.recipientEmail == "" {
		return fmt.Errorf("RECIPIENT_EMAIL no está configurado")
	}
	return nil
}

// buildEmailBody construye el cuerpo del correo electrónico en HTML
func (e *EmailService) buildEmailBody(contact *models.ContactRequest) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background: linear-gradient(135deg, #E75DEE 0%%, #C93DD1 100%%); color: white; padding: 30px; border-radius: 10px 10px 0 0; }
        .header h1 { margin: 0; font-size: 24px; }
        .content { background: #f9f9f9; padding: 30px; border-radius: 0 0 10px 10px; }
        .field { margin-bottom: 20px; }
        .label { font-weight: bold; color: #E75DEE; margin-bottom: 5px; }
        .value { background: white; padding: 12px; border-radius: 5px; border-left: 3px solid #E75DEE; }
        .message { background: white; padding: 20px; border-radius: 5px; border-left: 3px solid #E75DEE; white-space: pre-wrap; }
        .footer { text-align: center; margin-top: 20px; color: #666; font-size: 12px; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>📧 Nuevo Mensaje de Contacto</h1>
        </div>
        <div class="content">
            <div class="field">
                <div class="label">👤 Nombre Completo:</div>
                <div class="value">%s</div>
            </div>
            <div class="field">
                <div class="label">📧 Correo Electrónico:</div>
                <div class="value"><a href="mailto:%s">%s</a></div>
            </div>
            <div class="field">
                <div class="label">📋 Asunto:</div>
                <div class="value">%s</div>
            </div>
            <div class="field">
                <div class="label">💬 Mensaje:</div>
                <div class="message">%s</div>
            </div>
            <div class="footer">
                <p>Este mensaje fue enviado desde el formulario de contacto de tu portafolio</p>
            </div>
        </div>
    </div>
</body>
</html>
	`, 
		contact.GetFullName(),
		contact.Email,
		contact.Email,
		contact.Subject,
		strings.ReplaceAll(contact.Message, "\n", "<br>"),
	)
}

// buildMIMEMessage construye el mensaje en formato MIME
func (e *EmailService) buildMIMEMessage(subject, body string) string {
	headers := make(map[string]string)
	headers["From"] = e.smtpEmail
	headers["To"] = e.recipientEmail
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	return message
}
