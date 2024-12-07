package config

func GetServiceDomain(service string) string {
	switch service {
	case "user":
		return "http://localhost:3000/user"
	case "product":
		return "http://localhost:3001/product"
	case "sale":
		return "http://localhost:3002"
	default:
		return ""
	}
}
