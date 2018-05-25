package settings

type BaseSettings struct {
	frozen bool
	attributes map[string]string
}
type Settings struct {
	*BaseSettings

}

type CrawlerSettings struct {
	*Settings
	settingsModule string
}
