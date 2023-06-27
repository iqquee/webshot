package puppet

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
)

// NewPuppet sets up a new browser with the config provided
func NewPuppet(c NewConfig) (*Puppet, error) {
	// var webdriver selenium.WebDriver

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

	caps := selenium.Capabilities{"browserName": c.BrowserName}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("%s:%d", c.Address, c.Port))
	if err != nil {
		return nil, err
	}

	wd.MaximizeWindow("")

	fmt.Printf("The service ran successfully\n")
	return &Puppet{
		Webdriver: wd,
		Service:   service,
	}, nil
}
