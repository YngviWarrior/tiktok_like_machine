package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"
)

func RandonUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:60.0) Gecko/20100101 Firefox/60.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edge/91.0.864.59",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36",
	}

	// Gerar um número aleatório para escolher um User-Agent
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(userAgents)) // Escolhe um índice aleatório

	// Seleciona o User-Agent aleatório
	return userAgents[randomIndex]
}

func RandomTimer() time.Duration {
	return time.Duration(rand.Intn(7)+3) * time.Second
}

func RandomClick() time.Duration {
	return time.Duration(rand.Intn(30)+30) * time.Second
}

func RandomClickDelay() time.Duration {
	return time.Duration(rand.Intn(6000)+500) * time.Millisecond
}

func main() {
	profilePath := "./chromium-profile"

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		return
	}

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

	time.Sleep(5 * time.Second)

	// Create a new page
	page := browser.MustPage("https://www.tiktok.com/@maayrls/live").MustWaitStable()
	page.WaitLoad()

	for {
		element, err := page.Element(`div[class*="css-1mk0i7a-DivLikeBtnIcon ebnaa9i3"]`)
		if err != nil {
			log.Panic(err)
		}
		element.MustClick()
	}
}
