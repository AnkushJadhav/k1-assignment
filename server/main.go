package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AnkushJadhav/k1-assignment/server/pkg/transport/http"

	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"
	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance/mysql"
)

func main() {
	if err := bootup(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type configDetails struct {
	defaultValue string
	isMandatory  bool
	value        *string
	usage        string
}

func bootup() error {
	// init app config
	ac := generateAppConfig()
	// validate input flags
	if err := validateMandatoryFlags(&ac); err != nil {
		return err
	}
	// create db client
	store, err := getStore(*ac["dsnstring"].value)
	if err != nil {
		return err
	}
	// create server
	srvr, err := http.New(*ac["bindip"].value, *ac["bindport"].value)
	if err != nil {
		return err
	}
	srvr.SetStore(store)
	srvr.SetAuth(*ac["authsecret"].value)
	srvr.AttachRoutes()
	// start server
	err = srvr.Start()
	if err != nil {
		return err
	}

	return nil
}

func generateAppConfig() map[string]*configDetails {
	ac := make(map[string]*configDetails)

	// dsnstring
	ac["dsnstring"] = &configDetails{
		defaultValue: "",
		isMandatory:  true,
		usage:        "db connection string",
	}

	// authsecret
	ac["authsecret"] = &configDetails{
		defaultValue: "",
		isMandatory:  true,
		usage:        "secret string for JWT auth",
	}

	// bindip
	ac["bindip"] = &configDetails{
		defaultValue: "0.0.0.0",
		isMandatory:  false,
		usage:        "ip to bind to",
	}

	//bind port
	ac["bindport"] = &configDetails{
		defaultValue: "8080",
		isMandatory:  false,
		usage:        "port to bind to",
	}

	populateAppConfig(&ac)
	return ac
}

func validateMandatoryFlags(ac *map[string]*configDetails) error {
	for prop, val := range *ac {
		if val.isMandatory && *val.value == val.defaultValue {
			return fmt.Errorf(prop + " not provided")
		}
	}

	return nil
}

func populateAppConfig(ac *map[string]*configDetails) {
	for prop, val := range *ac {
		val.value = flag.String(prop, val.defaultValue, val.usage)
	}
	flag.Parse()
}

func getStore(dsn string) (persistance.Client, error) {
	return mysql.New(dsn)
}
