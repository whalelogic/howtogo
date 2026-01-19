// Package middleware provides HTTP middleware functions
package middleware

import (	
	"net"
	"net/http"
	"strings"
)


// get client IP address from request
func getClientIP(r *http.Request) string {
	// X-Forwarded-For header? Check first
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		// Return first IP 
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// Check X-Real-IP header
	xrip := r.Header.Get("X-Real-IP")
	if xrip != "" {
		return strings.TrimSpace(xrip)
	}

	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}


// ClientIPMiddleware is a middleware that adds the client IP address to the request context
func ClientIPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := getClientIP(r)
		// Store the client IP in request context or headers as needed
		r.Header.Set("Client-IP", clientIP)
		next.ServeHTTP(w, r)
	})
}




