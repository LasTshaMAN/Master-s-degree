package main

import (
	"time"

	"github.com/fedesog/webdriver"
)

var session *webdriver.Session

func initSession(chromeDriver *webdriver.ChromeDriver) (err error) {
	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err = chromeDriver.NewSession(desired, required)
	if err != nil {
		return
	}

	var wh webdriver.WindowHandle
	wh, err = session.WindowHandle()
	if err != nil {
		return
	}
	wh.MaximizeWindow()

	return
}

func setSessionURL(url string) error {
	err := session.Url(url)
	if err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond)
	return nil
}
