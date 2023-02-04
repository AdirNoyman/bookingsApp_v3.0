package render

import (
	"bytes"
	"github.com/AdirNoyman/hotelBookings/pkg/config"
	"github.com/AdirNoyman/hotelBookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	// Assign the template cache to the 'app' variable
	app = a
}

// AddDefaultData - Get the data that is already in the repository
// Here we can add data that we want to be available a cross every page in our site
func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders html templates
// tmpl = the template name
// w = http response we are writing the template to
// tmpl = the template we want to render
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var templateCache map[string]*template.Template

	// If in production environment, use the cache
	if app.UseCache {

		// Get the template cache from the app config
		templateCache = app.TemplateCache
		// If in development environment, DO NOT use the cache
	} else {

		// Create the template cache
		templateCache, _ = CreateTemplateCache()
	}

	// 1. create a template cache
	//templateCache, err := CreateTemplateCache()
	//if err != nil {
	//	// Kill the app
	//	println("Problem creating the templates cache. Killing the App 😥")
	//	log.Fatal(err)
	//}
	// 2. get the requested template from the cache
	renderedTemplate, isRenderedTemplateFound := templateCache[tmpl]

	if !isRenderedTemplateFound {
		// Kill the app
		log.Fatal("Problem creating the templates cache. Killing the App 😥")
	}

	// create a buffer that will hold the bytes we got from the rendered template
	buf := new(bytes.Buffer)

	// Get the data from our main repository
	td = AddDefaultData(td)
	// Check if the value (bytes that represent the rendered template) we got form the cache, we can write to the http response
	err := renderedTemplate.Execute(buf, td)
	if err != nil {
		// Kill the app
		println("Problem with the templates value in the cache")
	}
	// 3. render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		// Kill the app
		println("Error writing to the http response writer: ", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// create an empty templates dictionary(map)
	myCache := map[string]*template.Template{}

	// You need to get all the files named *.page.tmpl from disk ("./templates")
	// Glob returns an array with the matching files it found
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {

		return myCache, err
	}

	// Range through all the files suffixed with '.page.tmpl
	// first place is for index, but we are not using in it here
	for _, page := range pages {

		// Get the fileName of the template from the file path (which is in 'page' variable). It returns the end point of the path (also called Base)
		fileName := filepath.Base(page)
		// create a new parsed template with its fileName (without the all path...just the fileName)
		templateSet, err := template.New(fileName).ParseFiles(page)
		if err != nil {

			return myCache, err
		}

		// Look for layouts if exist
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		// Glob returns an array with the matching files it found (in our case -> just one layout)
		if len(matches) > 0 {
			// Combine and parse the template page with the layout the page asks for
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {

				return myCache, err
			}
		}
		// Add the final result of the parsed template (+ her layout) to the templates cache
		myCache[fileName] = templateSet
	}
	return myCache, nil
}

//var templateCache = make(map[string]*template.Template)

//func RenderTemplate(w http.ResponseWriter, t string) {
//
//	var tmpl *template.Template
//	var err error
//
//	// Check to see if we already have the template in our cache. The template file name(string) is in the variable t, and this represents the key
//	_, isTemplateKeyInMap := templateCache[t]
//
//	if !isTemplateKeyInMap {
//		// Template not in cache. Need to create it (read it from disk and parse it)
//		println("Trying to create template and add it to cache 😎🤞")
//		err = createTemplateCache(t)
//		if err != nil {
//			println("Error making a new entry to the template cache 🤨")
//		}
//	} else {
//
//		// The template is already in the cache
//		println("Using cached template 🤓")
//
//	}
//
//	tmpl = templateCache[t]
//
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		println("Error writing the parsed template in to the http response 🤨")
//	}
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
//	}
//
//	// parse the template (each item in the templates array will be parsed
//	tmpl, err := template.ParseFiles(templates...)
//
//	if err != nil {
//		println("Error parsing templates 😩")
//		return err
//	}
//
//	// Add the parsed templates to the templates cache(map)
//	// Here we put the name of the template as Key (e.g. "about.page.tmpl") and the parsed result(tmpl) as Value
//	templateCache[t] = tmpl
//
//	return nil
//}
