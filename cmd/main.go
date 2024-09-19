package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Configuration for backend services
type Service struct {
	Name string
	URL  string
}

var services = []Service{
	{Name: "users", URL: "http://localhost:4001"},
	{Name: "orders", URL: "http://localhost:4002"},
	{Name: "products", URL: "http://localhost:4003"},
}

func main() {
	r := chi.NewRouter()

	// Global Middlewares
	r.Use(middleware.Logger)            // Logs each request
	r.Use(middleware.Recoverer)         // Recovers from panics
	r.Use(middleware.Timeout(60 * 1e9)) // Sets a timeout of 60 seconds

	// CORS Middleware (Optional: Uncomment if needed)
	/*
		r.Use(middleware.AllowContentType("application/json"))
		r.Use(middleware.CORS(
			middleware.CORSOptions{
				AllowedOrigins: []string{"*"},
				AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			},
		))
	*/

	// Root Route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API internal"))
	})

	// Health Check Route
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Dynamically set up proxy routes for each service
	for _, service := range services {
		svc := service // capture range variable
		r.Route("/"+svc.Name, func(r chi.Router) {
			// Subrouter Middlewares (Optional: Add service-specific middleware)
			r.Use(serviceMiddleware(svc.Name))

			// Proxy all subpaths to the respective service
			r.Handle("/*", proxyHandler(svc.URL))
		})
	}

	// Start the server
	port := ":3000"
	fmt.Printf("API internal is listening on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// serviceMiddleware is an example middleware for specific services
func serviceMiddleware(serviceName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Example: Add a header to indicate which service is handling the request
			w.Header().Set("X-Served-By", serviceName)
			// Example: Implement authentication or other checks here
			// For instance:
			/*
				token := r.Header.Get("Authorization")
				if token != "expected-token" {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
			*/
			next.ServeHTTP(w, r)
		})
	}
}

// proxyHandler creates a reverse proxy to the given target
func proxyHandler(target string) http.Handler {
	// Parse the target URL
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid target URL %s: %v", target, err)
	}

	// Create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Modify the request to direct it to the target
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// Optionally, modify the request here (e.g., add headers)
		req.Host = targetURL.Host
	}

	// Optional: Customize the errors handler
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy errors: %v", err)
		http.Error(w, "Bad internal", http.StatusBadGateway)
	}

	return proxy
}
