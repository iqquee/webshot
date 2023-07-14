# webshot
webshot is a website screenshotting library in golang

# Get Started
In other to use this package, you need to first install `tesseract` on your machine. Download GekoDriver for you os from [https://github.com/mozilla/geckodriver/releases](https://github.com/mozilla/geckodriver/releases).

### NOTE: The broswer used in this package by default is firefox. Kindly install firefox if you don't have it on your machine already.
# Installation
This package can in installed using the go command below.
```sh
go get github.com/iqquee/webshot
```
# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```

# ScreenShot
ScreenShot is the default screenshotter which can take the screesnhot of webpages

```go
package main

import (
	"fmt"
	"os"

	"github.com/iqquee/webshot"
)

func main() {
	c := webshot.NewConfig{
		Address:     "http://localhost",
		Port:        4444, // you can change accordingly to which ever port you wish
		BrowserName: "firefox",
		DebugMode:   true, // set to true if you want to get the logs
		DriverPath:  "", // your gekodriver path goes in here
	}

	driver, err := webshot.NewWebshot(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	url := "https://google.com"

	byteImage, err := driver.DefaultScreenShot(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := "screenshot" + ".png"
	pngData, _ := os.Create(fileName)
	pngData.Write([]byte(byteImage))
}
```


# ImageProcessing
ImageProcessing does the optical character recognition(OCR)

This method processess an image and returns the text in that image in a .txt file.

```go
package main

import (
	"fmt"

	"github.com/iqquee/webshot"
)

func main() {

    filePath := ""
    fileName := ""
	if err := webshot.ImageProcessing(filePath, fileName); err != nil {
		fmt.Println("Image processing err: \n", err)
		return
	}
}
```