package puppet

import (
	"time"

	"github.com/tebeka/selenium"
)

// FacebookScreenShot gets a screenshot off facebook
// name can be your phone number or email address while password is your password
func (p *Puppet) FacebookScreenShot(name, password, requestURL string) ([]byte, error) {
	if err := p.Webdriver.Get(FacebookLoginRoute); err != nil {
		return nil, err
	}

	us, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "#email")
	if err != nil {
		return nil, err
	}
	us.SendKeys(name)

	ps, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "#pass")
	if err != nil {
		return nil, err
	}
	ps.SendKeys(password)

	btn, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "button[name=\"login\"]")
	if err != nil {
		return nil, err
	}

	if err := btn.Click(); err != nil {
		return nil, err
	}

	time.Sleep(5 * time.Second)

	if err := p.Webdriver.Get(requestURL); err != nil {
		return nil, err
	}

	time.Sleep(5 * time.Second)

	imgByte, err := p.Webdriver.Screenshot()
	if err != nil {
		return nil, err
	}

	defer p.Service.Stop()
	defer p.Webdriver.Quit()

	return imgByte, nil
}

// InstagramScreenShot gets a screenshot off instagram
func (p *Puppet) InstagramScreenShot(username, password, requestURL string) ([]byte, error) {
	if err := p.Webdriver.Get(InstagramLoginRoute); err != nil {
		return nil, err
	}

	time.Sleep(time.Second * 5)

	us, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "input[name=\"username\"]")
	if err != nil {
		return nil, err
	}
	us.SendKeys(username)

	ps, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "input[name=\"password\"]")
	if err != nil {
		return nil, err
	}
	ps.SendKeys(password)

	btn, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "button[type=\"submit\"]")
	if err != nil {
		return nil, err
	}
	if err := btn.Click(); err != nil {
		return nil, err
	}

	if err := p.Webdriver.Get(requestURL); err != nil {
		return nil, err
	}

	time.Sleep(5 * time.Second)
	imgByte, err := p.Webdriver.Screenshot()
	if err != nil {
		return nil, err
	}

	defer p.Service.Stop()
	defer p.Webdriver.Quit()

	return imgByte, nil
}

// TwitterScreenShot gets a screenshot off twitter
func (p *Puppet) TwitterScreenShot(name, password, requestURL string) ([]byte, error) {
	if err := p.Webdriver.Get(TwitterLoginRoute); err != nil {
		return nil, err
	}

	time.Sleep(time.Second * 5)
	loginbtn, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "a[data-testid=\"login\"]")
	if err != nil {
		return nil, err
	}
	if err := loginbtn.Click(); err != nil {
		return nil, err
	}

	time.Sleep(5 * time.Second)
	us, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "input[autocomplete=\"username\"]")
	if err != nil {
		return nil, err
	}
	us.SendKeys(name)
	next, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "div[role=\"button\"]")
	if err != nil {
		return nil, err
	}
	if err := next.Click(); err != nil {
		return nil, err
	}

	time.Sleep(10 * time.Second)
	ps, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "input[name=\"password\"]")
	if err != nil {
		return nil, err
	}
	ps.SendKeys(password)

	btn, err := p.Webdriver.FindElement(selenium.ByCSSSelector, "div[role=\"button\"]")
	if err != nil {
		return nil, err
	}
	if err := btn.Click(); err != nil {
		return nil, err
	}

	time.Sleep(5 * time.Second)

	if err := p.Webdriver.Get(requestURL); err != nil {
		return nil, err
	}

	imgByte, err := p.Webdriver.Screenshot()
	if err != nil {
		return nil, err
	}

	defer p.Service.Stop()
	defer p.Webdriver.Quit()

	return imgByte, nil
}

// DefaultScreenShot is the default screenshotter which can screesnhot of webpages that does not require any form of authentication
func (p *Puppet) DefaultScreenShot(requestURL string) ([]byte, error) {
	p.Webdriver.Get(requestURL)

	imgByte, err := p.Webdriver.Screenshot()
	if err != nil {
		return nil, err
	}

	defer p.Service.Stop()
	defer p.Webdriver.Quit()

	return imgByte, nil
}
