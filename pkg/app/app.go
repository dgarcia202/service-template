package app

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/dgarcia202/service-template/internal/cmd"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	consul "github.com/hashicorp/consul/api"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// HTTPSetupFunc is a function that is able to configure the web server in terms
// of routes and middleware
type HTTPSetupFunc func(*gin.Engine)

// App represents the service
type App struct {
	serviceName      string
	shortDescription string
	longDescription  string
	version          string
	consulID         string
	consulTTL        time.Duration

	httpSetupFuncs []HTTPSetupFunc
	models         []interface{}

	consulAgent *consul.Agent
	ginEngine   *gin.Engine
	db          *gorm.DB
}

var std = NewApp()

// NewApp constructs a new App object
func NewApp() *App {

	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		log.Fatal("Can't create a consul client")
	}

	duration, _ := time.ParseDuration("10s") // TODO: move this to a constant or config

	return &App{
		consulAgent: c.Agent(),
		consulTTL:   duration}
}

// runs the app either bringing up the service or other action like showing version number
func (a *App) run() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig.String() == "interrupt" {
				shutDown()
			}
		}
	}()

	info := cmd.ServiceInfo{Name: a.serviceName, Short: a.shortDescription, Long: a.longDescription, Version: a.version}
	cmd.Execute(&info, startUp)
}

// setupRoutes allows to modify routing configuration
func (a *App) addHTTPSetup(fn HTTPSetupFunc) {
	a.httpSetupFuncs = append(a.httpSetupFuncs, fn)
}

func (a *App) registerConsulService() {

	log.Debug("Registering service in consul")

	if a.consulAgent == nil {
		log.Warn("No consul agent available, skipping service registration")
		return
	}

	portInt, err := strconv.Atoi(viper.GetString("port"))
	if err != nil {
		log.Fatal("Can't parse port value ", viper.GetString("port"))
	}

	serviceDef := &consul.AgentServiceRegistration{
		Name:    a.serviceName,
		ID:      uuid.New().String(),
		Address: viper.GetString("address"),
		Port:    portInt,
		Check: &consul.AgentServiceCheck{
			CheckID: uuid.New().String(),
			Name:    fmt.Sprintf("checkTTL-%s", a.serviceName),
			TTL:     a.consulTTL.String()}}

	if err = a.consulAgent.ServiceRegister(serviceDef); err != nil {
		log.Error("Can't register service: ", err.Error())
		return
	}

	a.consulID = serviceDef.ID

	go a.heartbeat(serviceDef.Check.CheckID)
}

func (a *App) deregisterConsulService() {

	if len(a.consulID) == 0 {
		log.Warn("No service ID available, skipping consul deregistration")
		return
	}

	log.Debug("Deregistering service in consul")

	if err := a.consulAgent.ServiceDeregister(a.consulID); err != nil {
		log.Error("Error while deregistering service with consul: ", err)
	}
}

func (a *App) heartbeat(checkID string) {
	log.Trace("Started heartbeat routine")
	ticker := time.NewTicker(a.consulTTL / 2)
	for range ticker.C {
		log.Trace("Sending heartbeat")
		if err := a.consulAgent.UpdateTTL(checkID, "Service heartbeat...", "pass"); err != nil {
			log.Error("Error sending heartbeat: ", err.Error())
		}
	}
}
