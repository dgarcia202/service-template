package app

import (
	"github.com/jinzhu/gorm"
)

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

// AddHTTPSetup allows to modify HTTP server configuration
func AddHTTPSetup(fn HTTPSetupFunc) {
	std.addHTTPSetup(fn)
}

// AddModel sets an struct to be model for database migration
func AddModel(value interface{}) {
	std.addModel(value)
}

// Db returns the used instance of the GORM Db object
func Db() *gorm.DB {
	return std.db
}

// Run the app
func Run() {
	std.run()
}
