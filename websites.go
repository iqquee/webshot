package puppet

import (
	"os"
	"time"
)

// FacebookScreenShot gets a screenshot off facebook
func (p *Puppet) FacebookScreenShot(requestURL string) error {
	p.Webdriver.Get(FacebookLoginRoute)

	username, err := p.Webdriver.FindElement("id", "email")
	if err != nil {
		return err
	}
	username.SendKeys("your_username")

	password, err := p.Webdriver.FindElement("id", "pass")
	if err != nil {
		return err
	}
	password.SendKeys("your_password")

	btn, err := p.Webdriver.FindElement("name", "login")
	if err != nil {
		return err
	}

	if err := btn.Click(); err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	p.Webdriver.Get(requestURL)

	pngBase64, err := p.Webdriver.Screenshot()
	if err != nil {
		return err
	}

	fileName := "screenshot" + time.Now().String() + ".png"
	pngData, _ := os.Create(fileName)
	pngData.Write([]byte(pngBase64))

	return nil
}

// InstagramScreenShot gets a screenshot off instagram
func (p *Puppet) InstagramScreenShot(requestURL string) error {
	p.Webdriver.Get(InstagramLoginRoute)

	username, err := p.Webdriver.FindElement("name", "username")
	if err != nil {
		return err
	}
	username.SendKeys("your_username")

	password, err := p.Webdriver.FindElement("name", "password")
	if err != nil {
		return err
	}
	password.SendKeys("your_password")

	btn, err := p.Webdriver.FindElement("xpath", "//button[contains(.,'Log In')]")
	if err != nil {
		return err
	}
	if err := btn.Click(); err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	p.Webdriver.Get(requestURL)

	pngBase64, err := p.Webdriver.Screenshot()
	if err != nil {
		return err
	}

	fileName := "screenshot" + time.Now().String() + ".png"
	pngData, _ := os.Create(fileName)
	pngData.Write([]byte(pngBase64))

	return nil
}

// TwitterScreenShot gets a screenshot off twitter
func (p *Puppet) TwitterScreenShot(requestURL string) error {
	p.Webdriver.Get(TwitterLoginRoute)

	username, err := p.Webdriver.FindElement("name", "session[username_or_email]")
	if err != nil {
		return err
	}
	username.SendKeys("your_username")

	password, err := p.Webdriver.FindElement("name", "session[password]")
	if err != nil {
		return err
	}
	password.SendKeys("your_password")

	btn, err := p.Webdriver.FindElement("name", "session[remember_me]")
	if err != nil {
		return err
	}
	if err := btn.Click(); err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	p.Webdriver.Get(requestURL)

	pngBase64, err := p.Webdriver.Screenshot()
	if err != nil {
		return err
	}

	fileName := "screenshot" + time.Now().String() + ".png"
	pngData, _ := os.Create(fileName)
	pngData.Write([]byte(pngBase64))

	return nil
}

// JijiDotNgScreenShot gets a screenshot off jiji.ng
// Jiji is a nigerian ecommerce company
func (p *Puppet) JijiDotNgScreenShot(requestURL string) error {
	p.Webdriver.Get(JijiDotNgLoginRoute)

	username, err := p.Webdriver.FindElement("name", "email")
	if err != nil {
		return err
	}
	username.SendKeys("your_username")

	password, err := p.Webdriver.FindElement("name", "password")
	if err != nil {
		return err
	}
	password.SendKeys("your_password")

	btn, err := p.Webdriver.FindElement("class name", "btn.btn-primary.btn-block")
	if err != nil {
		return err
	}
	if err := btn.Click(); err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	p.Webdriver.Get(requestURL)

	pngBase64, err := p.Webdriver.Screenshot()
	if err != nil {
		return err
	}

	fileName := "screenshot" + time.Now().String() + ".png"
	pngData, _ := os.Create(fileName)
	pngData.Write([]byte(pngBase64))

	return nil
}
