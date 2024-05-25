package internal

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

func ScheduleScrapping(config Config) {
	c := buildScrapPipeline(config)

	for {
		log.Info("starting scrap run")

		for _, url := range config.urls {
			err := c.Visit(url.String())
			if err != nil {
				log.Errorf("error while visiting '%s','%s", url.String(), err.Error())
			}
		}

		log.Info("scrap run ended")
		time.Sleep(time.Hour * 1)
	}
}

func buildScrapPipeline(config Config) *colly.Collector {

	pattern, _ := regexp.Compile("https://play\\.google\\.com/store/apps/details\\?id=(([A-Za-z][A-Za-z\\\\d_]*\\.)*[A-Za-z][A-Za-z\\\\d_]*)")

	c := colly.NewCollector(colly.AllowURLRevisit())
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if pattern.MatchString(link) {
			packageName := pattern.FindStringSubmatch(link)[1]
			log.Info(fmt.Sprintf("Found link '%s' with package name '%s'", link, packageName))
			submitNewAppPackage(config, packageName)
		}
	})

	c.OnRequest(func(request *colly.Request) {
		log.Info(fmt.Sprintf("Visiting '%s'", request.URL.String()))
	})

	return c
}
