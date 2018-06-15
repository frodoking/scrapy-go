package scrapy

import "scrapy/settings"

func InsideProject() bool {
	return true
}

func GetProjectSettings() *settings.CrawlerSettings {
	return nil
}
