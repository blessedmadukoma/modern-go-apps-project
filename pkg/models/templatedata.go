package models

// TemplateData holds the data passed to the templates (HTML)
type TemplateData struct {
	StringMap map[string]string
	FloatMap  map[string]float32
	IntMap    map[string]int
	Data      map[string]interface{}
	CSRFToken string
	Flash     string // message
	Warning   string // message
	Error     string // message
}
