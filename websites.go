package puppet

import (
	"fmt"
	"os"
	"time"
)

func (p *Puppet) FacebookScreenShot() {

}

// GetInstagram is able to get a screen shot off instagram
func (p *Puppet) InstagramScreenShot(requestURL string) {
	p.Webdriver.Get(InstagramLoginRoute)

	username, err := p.Webdriver.FindElement("name", "username")
	if err != nil {
		fmt.Println("error finding element: ", err)
		return
	}
	username.SendKeys("your_username")

	password, err := p.Webdriver.FindElement("name", "password")
	if err != nil {
		fmt.Println("error finding element: ", err)
		return
	}
	password.SendKeys("your_password")

	// Click login button
	btn, err := p.Webdriver.FindElement("xpath", "//button[contains(.,'Log In')]")
	if err != nil {
		fmt.Println("error finding the login button element: ", err)
		return
	}
	if err := btn.Click(); err != nil {
		fmt.Println("error clicking the login button: ", err)
		return
	}

	// Wait for Instagram to load
	time.Sleep(5 * time.Second)

	// Navigate to post URL
	p.Webdriver.Get(requestURL)

	// Take screenshot and save it
	pngBase64, err := p.Webdriver.Screenshot()
	if err != nil {
		fmt.Println("an error occurened while taking screenshot: ", err)
		return
	}
	pngData, _ := os.Create("screenshot.png")
	pngData.Write([]byte(pngBase64))
}
func (p *Puppet) TwitterScreenShot() {}
