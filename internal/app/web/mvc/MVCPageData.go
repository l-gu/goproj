package mvc

import (
	"fmt"
	"reflect"
)

type MVCPage struct {
	// Page title
	Title string
	// href URI prefix for links ( eg "foo" for "/foo/form" and "/foo/list" )
	Prefix string
	// form mode (Creation or Update mode)
	UpdateMode bool // false by default ("zero value" is false for boolean)
	// List of messages to be printed in the page
	Messages []string
}

type ErrorsProvider interface {
	getErrors() []string
}

//-----------------------------------------------------------------------------
// LIST DATA
//-----------------------------------------------------------------------------
// Data to be used in LIST page
type MVCListData struct {
	Page MVCPage
}

func (this *MVCListData) init(title string, prefix string) {

	log("MVCListData", "init()", fmt.Sprintf("title : %s", title))
	this.Page.Title = title

	log("MVCListData", "init()", fmt.Sprintf("prefix : %s", prefix))
	this.Page.Prefix = prefix
}

//-----------------------------------------------------------------------------
// FORM DATA
//-----------------------------------------------------------------------------
// Data to be used in FORM page
type MVCFormData struct {
	Page MVCPage
}

func (this *MVCFormData) addMessage(m string) {
	this.Page.Messages = append(this.Page.Messages, m)
	log("MVCFormData", "addMessage", m)
}

func (this *MVCFormData) addMessages(messages ...string) {
	log("MVCFormData", "addMessages", messages)
	for _, m := range messages {
		this.Page.Messages = append(this.Page.Messages, m)
	}
}

func (this *MVCFormData) addErrorMessages(messages []string) {
	log("MVCFormData", "addErrorMessages", messages)
	for _, m := range messages {
		this.Page.Messages = append(this.Page.Messages, m)
	}
}

func (this *MVCFormData) init(title string, prefix string, mode int, errProvider ErrorsProvider) {

	log("MVCFormData", "init()", fmt.Sprintf("title : %s", title))
	this.Page.Title = title

	log("MVCFormData", "init()", fmt.Sprintf("prefix : %s", prefix))
	this.Page.Prefix = prefix

	// init form mode : UPDATE or CREATION mode
	log("MVCFormData", "init()", fmt.Sprintf("mode : %d", mode))
	if mode == UPDATE_MODE {
		this.Page.UpdateMode = true
	} else {
		this.Page.UpdateMode = false
	}

	// init form messages (from validator)
	log("MVCFormData", "init()", fmt.Sprintf("errProvider : %T", errProvider))
	if errProvider != nil {
		// NB : An interface holding nil value is not nil !
		// => check value is a not nil pointer
		v := reflect.ValueOf(errProvider)
		if v.Kind() == reflect.Ptr { // it's a pointer
			if !v.IsNil() { // the pointer is not nil
				this.addErrorMessages(errProvider.getErrors())
			}
		}
	}
}
