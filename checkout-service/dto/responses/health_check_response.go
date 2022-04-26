package responses

// Contains a Response DTO of the health of the various components of the service.
type HealthCheckResponse struct {
	ServiceHealth     string `json:"service_health"`      // Service Health
	KafkaServerHealth string `json:"kafka_server_health"` // Kafka Server Health
	DBHealth          string `json:"db_health"`           // DB Health
}
