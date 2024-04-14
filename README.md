# inputcheck

[![Go Reference](https://pkg.go.dev/badge/github.com/jaavier/inputcheck.svg)](https://pkg.go.dev/github.com/jaavier/inputcheck)

Package inputcheck provides advanced input validation for Go applications, including IPv4, IPv6, domain validation, and more.

## Install

`go get -u https://github.com/jaavier/inputcheck`


## Import

```go
package main

import (
	"github.com/jaavier/inputcheck"
)

func main() {
    // Your code here...
}

```
## Features

### Email Validation

```go
email := "user@example.com"
if inputcheck.IsEmail(email) {
    fmt.Println("Valid email address")
} else {
    fmt.Println("Invalid email address")
}
```

### Phone Number Validation

```go
phone := "+1234567890"
if inputcheck.IsPhone(phone) {
    fmt.Println("Valid phone number")
} else {
    fmt.Println("Invalid phone number")
}
```

### Domain Validation

```go
domain := "example.com"
if inputcheck.IsDomain(domain) {
    fmt.Println("Valid domain name")
} else {
    fmt.Println("Invalid domain name")
}
```

### Alphanumeric Validation

```go
alphanumeric := "abc123"
if inputcheck.IsAlphaNum(alphanumeric) {
    fmt.Println("Valid alphanumeric string")
} else {
    fmt.Println("Invalid alphanumeric string")
}
```

### Alpha Validation

```go
alpha := "abc"
if inputcheck.IsAlpha(alpha) {
    fmt.Println("Valid alphabetical string")
} else {
    fmt.Println("Invalid alphabetical string")
}
```

### Numeric Validation

```go
numeric := "123"
if inputcheck.IsNumeric(numeric) {
    fmt.Println("Valid numeric string")
} else {
    fmt.Println("Invalid numeric string")
}
```

### URL Validation

```go
url := "https://www.example.com"
if inputcheck.IsURL(url) {
    fmt.Println("Valid URL")
} else {
    fmt.Println("Invalid URL")
}
```

### IPv4 Validation

```go
ipv4 := "192.168.0.1"
if inputcheck.IsIPv4(ipv4) {
    fmt.Println("Valid IPv4 address")
} else {
    fmt.Println("Invalid IPv4 address")
}
```

### IPv6 Validation

```go
ipv6 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
if inputcheck.IsIPv6(ipv6) {
    fmt.Println("Valid IPv6 address")
} else {
    fmt.Println("Invalid IPv6 address")
}
```

### Hex Color Validation

```go
color := "#FF0000"
if inputcheck.IsHexColor(color) {
    fmt.Println("Valid hexadecimal color")
} else {
    fmt.Println("Invalid hexadecimal color")
}
```

### Case Sensitivity Validation

```go
lower := "lowercase"
if inputcheck.IsLowerCase(lower) {
    fmt.Println("Contains only lowercase characters")
} else {
    fmt.Println("Does not contain only lowercase characters")
}

upper := "UPPERCASE"
if inputcheck.IsUpperCase(upper) {
    fmt.Println("Contains only uppercase characters")
} else {
    fmt.Println("Does not contain only uppercase characters")
}
```

### UUID Validation

```go
uuid := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
if inputcheck.IsUUID(uuid) {
    fmt.Println("Valid UUID")
} else {
    fmt.Println("Invalid UUID")
}
```

## Installation

```bash
go get github.com/jaavier/inputcheck
```

## Contributing

Contributions are welcome! Feel free to open issues or pull requests for any improvements or new features you'd like to see.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.