package middleware

import (
	"net/http"
	"sync"
)

// EndpointStats represents the statistics for a specific endpoint
type EndpointStats struct {
	Count            int            `json:"count"`
	UniqueUserAgents map[string]int `json:"unique_user_agents"`
}

// Statistics represents the overall statistics for all endpoints
type Statistics struct {
	Endpoints map[string]EndpointStats `json:"endpoints"`
	Mutex     sync.RWMutex
}

var stats Statistics

// Middleware function to track endpoint statistics
func TrackStatistics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Increment count for the current endpoint
		stats.Mutex.Lock()

		pathKey := r.URL.Path
		userAgents := r.UserAgent()
		count := stats.Endpoints[pathKey].Count + 1

		// Check userAgents.
		// Initializer the map if empty.
		data, ok := stats.Endpoints[pathKey].UniqueUserAgents[userAgents]
		if !ok {
			stats.Endpoints[pathKey] = EndpointStats{
				Count: count,
				UniqueUserAgents: map[string]int{
					userAgents: 1,
				},
			}
			// Re-assign.
			data = stats.Endpoints[pathKey].UniqueUserAgents[userAgents]
		}

		stats.Endpoints[pathKey] = EndpointStats{
			Count: count,
			UniqueUserAgents: map[string]int{
				userAgents: data,
			},
		}
		stats.Mutex.Unlock()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Handler function to retrieve the statistics
func GetStatistics(w http.ResponseWriter, r *http.Request) {
	stats.Mutex.RLock()
	defer stats.Mutex.RUnlock()

	// Return the statistics in JSON format.
	responseWriter := w.(*ResponseWriter)
	responseWriter.ResponseData = stats.Endpoints
}

// Initialize the statistics map
func InitStatistics() {
	stats.Endpoints = make(map[string]EndpointStats)
}
