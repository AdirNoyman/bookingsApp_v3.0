package models

// TemplateData used for containing data that might be passed by the handlers to the templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	// When you don't know the data type in advance you define the variable type as 'interface'
	Data         map[string]interface{}
	CSRFToken    string
	FlashMessage string
	Warning      string
	ErrorMessage string
}
