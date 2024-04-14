package inputcheck

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	// Casos de prueba para IsEmail
	cases := []struct {
		email    string
		expected bool
	}{
		{"correo@example.com", true},
		{"correo@ejemplo", false},
		{"", false}, // Email vacío
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsEmail(c.email)
		if result != c.expected {
			t.Errorf("IsEmail(%q) == %v, expected %v", c.email, result, c.expected)
		}
	}
}

func TestIsPhone(t *testing.T) {
	// Casos de prueba para IsPhone
	cases := []struct {
		phone    string
		expected bool
	}{
		{"+1234567890", true},
		{"123", false},
		{"", false}, // Teléfono vacío
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsPhone(c.phone)
		if result != c.expected {
			t.Errorf("IsPhone(%q) == %v, expected %v", c.phone, result, c.expected)
		}
	}
}

func TestIsDomain(t *testing.T) {
	// Casos de prueba para IsDomain
	cases := []struct {
		domain   string
		expected bool
	}{
		{"example.com", true},
		{"-example.com", false},
		{"", false}, // Dominio vacío
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsDomain(c.domain)
		if result != c.expected {
			t.Errorf("IsDomain(%q) == %v, expected %v", c.domain, result, c.expected)
		}
	}
}

func TestIsLowerCase(t *testing.T) {
	// Casos de prueba para IsLowerCase
	cases := []struct {
		str      string
		expected bool
	}{
		{"abc", true},
		{"Abc", false},
		{"", false}, // Cadena vacía
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsLowerCase(c.str)
		if result != c.expected {
			t.Errorf("IsLowerCase(%q) == %v, expected %v", c.str, result, c.expected)
		}
	}
}

func TestIsUpperCase(t *testing.T) {
	// Casos de prueba para IsUpperCase
	cases := []struct {
		str      string
		expected bool
	}{
		{"ABC", true},
		{"AbC", false},
		{"", false}, // Cadena vacía
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsUpperCase(c.str)
		if result != c.expected {
			t.Errorf("IsUpperCase(%q) == %v, expected %v", c.str, result, c.expected)
		}
	}
}

func TestIsAlphaNum(t *testing.T) {
	// Casos de prueba para IsAlphaNum
	cases := []struct {
		str      string
		expected bool
	}{
		{"abc123", true},
		{"abc", true},
		{"123", true},
		{"abc$123", false},
		{"", false}, // Cadena vacía
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsAlphaNum(c.str)
		if result != c.expected {
			t.Errorf("IsAlphaNum(%q) == %v, expected %v", c.str, result, c.expected)
		}
	}
}

func TestIsAlpha(t *testing.T) {
	// Casos de prueba para IsAlpha
	cases := []struct {
		str      string
		expected bool
	}{
		{"abc", true},
		{"ABC", true},
		{"abc123", false},
		{"123", false},
		{"", false}, // Cadena vacía
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsAlpha(c.str)
		if result != c.expected {
			t.Errorf("IsAlpha(%q) == %v, expected %v", c.str, result, c.expected)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	// Casos de prueba para IsNumeric
	cases := []struct {
		str      string
		expected bool
	}{
		{"123", true},
		{"123abc", false},
		{"", false}, // Cadena vacía
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsNumeric(c.str)
		if result != c.expected {
			t.Errorf("IsNumeric(%q) == %v, expected %v", c.str, result, c.expected)
		}
	}
}

func TestIsURL(t *testing.T) {
	// Casos de prueba para IsURL
	cases := []struct {
		url      string
		expected bool
	}{
		{"https://www.example.com", true},
		{"http://example.com", true},
		{"ftp://example.com", true},
		{"example.com", false},
		{"", false}, // URL vacía
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsURL(c.url)
		if result != c.expected {
			t.Errorf("IsURL(%q) == %v, expected %v", c.url, result, c.expected)
		}
	}
}

func TestIsIPv4(t *testing.T) {
	// Casos de prueba para IsIPv4
	cases := []struct {
		ip       string
		expected bool
	}{
		{"192.168.0.1", true},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", false}, // IPv6
		{"", false}, // IP vacía
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsIPv4(c.ip)
		if result != c.expected {
			t.Errorf("IsIPv4(%q) == %v, expected %v", c.ip, result, c.expected)
		}
	}
}

func TestIsIPv6(t *testing.T) {
	// Casos de prueba para IsIPv6
	cases := []struct {
		ip       string
		expected bool
	}{
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"192.168.0.1", false}, // IPv4
		{"", false},            // IP vacía
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsIPv6(c.ip)
		if result != c.expected {
			t.Errorf("IsIPv6(%q) == %v, expected %v", c.ip, result, c.expected)
		}
	}
}

func TestIsHexColor(t *testing.T) {
	// Casos de prueba para IsHexColor
	cases := []struct {
		color    string
		expected bool
	}{
		{"#FF0000", true},
		{"#abc", true},
		{"#1234567", false}, // Demasiados dígitos
		{"abc123", false},   // No comienza con #
		{"", false},         // Color vacío
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsHexColor(c.color)
		if result != c.expected {
			t.Errorf("IsHexColor(%q) == %v, expected %v", c.color, result, c.expected)
		}
	}
}

func TestIsUUID(t *testing.T) {
	// Casos de prueba para IsUUID
	cases := []struct {
		uuid     string
		expected bool
	}{
		{"f47ac10b-58cc-4372-a567-0e02b2c3d479", true},
		{"F47AC10B-58CC-4372-A567-0E02B2C3D479", true},
		{"f47ac10b-58cc-4372-a567-0e02b2c3d4791", false}, // Demasiados dígitos
		{"f47ac10b-58cc-4372-a567-0e02b2c3d47", false},   // Demasiados dígitos
		{"", false}, // UUID vacío
		// Agrega más casos de prueba según sea necesario...
	}

	for _, c := range cases {
		result := IsUUID(c.uuid)
		if result != c.expected {
			t.Errorf("IsUUID(%q) == %v, expected %v", c.uuid, result, c.expected)
		}
	}
}
