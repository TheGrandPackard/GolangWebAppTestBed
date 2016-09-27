package template

import (
	"html/template"
	"io/ioutil"
	"log"
)

var isProduction = false
var isTemplateLoaded = false
var contentTemplate map[string]*template.Template

// LoadTemplates Before Rendering
func LoadTemplates() (err error) {
	if isTemplateLoaded && isProduction {
		return
	}

	isTemplateLoaded = true

	contentTemplate = make(map[string]*template.Template)
	log.Println("Loading Templates")

	bData, err := ioutil.ReadFile("templates/_template.tpl")
	if err != nil {
		return err
	}

	contentTemplate["_template"], err = template.New("_template").Parse(string(bData))
	if err != nil {
		return err
	}

	bData, err = ioutil.ReadFile("templates/_header.tpl")
	if err != nil {
		return err
	}

	_, err = contentTemplate["_template"].New("header").Parse(string(bData))
	if err != nil {
		return err
	}

	paths := []string{
		"view",
		"edit",
		"pages",
		"login",
		"logout",
		"signup",
	}

	for _, path := range paths {
		contentTemplate[path], err = contentTemplate["_template"].Clone()
		if err != nil {
			return err
		}

		bData, err = ioutil.ReadFile("templates/" + path + ".tpl")
		if err != nil {
			return err
		}
		_, err = contentTemplate[path].New("content").Parse(string(bData))
		if err != nil {
			return err
		}
	}

	return nil
}
