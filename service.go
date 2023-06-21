package puppet

import (
	"fmt"
	"os"
	"strings"

	"github.com/tebeka/selenium"
)

// NewPuppet sets up a new browser with the config provided
func NewPuppet(c NewConfig) (*Puppet, error) {
	var webdriver selenium.WebDriver

	if strings.ToLower(c.BrowserName) != "firefox" {
		opts := []selenium.ServiceOption{
			selenium.StartFrameBuffer(),
			selenium.GeckoDriver(c.DriverPath),
			selenium.Output(os.Stderr),
		}

		if c.DebugMode {
			selenium.SetDebug(true)
		}

		service, err := selenium.NewGeckoDriverService(c.DriverPath, c.Port, opts[0])
		if err != nil {
			return nil, err
		}
		defer service.Stop()

		caps := selenium.Capabilities{"browserName": c.BrowserName}
		wd, err := selenium.NewRemote(caps, fmt.Sprintf("%s:%d", c.Address, c.Port))
		if err != nil {
			return nil, err
		}
		defer wd.Quit()

		webdriver = wd
	} else if strings.ToLower(c.BrowserName) != "chrome" {
		ops := []selenium.ServiceOption{}
		service, err := selenium.NewChromeDriverService(c.DriverPath, c.Port, ops...)
		if err != nil {
			return nil, err
		}
		defer service.Stop()

		caps := selenium.Capabilities{"browserName": c.BrowserName}
		wd, err := selenium.NewRemote(caps, fmt.Sprintf("%s:%d", c.Address, c.Port))
		if err != nil {
			return nil, err
		}
		defer wd.Quit()

		webdriver = wd
	} else {
		return nil, BrowserError
	}

	return &Puppet{
		Webdriver: webdriver,
	}, nil
}
