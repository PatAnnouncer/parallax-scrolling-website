package handlers

import (
	"encoding/json"
	"net/http"

	"parallax-scrolling/internal/models"
	"parallax-scrolling/internal/services"
)

// StatusHandler handles the /api/status endpoint
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// ==========================================
	// ВНИМАНИЕ: ЗАМЕНИ НА СВОИ IP И ПОРТЫ!
	// ==========================================
	mcStatus := services.CheckPort("8.8.8.8", "25565")  // Пример IP и порта для Minecraft
	dndStatus := services.CheckPort("8.8.8.8", "30100") // Пример IP и порта для D&D / Foundry VTT

	status := models.ServerStatus{
		Minecraft: mcStatus,
		DnD:       dndStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
