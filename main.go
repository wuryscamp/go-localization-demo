package main

import (
	"fmt"
	"os"

	"github.com/Bhinneka/go-localization-demo/config/spik"
)

// en_US
// in_ID
func main() {
	var locale string

	lang := os.Getenv("lang")

	if len(lang) > 0 {
		locale = lang
	} else {
		locale = "en_US"
	}

	langFiles := []string{"config/spik/en_US.yaml", "config/spik/in_ID.yaml"}

	// Create a Bundle to use for the lifetime of your application
	app, err := spik.New(langFiles)
	if err != nil {
		panic(err)
	}

	welcomeMsg, err := app.Translate("welcome", locale)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(welcomeMsg)

}
