package mvc

import (
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/l-gu/goproject/internal/app/logger"
)

type MVCControllerAdapter interface {
	//checkPrerequisites()
	getListData(r *http.Request) interface{}
	getVoidFormData(r *http.Request) interface{}
	getEntityFormData(r *http.Request) interface{}
	create(r *http.Request) interface{}
	delete(r *http.Request) bool
	update(r *http.Request) interface{}
}

type MVCController struct {
	adapter          MVCControllerAdapter
	listTemplateName string
	formTemplateName string
	name             string
}

func (this *MVCController) log(v ...interface{}) {
	if LogFlag {
		logger.Log(this.name, v...) // v... : treat input to function as variadic
	}
}

func recovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		var msg string
		switch value := r.(type) {
		case string:
			msg = value
		case error:
			msg = value.Error()
		default:
			msg = "unknown error type "
		}
		stack := debug.Stack()
		errorPage(w, "Runtime error", msg, stack)
	}
}

//------------------------------------------------------------------------
// HTTP REQUEST HANDLER (REQUEST ENTRY POINT)
//------------------------------------------------------------------------
func (this *MVCController) HttpHandler(w http.ResponseWriter, r *http.Request) {
	defer recovery(w)
	this.log("HttpHandler :", r.URL.Path)
	this.processRequest(w, r)
}

// Process request : URI parsing
func (this *MVCController) processRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	this.log("processRequest :", path)
	if r.Method == "POST" {
		r.ParseForm()
	}

	if strings.HasSuffix(path, "/list") {
		this.processList(w, r)
	} else if strings.HasSuffix(path, "/form") {
		this.processForm(w, r)
	} else if strings.HasSuffix(path, "/create") {
		this.processCreate(w, r)
	} else if strings.HasSuffix(path, "/update") {
		this.processUpdate(w, r)
	} else if strings.HasSuffix(path, "/delete") {
		this.processDelete(w, r)
	} else {
		errorPage(w, "Unexpected action (path = '"+path+"')", "", nil)
	}
}

// Process "xxx/list" request
func (this *MVCController) processList(w http.ResponseWriter, r *http.Request) {
	this.log("processList ")
	// call concrete controller to get list data
	//data := this.adapter.findAll(r)
	data := this.adapter.getListData(r)
	// forward to view ( list page )
	applyTemplate(w, this.listTemplateName, data)
}

// Process "xxx/form" request
func (this *MVCController) processForm(w http.ResponseWriter, r *http.Request) {
	var formData interface{}
	paramCount := len(r.URL.Query())
	if paramCount == 0 {
		this.log("processFormGet :  no parameter")
		// No paramater => call concrete controller to get void form data
		formData = this.adapter.getVoidFormData(r)
	} else {
		this.log("processFormGet :  ", paramCount, " parameter(s)")
		// paramater(s) => call concrete controller to get entity form data
		formData = this.adapter.getEntityFormData(r)
	}
	// forward to view ( form page )
	applyTemplate(w, this.formTemplateName, formData)
}

// Process "xxx/create" request
func (this *MVCController) processCreate(w http.ResponseWriter, r *http.Request) {
	this.log("processCreate ")
	// call concrete controller to get list data
	data := this.adapter.create(r)
	// refresh FORM page
	applyTemplate(w, this.formTemplateName, data)
}

// Process "xxx/delete" request
func (this *MVCController) processDelete(w http.ResponseWriter, r *http.Request) {
	this.log("processDelete ")
	// call concrete controller
	this.adapter.delete(r)
	// always go to list page (if not deleted => doesn't exist)
	this.processList(w, r)
}

// Process "xxx/update" request
func (this *MVCController) processUpdate(w http.ResponseWriter, r *http.Request) {
	this.log("processUpdate ")
	// call concrete controller
	data := this.adapter.update(r)
	// forward to view ( form page )
	applyTemplate(w, this.formTemplateName, data)
}
