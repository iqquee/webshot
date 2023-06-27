package puppet

import (
	"errors"

	"github.com/tebeka/selenium"
)

var (
	// FacebookLoginRoute is the facebook login route
	FacebookLoginRoute = "https://web.facebook.com/login"
	// InstagramLoginRoute is the instagram login route
	InstagramLoginRoute = "https://instagram.com/"
	// TwitterLoginRoute is the twitter login route
	TwitterLoginRoute = "https://twitter.com/"
	// JijiDotNgLoginRoute is the jiji.ng login route
	JijiDotNgLoginRoute = "https://www.jiji.ng/login"
	// ChromeBrowser for the chrome browser
	ChromeBrowser = "chrome"
	// FirefoxBrowser for the firefox browser
	FirefoxBrowser = "firefox"
	// BrowserError is the error returned when the browser type entered by the user is neither chrome or firefox
	ErrBrowser = errors.New("the supported browsers are chrome and firefox")
)

type (
	// NewConfig is the config that sets up puppet for use
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
	Puppet struct {
		// webdriver is the selenium we driver
		Webdriver selenium.WebDriver
		// Service is the selenuim service
		Service *selenium.Service
	}
)
