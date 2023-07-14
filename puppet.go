package webshot

import (
	"errors"

	"github.com/tebeka/selenium"
)

var (
	// ChromeBrowser for the chrome browser
	ChromeBrowser = "chrome"
	// FirefoxBrowser for the firefox browser
	FirefoxBrowser = "firefox"
	// BrowserError is the error returned when the browser type entered by the user is not firefox
	ErrBrowser = errors.New("the supported browser is firefox")
)

type (
	// NewConfig is the config that sets up webshot for use
	NewConfig struct {
		// DriverPath is the webdriver for the browser
		DriverPath string //
		// Address for the address e.g 127.0.0.1
		Address string
		// Port for the port to run the selenium server on e.g 8080
		Port int
		// DebugMode turn this to true if you want to get logs for the process happening
		DebugMode bool
		// BrowserName is the name of the browser to use
		BrowserName string
	}

	// Puppet is the struct for the webdriver
	Webshot struct {
		// webdriver is the selenium we driver
		Webdriver selenium.WebDriver
		// Service is the selenuim service
		Service *selenium.Service
	}
)
