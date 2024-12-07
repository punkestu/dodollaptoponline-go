package config

func GetServiceDomain(service string) string {
	switch service {
	case "user":
		return "http://localhost:3000"
	case "product":
		return "http://localhost:3001"
	case "sale":
		return "http://localhost:3002"
	default:
		return ""
	}
}
