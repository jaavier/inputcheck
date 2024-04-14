package inputcheck

import (
	"regexp"
	"strings"
)

// IsEmail valida si una cadena es una dirección de correo electrónico válida.
func IsEmail(email string) bool {
	// Expresión regular para validar direcciones de correo electrónico
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

// IsPhone valida si una cadena es un número de teléfono válido.
func IsPhone(phone string) bool {
	// Expresión regular para validar números de teléfono en formato internacional
	regex := `^\+[1-9]\d{1,14}$`
	return regexp.MustCompile(regex).MatchString(phone)
}

// IsDomain valida si una cadena es un nombre de dominio válido.
func IsDomain(domain string) bool {
	// Expresión regular para validar nombres de dominio
	regex := `^([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\.)+[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(domain)
}

// IsAlphaNum verifica si una cadena contiene solo caracteres alfanuméricos.
func IsAlphaNum(str string) bool {
	regex := "^[a-zA-Z0-9]+$"
	return regexp.MustCompile(regex).MatchString(str)
}

// IsAlpha verifica si una cadena contiene solo letras del alfabeto.
func IsAlpha(str string) bool {
	regex := "^[a-zA-Z]+$"
	return regexp.MustCompile(regex).MatchString(str)
}

// IsNumeric verifica si una cadena contiene solo caracteres numéricos.
func IsNumeric(str string) bool {
	regex := "^[0-9]+$"
	return regexp.MustCompile(regex).MatchString(str)
}

// IsURL verifica si una cadena es una URL válida.
func IsURL(url string) bool {
	// Expresión regular para validar URLs
	regex := `^(https?|ftp)://[^\s/$.?#].[^\s]*$`
	return regexp.MustCompile(regex).MatchString(url)
}

// IsIPv4 verifica si una cadena es una dirección IPv4 válida.
func IsIPv4(ip string) bool {
	// Expresión regular para validar direcciones IPv4
	regex := `^(\d{1,3}\.){3}\d{1,3}$`
	return regexp.MustCompile(regex).MatchString(ip)
}

// IsIPv6 verifica si una cadena es una dirección IPv6 válida.
func IsIPv6(ip string) bool {
	// Expresión regular para validar direcciones IPv6
	regex := `^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`
	return regexp.MustCompile(regex).MatchString(ip)
}

// IsHexColor verifica si una cadena es un código de color hexadecimal válido.
func IsHexColor(color string) bool {
	regex := "^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$"
	return regexp.MustCompile(regex).MatchString(color)
}

// IsLowerCase verifica si una cadena contiene solo letras minúsculas.
func IsLowerCase(str string) bool {
	if str == "" { // Si la cadena está vacía, no es minúscula
		return false
	}
	return str == strings.ToLower(str)
}

// IsUpperCase verifica si una cadena contiene solo letras mayúsculas.
func IsUpperCase(str string) bool {
	if str == "" { // Si la cadena está vacía, no es mayúscula
		return false
	}
	return str == strings.ToUpper(str)
}

// IsUUID verifica si una cadena es un UUID (identificador único universal) válido.
func IsUUID(uuid string) bool {
	if uuid == "" { // Si el UUID está vacío, no es válido
		return false
	}
	regex := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	return regexp.MustCompile(regex).MatchString(uuid)
}
