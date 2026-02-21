package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	zabbix "github.com/tpretz/go-zabbix-api"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/version", handleVersion)
	http.HandleFunc("/api/hosts", handleHosts)
	http.HandleFunc("/api/groups", handleGroups)
	http.HandleFunc("/health", handleHealth)

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	response := APIResponse{
		Success: true,
		Data: map[string]string{
			"service": "Zabbix API Server",
			"version": "1.0.0",
			"endpoints": "/api/version, /api/hosts, /api/groups, /health",
		},
	}
	writeJSON(w, response)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := APIResponse{
		Success: true,
		Data:    map[string]string{"status": "healthy"},
	}
	writeJSON(w, response)
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	zabbixURL := r.URL.Query().Get("url")
	if zabbixURL == "" {
		zabbixURL = os.Getenv("ZABBIX_URL")
	}
	if zabbixURL == "" {
		writeError(w, "Zabbix URL is required")
		return
	}

	config := zabbix.Config{
		Url: zabbixURL,
	}
	api := zabbix.NewAPI(config)

	version, err := api.Version()
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to get version: %v", err))
		return
	}

	response := APIResponse{
		Success: true,
		Data:    map[string]string{"version": version},
	}
	writeJSON(w, response)
}

func handleHosts(w http.ResponseWriter, r *http.Request) {
	zabbixURL := r.URL.Query().Get("url")
	user := r.URL.Query().Get("user")
	password := r.URL.Query().Get("password")

	if zabbixURL == "" {
		zabbixURL = os.Getenv("ZABBIX_URL")
	}
	if user == "" {
		user = os.Getenv("ZABBIX_USER")
	}
	if password == "" {
		password = os.Getenv("ZABBIX_PASSWORD")
	}

	if zabbixURL == "" || user == "" || password == "" {
		writeError(w, "Zabbix URL, user and password are required")
		return
	}

	config := zabbix.Config{
		Url: zabbixURL,
	}
	api := zabbix.NewAPI(config)

	_, err := api.Login(user, password)
	if err != nil {
		writeError(w, fmt.Sprintf("Login failed: %v", err))
		return
	}

	hosts, err := api.HostsGet(zabbix.Params{})
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to get hosts: %v", err))
		return
	}

	response := APIResponse{
		Success: true,
		Data:    hosts,
	}
	writeJSON(w, response)
}

func handleGroups(w http.ResponseWriter, r *http.Request) {
	zabbixURL := r.URL.Query().Get("url")
	user := r.URL.Query().Get("user")
	password := r.URL.Query().Get("password")

	if zabbixURL == "" {
		zabbixURL = os.Getenv("ZABBIX_URL")
	}
	if user == "" {
		user = os.Getenv("ZABBIX_USER")
	}
	if password == "" {
		password = os.Getenv("ZABBIX_PASSWORD")
	}

	if zabbixURL == "" || user == "" || password == "" {
		writeError(w, "Zabbix URL, user and password are required")
		return
	}

	config := zabbix.Config{
		Url: zabbixURL,
	}
	api := zabbix.NewAPI(config)

	_, err := api.Login(user, password)
	if err != nil {
		writeError(w, fmt.Sprintf("Login failed: %v", err))
		return
	}

	groups, err := api.HostGroupsGet(zabbix.Params{})
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to get groups: %v", err))
		return
	}

	response := APIResponse{
		Success: true,
		Data:    groups,
	}
	writeJSON(w, response)
}

func writeJSON(w http.ResponseWriter, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func writeError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(APIResponse{
		Success: false,
		Error:   message,
	})
}
