package mvc

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/l-gu/goproject/internal/app/logger"
)

type MVCValidator struct {
	name   string
	errors []string
}

func NewMVCValidator() MVCValidator {
	v := new(MVCValidator)
	v.errors = make([]string, 0, 5)
	return *v
}

func (this *MVCValidator) log(v ...interface{}) {
	if LogFlag {
		logger.Log(this.name, v...) // v... : treat input to function as variadic
	}
}
func (this *MVCValidator) addError(err string) {
	this.log("addError :", err)
	this.errors = append(this.errors, err)
	this.log("addError : len = ", len(this.errors))
}
func (this *MVCValidator) hasError() bool {
	this.log("hasError : len = ", len(this.errors))
	return len(this.errors) > 0
}
func (this *MVCValidator) getErrors() []string {
	this.log("getErrors : len = ", len(this.errors))
	return this.errors
}

// GET GENERIC PARAMETER
func (this *MVCValidator) getParam(r *http.Request, name string) (string, bool) {
	switch r.Method {
	case "GET":
		return this.getUrlParam(r, name)
	case "POST":
		return this.getFormParam(r, name)
	default:
		return "", false // not found
	}
}

func (this *MVCValidator) getUrlParam(request *http.Request, name string) (string, bool) {
	// r.URL.Query() returns a 'Values' type
	// which is simply a map[string][]string of the QueryString parameters.
	queryValues := request.URL.Query()

	// Query()["key"] will return an array of items,
	// we only want a single item => use the first one
	values, found := queryValues[name]
	if found {
		if len(values) > 0 {
			return values[0], true // found
		}
	}
	return "", false // not found
}

func (this *MVCValidator) getFormParam(request *http.Request, name string) (string, bool) {

	//	values := request.Form.Values
	//	request.Form.Get(name)
	values, found := request.Form[name]
	if found {
		if len(values) > 0 {
			return values[0], true // found
		}
	}
	return "", false // not found
}

//------------------------------------------------------------------------------
func (this *MVCValidator) paramNotFound(paramName string, required bool) {
	if required {
		m := fmt.Sprintf("parameter '%s' is required", paramName)
		this.addError(m)
	}
}
func (this *MVCValidator) invalidParameter(paramName string, paramVal string, paramType string) {
	m := fmt.Sprintf("parameter '%s' is invalid (value '%s' for type '%s')", paramName, paramVal, paramType)
	this.addError(m)
}

//------------------------------------------------------------------------------
// GET TYPED PARAMETERS VALUES
//------------------------------------------------------------------------------
func (this *MVCValidator) getStringParam(r *http.Request, paramName string, required bool) (string, bool) {
	value, found := this.getParam(r, paramName)
	if found {
		return value, true
	} else {
		this.paramNotFound(paramName, required)
		return "", false
	}
}

func (this *MVCValidator) getIntParam(r *http.Request, paramName string, required bool) (int, bool) {
	paramVal, found := this.getParam(r, paramName)
	if found {
		i, e := stringToInt(paramVal)
		if e == nil {
			return i, true // value OK and found
		} else {
			//m := fmt.Sprintf("parameter '%s' = '%s' : invalid integer", paramName, paramVal)
			//this.addError(m)
			this.invalidParameter(paramName, paramVal, "int")
		}
	} else {
		this.paramNotFound(paramName, required)
	}
	return 0, false
}

func (this *MVCValidator) getInt32Param(r *http.Request, paramName string, required bool) (int32, bool) {
	paramVal, found := this.getParam(r, paramName)
	if found {
		i, e := stringToInt32(paramVal)
		if e == nil {
			return i, true // value OK and found
		} else {
			this.invalidParameter(paramName, paramVal, "int32")
		}
	} else {
		this.paramNotFound(paramName, required)
	}
	return 0, false
}

func (this *MVCValidator) getInt64Param(r *http.Request, paramName string, required bool) (int64, bool) {
	paramVal, found := this.getParam(r, paramName)
	if found {
		i, e := stringToInt64(paramVal)
		if e == nil {
			return i, true // value OK and found
		} else {
			this.invalidParameter(paramName, paramVal, "int64")
		}
	} else {
		this.paramNotFound(paramName, required)
	}
	return 0, false
}

func (this *MVCValidator) getFloat32Param(r *http.Request, paramName string, required bool) (float32, bool) {
	paramVal, found := this.getParam(r, paramName)
	if found {
		v, e := stringToFloat32(paramVal)
		if e == nil {
			return v, true // value OK and found
		} else {
			this.invalidParameter(paramName, paramVal, "float32")
		}
	} else {
		this.paramNotFound(paramName, required)
	}
	return 0, false
}

func (this *MVCValidator) getFloat64Param(r *http.Request, paramName string, required bool) (float64, bool) {
	paramVal, found := this.getParam(r, paramName)
	if found {
		v, e := stringToFloat64(paramVal)
		if e == nil {
			return v, true // value OK and found
		} else {
			this.invalidParameter(paramName, paramVal, "float64")
		}
	} else {
		this.paramNotFound(paramName, required)
	}
	return 0, false
}

func (this *MVCValidator) getBoolParam(r *http.Request, paramName string, required bool) (bool, bool) {
	paramVal, found := this.getParam(r, paramName)
	if found {
		v, e := stringToBool(paramVal)
		if e == nil {
			return v, true // value OK and found
		} else {
			this.invalidParameter(paramName, paramVal, "bool")
		}
	} else {
		this.paramNotFound(paramName, required)
	}
	return false, false
}

//------------------------------------------------------------------------------
// VALUE CHECKING
//------------------------------------------------------------------------------
func (this *MVCValidator) checkNotEmpty(paramName string, v string) {
	if v == "" {
		this.addError(paramName + " : cannot be empty")
	}
}

func (this *MVCValidator) checkNotBlank(paramName string, v string) {
	// remove all leading and trailing white space as defined by Unicode ( ' ', '\t' )
	if strings.TrimSpace(v) == "" {
		this.addError(paramName + " : cannot be blank")
	}
}

func (this *MVCValidator) checkMaxLength(paramName string, v string, max int) {
	len := len(v)
	if len > max {
		this.addError(fmt.Sprintf("%s : length %d > %d (max)", paramName, len, max))
	}
}

func (this *MVCValidator) checkMinLength(paramName string, v string, min int) {
	len := len(v)
	if len < min {
		this.addError(fmt.Sprintf("%s : length %d < %d (min)", paramName, len, min))
	}
}
