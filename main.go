package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/yaml.v2"
)

// ConfigFile to watch
const ConfigFile = "/config/config.yaml"

// Bind adress to listen to
const Bind = "0.0.0.0:8080"

func check(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

// Config This is the struct that holds our application's configuration
type Config struct {
	Message string `yaml:"message"`
}

/*
 Simple Yaml Config file loader
*/
func loadConfig(configFile string) *Config {
	conf := &Config{}
	configData, err := ioutil.ReadFile(configFile)
	check(err)

	err = yaml.Unmarshal(configData, conf)
	check(err)
	return conf
}

func main() {
	confManager := NewMutexConfigManager(loadConfig(ConfigFile))

	// Create a single GET Handler to print out our simple config message
	router := httprouter.New()
	router.GET("/", func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
		conf := confManager.Get()
		fmt.Fprintf(resp, "%s", conf.Message)
	})

	// Watch the file for modification and update the config manager with the new config when it's available
	watcher, err := WatchFile(ConfigFile, time.Second, func() {
		fmt.Printf("Configfile Updated\n")
		conf := loadConfig(ConfigFile)
		confManager.Set(conf)
	})
	check(err)

	// Clean up
	defer func() {
		watcher.Close()
		confManager.Close()
	}()

	fmt.Printf("Listening on '%s'....\n", Bind)
	err = http.ListenAndServe(Bind, router)
	check(err)
}
