package mvc

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/l-gu/goproject/internal/app/config"
	"github.com/l-gu/goproject/internal/app/logger"
)

func log(v ...interface{}) {
	if LogFlag {
		logger.Log("MVC", v...) // v... : treat input to function as variadic
	}
}

var mvcTemplate *template.Template

func InitTemplates() {
	log("Templates initialization")
	pattern := config.GetWebDir() + "/templates/*.gohtml"
	var err error
	mvcTemplate, err = template.ParseGlob(pattern)
	if err != nil {
		log("ERROR in templates parsing")
		panic(err)
	}
	// print available templates
	templates := mvcTemplate.Templates()
	log(fmt.Sprintf("%d template(s) parsed.", len(templates)))
	for i, t := range templates {
		log(fmt.Sprintf(" . %d : %s", i, t.Name()))
	}
}

//------------------------------------------------------------------------------
// CONVERSION TO STRING
//------------------------------------------------------------------------------
func intToString(v int) string {
	return strconv.Itoa(v)
}
func int32ToString(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}
func int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}
func float32ToString(v float32) string {
	return fmt.Sprintf("%.2f", v)
}
func float64ToString(v float64) string {
	return fmt.Sprintf("%.2f", v)
}
func boolToString(v bool) string {
	return strconv.FormatBool(v) // returns "true" or "false"
}

//------------------------------------------------------------------------------
// CONVERSION FROM STRING
//------------------------------------------------------------------------------
func stringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
func stringToInt32(s string) (int32, error) {
	// bitSize : 0 for int, 8 for int8, 16 for int16, 32 for int32, and 64 for int64.
	i, e := strconv.ParseInt(s, 10, 32) // returns int64, error
	if e == nil {
		return int32(i), nil
	} else {
		return 0, e
	}
}
func stringToInt64(s string) (int64, error) {
	// bitSize : 0 for int, 8 for int8, 16 for int16, 32 for int32, and 64 for int64.
	return strconv.ParseInt(s, 10, 64)
}

func stringToFloat32(s string) (float32, error) {
	// bitSize: 32 for float32, or 64 for float64.
	v, e := strconv.ParseFloat(s, 32) // returns float64, error
	if e == nil {
		return float32(v), nil
	} else {
		return 0, e
	}
}
func stringToFloat64(s string) (float64, error) {
	// bitSize: 32 for float32, or 64 for float64.
	return strconv.ParseFloat(s, 64) // returns float64, error
}

func stringToBool(s string) (bool, error) {
	return strconv.ParseBool(s) // returns bool, error
}

//------------------------------------------------------------------------------
func errorPage(w http.ResponseWriter, msg1 string, msg2 string, stack []byte) {
	fmt.Fprint(w, "<!DOCTYPE html>")
	fmt.Fprint(w, "<html>")
	fmt.Fprint(w, "<body>")
	fmt.Fprint(w, "<h1>ERROR</h1>\n")
	// Message 1
	fmt.Fprintf(w, "<h2>%s</h2>\n", msg1)
	// Message 2 (if any)
	if msg2 != "" {
		fmt.Fprintf(w, "<h2>%s</h2>\n", msg2)
	}
	// Stack (if any)
	if stack != nil {
		fmt.Fprint(w, "<pre> <code>")
		fmt.Fprint(w, string(stack))
		fmt.Fprint(w, "</code> </pre>")
	}
	fmt.Fprint(w, "</body>")
	fmt.Fprint(w, "</html>")
}

func buildTemplatePath(templateName string) string {
	return config.GetWebDir() + "/templates/" + templateName + ".gohtml"
}

func applyTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	log("applyTemplate(" + templateName + ")")
	// TODO : choose mode according config
	applyTemplateDevMode(w, templateName, data)
}

func applyTemplateProdMode(w http.ResponseWriter, templateName string, data interface{}) {
	log("applyTemplateProdMode(" + templateName + ")")
	fullTemplateName := templateName + ".gohtml"
	template := mvcTemplate.Lookup(fullTemplateName)
	if template != nil {
		// Merge the template with the given data to produce the responce
		err := template.Execute(w, data)
		if err != nil {
			errorPage(w, "Cannot execute template '"+fullTemplateName+"'", err.Error(), nil)
		}
	} else {
		errorPage(w, "Template '"+fullTemplateName+"' not found", "", nil)
	}
}
func applyTemplateDevMode(w http.ResponseWriter, templateName string, data interface{}) {
	log("applyTemplateDevMode(" + templateName + ")")

	// Build path
	templatePath := buildTemplatePath(templateName)
	nestedTemplatesPath := buildTemplatePath("nested-templates")

	// Parse the template file
	template, err := template.ParseFiles(templatePath, nestedTemplatesPath)
	if err != nil {
		errorPage(w, "Cannot parse template file '"+templatePath+"'", err.Error(), debug.Stack())
		return
	}

	// Merge the template with the given data to produce the responce
	err = template.Execute(w, data)
	if err != nil {
		errorPage(w, "Cannot execute template file '"+templatePath+"'", err.Error(), debug.Stack())
	}
}

//------------------------------------------------------------------------------
// URL PARAMETERS
//------------------------------------------------------------------------------
func urlParam(request *http.Request, name string) string {
	// r.URL.Query() returns a 'Values' type
	// which is simply a map[string][]string of the QueryString parameters.
	queryValues := request.URL.Query()

	// Query()["key"] will return an array of items,
	// we only want a single item => use the first one
	values, ok := queryValues[name]
	if ok && len(values) > 0 {
		return values[0] // found
	} else {
		return "" // not found
	}
}

func urlParamAsString(request *http.Request, name string, defaultValue string) string {
	return paramAsString(urlParam(request, name), defaultValue)
}
func urlParamAsInt(request *http.Request, name string, defaultValue int) int {
	return paramAsInt(urlParam(request, name), defaultValue)
}

//------------------------------------------------------------------------------
// FORM PARAMETERS
//------------------------------------------------------------------------------

// Returns the request form parameter for the given name ( "" if not found )
func formParam(request *http.Request, name string) string {
	return request.Form.Get(name)
}

func formParamAsString(request *http.Request, name string, defaultValue string) string {
	return paramAsString(formParam(request, name), defaultValue)
}

// Returns the request parameter for the given name
// NB : 'ParseForm()' must be called before using this function
// . request : http request,
// . name : the parameter name
// . defaultValue : the default value to be returned if parameter not found or void
func formParamAsInt(request *http.Request, name string, defaultValue int) int {
	return paramAsInt(formParam(request, name), defaultValue)
}

func paramAsString(v string, defaultValue string) string {
	if v != "" {
		// param found
		return v
	} else {
		// not found
		return defaultValue
	}
}
func paramAsInt(v string, defaultValue int) int {
	if v != "" {
		// param found
		i, err := strconv.Atoi(v)
		if err != nil {
			// invalid int
			return defaultValue
		} else {
			// ok, valid
			return i
		}
	} else {
		// not found
		return defaultValue
	}
}

//------------------------------------------------------------------------------
// PARAMETERS VALIDATOR
//------------------------------------------------------------------------------
/*
Parameter :

name : KEY (string)
type

// 1
value ( interface{} ???)
found bool

// 2
required
min
max
minlength
maxlength
notblank
notempty ? = not found

function
*/
