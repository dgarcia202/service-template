#Service template

* cli interface & local configuration with cobra & viper
* HTTP Rest with GIN
* Logging with logrus
* Relational database support with GORM
* MongoDB support?
* Service discovery?
* Centralized configuration?
* Authentication?
* RabbitMQ??
* health info??

## Notes

* try logger inside handlers
* optimize app.Model for UUID fields etc. see whats best to store UUID in different dialects
* try to redirect GORM log to main log