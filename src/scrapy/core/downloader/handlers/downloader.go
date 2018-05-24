package handlers

type DownloadHandlers struct {
	crawler       interface{}
	schemes       map[string]string
	handlers      map[string]string
	notConfigured map[string]string
}
