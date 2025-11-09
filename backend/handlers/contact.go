package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xavierdev25/portfolio-backend/models"
	"github.com/xavierdev25/portfolio-backend/services"
)

// ContactHandler maneja las solicitudes de contacto
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud
	var contactReq models.ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&contactReq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Datos inválidos en la solicitud")
		return
	}

	// Validar los datos
	if err := contactReq.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Enviar el correo electrónico
	emailService := services.NewEmailService()
	if err := emailService.SendContactEmail(&contactReq); err != nil {
		log.Printf("Error al enviar el correo: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error al enviar el mensaje. Por favor, intenta más tarde.")
		return
	}

	// Responder con éxito
	respondWithJSON(w, http.StatusOK, models.ContactResponse{
		Success: true,
		Message: "Mensaje enviado correctamente. Te responderé pronto!",
	})
}

// HealthCheckHandler verifica el estado del servidor
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "ok",
		"message": "Portfolio Backend API is running",
		"version": "1.0.0",
	})
}

// respondWithError envía una respuesta de error en formato JSON
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.ContactResponse{
		Success: false,
		Message: message,
	})
}

// respondWithJSON envía una respuesta en formato JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "message": "Error interno del servidor"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
