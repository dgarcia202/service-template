package app

import "github.com/gin-gonic/gin"

// ServiceName sets a short service name, avoid spaces
func ServiceName(name string) {
	std.serviceName = name
}

// ShortDescription sets a short description text for the service
func ShortDescription(short string) {
	std.shortDescription = short
}

// LongDescription sets bigger description test for the service
func LongDescription(long string) {
	std.longDescription = long
}

// Version sets service version tag
func Version(version string) {
	std.version = version
}

// SetupRoutes allows to modify routing configuration
func SetupRoutes(fn func(*gin.Engine)) {
	std.setupRoutes(fn)
}

// Run the app
func Run() {
	std.run()
}
