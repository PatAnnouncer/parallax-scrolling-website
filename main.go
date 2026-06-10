package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

// Структура ответа для фронтенда
type ServerStatus struct {
	Minecraft bool `json:"minecraft"`
	DnD       bool `json:"dnd"`
}

// Функция проверки доступности порта
func checkPort(host string, port string) bool {
	timeout := time.Second * 3
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return false
	}
	if conn != nil {
		defer conn.Close()
		return true
	}
	return false
}

// Обработчик эндпоинта /api/status
func statusHandler(w http.ResponseWriter, r *http.Request) {
	// ==========================================
	// ВНИМАНИЕ: ЗАМЕНИ НА СВОИ IP И ПОРТЫ!
	// ==========================================
	mcStatus := checkPort("8.8.8.8", "25565")  // Пример IP и порта для Minecraft
	dndStatus := checkPort("8.8.8.8", "30100") // Пример IP и порта для D&D / Foundry VTT

	status := ServerStatus{
		Minecraft: mcStatus,
		DnD:       dndStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func main() {
	// Раздаем все файлы из твоей папки ready-html
	fs := http.FileServer(http.Dir("./ready-html"))
	http.Handle("/", fs)

	// Роут для получения статуса
	http.HandleFunc("/api/status", statusHandler)

	fmt.Println("Сервер Fairy Forest запущен на порту :8080")
	// Запуск веб-сервера на стандартном порту
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
