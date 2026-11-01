package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

// Определяем location и sensorID по тем же правилам, что в задании
func resolveLocationAndID(location, sensorID string) (string, string) {
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
	return location, sensorID
}

// statusFromValue — простая бизнес-логика, согласованная с монолитом
func statusFromValue(v float64) string {
	switch {
	case v < 18:
		return "cold"
	case v > 27:
		return "hot"
	default:
		return "normal"
	}
}

func generateTemperature(location, sensorID string) TemperatureResponse {
	location, sensorID = resolveLocationAndID(location, sensorID)

	// Случайная температура 18.0 – 28.0 °C с одним знаком после запятой
	temp := 18.0 + rand.Float64()*10.0
	temp = float64(int(temp*10+0.5)) / 10

	return TemperatureResponse{
		Value:       temp,
		Unit:        "celsius",
		Timestamp:   time.Now().UTC(),
		Location:    location,
		Status:      statusFromValue(temp),
		SensorID:    sensorID,
		SensorType:  "temperature",
		Description: "Random temperature from external API",
	}
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	sensorID := r.URL.Query().Get("sensorId") // необязательный параметр
	resp := generateTemperature(location, sensorID)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// /temperature/{sensorID} — вызывается из GetTemperatureByID
func temperatureByIDHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/temperature/")
	sensorID := strings.TrimSpace(path)
	if sensorID == "" {
		http.Error(w, `{"error":"sensorID is required"}`, http.StatusBadRequest)
		return
	}

	// Валидация, что sensorID — число (не обязательно, но полезно)
	if _, err := strconv.Atoi(sensorID); err != nil {
		http.Error(w, `{"error":"invalid sensorID"}`, http.StatusBadRequest)
		return
	}

	resp := generateTemperature("", sensorID)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mux := http.NewServeMux()
	mux.HandleFunc("/temperature", temperatureHandler)       // ?location=...
	mux.HandleFunc("/temperature/", temperatureByIDHandler)  // /temperature/{id}
	mux.HandleFunc("/health", healthHandler)

	srv := &http.Server{
		Addr:         ":8081",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("temperature-api listening on :8081")
	log.Fatal(srv.ListenAndServe())
}