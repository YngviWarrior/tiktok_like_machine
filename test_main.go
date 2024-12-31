package main

import (
	"strings"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/stretchr/testify/assert"
)

func TestRandonUserAgent(t *testing.T) {
	userAgent := RandonUserAgent()
	assert.NotEmpty(t, userAgent, "User agent should not be empty")
	assert.True(t, strings.HasPrefix(userAgent, "Mozilla/5.0"), "User agent should start with 'Mozilla/5.0'")
}

func TestRandomTimer(t *testing.T) {
	duration := RandomTimer()
	assert.GreaterOrEqual(t, duration.Seconds(), 3.0, "Duration should be at least 3 seconds")
	assert.LessOrEqual(t, duration.Seconds(), 9.0, "Duration should be at most 9 seconds")
}

func TestRandomClick(t *testing.T) {
	duration := RandomClick()
	assert.GreaterOrEqual(t, duration.Seconds(), 30.0, "Duration should be at least 30 seconds")
	assert.LessOrEqual(t, duration.Seconds(), 60.0, "Duration should be at most 60 seconds")
}

func TestRandomClickDelay(t *testing.T) {
	duration := RandomClickDelay()
	assert.GreaterOrEqual(t, duration.Milliseconds(), 500, "Duration should be at least 500 milliseconds")
	assert.LessOrEqual(t, duration.Milliseconds(), 6500, "Duration should be at most 6500 milliseconds")
}

func TestMainFunction(t *testing.T) {
	profilePath := "./chromium-profile"
	userAgent := RandonUserAgent()

	l := launcher.New().Headless(true).
		Set("user-data-dir", profilePath).
		Append("disable-blink-features", "AutomationControlled").
		Append("disable-infobars").
		Append("disable-extensions").
		Append("start-maximized").
		Append("disable-software-rasterizer").
		Append("no-sandbox").
		Append("disable-dev-shm-usage").
		Append("user-agent", userAgent)

	defer l.Cleanup()
	url := l.MustLaunch()

	browser := rod.New().
		ControlURL(url).
		Trace(true).
		SlowMotion(2 * time.Second).
		MustConnect()

	defer browser.MustClose()

	page := browser.MustPage("https://www.tiktok.com/@maayrls/live").MustWaitStable()
	page.WaitLoad()

	assert.NotNil(t, page, "Page should not be nil")
}
