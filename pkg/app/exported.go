package app

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      app.Model
//    }
type Model struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

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

// AddModels sets structs to be models for database migration
func AddModels(values ...interface{}) {
	for _, value := range values {
		std.models = append(std.models, value)
	}
}

// Db returns the used instance of the GORM Db object
func Db() *gorm.DB {
	return std.db
}

// Run the app
func Run() {
	std.run()
}
