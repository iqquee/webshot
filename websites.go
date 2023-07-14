package webshot

import "time"

// ScreenShot is the default screenshotter which can take the screenshot of webpages
func (p *Webshot) ScreenShot(requestURL string) ([]byte, error) {
	p.Webdriver.Get(requestURL)

	time.Sleep(4 * time.Second)
	imgByte, err := p.Webdriver.Screenshot()
	if err != nil {
		return nil, err
	}

	defer p.Service.Stop()
	defer p.Webdriver.Quit()

	return imgByte, nil
}
