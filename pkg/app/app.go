package app

import "github.com/dgarcia202/service-template/internal/cmd"

// App represents the service
type App struct {
	ServiceName      string
	ShortDescription string
	LongDescription  string
	Version          string
}

// New creates the App object
func New() *App {
	app := App{}
	return &app
}

// Run runs the app either bringing up the service or other action like showing version number
func (a App) Run() {
	info := cmd.ServiceInfo{Name: a.ServiceName, Short: a.ShortDescription, Long: a.LongDescription, Version: a.Version}
	cmd.Execute(&info)
}
