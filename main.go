package main

import (
	"flag"

	v1 "github.com/njxxdev/leshy-sample/pkg/v1"
	"github.com/njxxdev/leshy/pkg/api"
	"github.com/njxxdev/leshy/pkg/component"
	"github.com/njxxdev/leshy/pkg/config"
)

func loadModules() {
	// component.GetComponentManager().Append(storages.NewPostgresRepository("postgres"))
	v1.CreateApiV1()
}

func runServices() error {
	service, err := component.GetComponentManager().GetComponent("api_v1")
	if err != nil {
		panic("Can't found component \"api_v1\": " + err.Error())
	}

	// Run API
	errChan := make(chan error)
	/// New API
	go func() {
		errChan <- service.(*api.APIServer).Run()
	}()

	return <-errChan
}

func main() {
	configFilename := flag.String("config", "./configs/prod.yaml", "static config path")
	_ = flag.String("vars", "./configs/vars.yaml", "static config vars path")
	flag.Parse()
	config.LoadConfigs(*configFilename)

	loadModules()

	if runServices() != nil {
		panic("Servers down")
	}
}
