package settings

type BaseSettings struct {
	frozen     bool
	attributes map[string]string
}

func (s *BaseSettings) Get(key string) string {
	return s.attributes[key]
}

func (s *BaseSettings) Set(key string, value string) {
	s.attributes[key] = value
}

func (s *BaseSettings) SetDict(key string, value string) {
	s.attributes[key] = value
}

type Settings struct {
	*BaseSettings
}

type CrawlerSettings struct {
	*Settings
	settingsModule string
}

func NewCrawlerSettings() *CrawlerSettings {
	settings := &CrawlerSettings{}
	return settings
}
