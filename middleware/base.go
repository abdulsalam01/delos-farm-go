package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response is the generic response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// JSONMiddleware is the middleware function that wraps the handlerFunc
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Create a new Response instance
		response := Response{
			Success: true,
		}

		// Wrap the original ResponseWriter with our custom implementation
		wrappedWriter := NewResponseWriter(w)
		// Call the next handler in the chain
		next.ServeHTTP(wrappedWriter, r)

		// Check if an error occurred
		if wrappedWriter.StatusCode >= http.StatusBadRequest {
			response.Success = false
			response.Message = http.StatusText(wrappedWriter.StatusCode)
		}

		// Set the response status code
		w.WriteHeader(wrappedWriter.StatusCode)

		// Store the response data in the Response struct
		response.Data = wrappedWriter.ResponseData

		// Marshal the response to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println("Error marshaling JSON response:", err)
			return
		}

		// Write the JSON response to the client
		w.Write(jsonResponse)
	})
}

// ResponseWriter is a custom implementation of http.ResponseWriter
type ResponseWriter struct {
	http.ResponseWriter
	StatusCode   int
	ResponseData interface{}
}

// NewResponseWriter creates a new instance of ResponseWriter
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}
