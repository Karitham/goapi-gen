// Package parameters provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version (devel) DO NOT EDIT.
package parameters

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/discord-gophers/goapi-gen/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// ComplexObject defines model for ComplexObject.
type ComplexObject struct {
	ID      int    `json:"Id"`
	IsAdmin bool   `json:"IsAdmin"`
	Object  Object `json:"Object"`
}

// Object defines model for Object.
type Object struct {
	FirstName string `json:"firstName"`
	Role      string `json:"role"`
}

// GetCookieParams defines parameters for GetCookie.
type GetCookieParams struct {
	// primitive
	P *int32 `json:"p,omitempty"`

	// primitive
	Ep *int32 `json:"ep,omitempty"`

	// exploded array
	Ea []int32 `json:"ea,omitempty"`

	// array
	A []int32 `json:"a,omitempty"`

	// exploded object
	Eo *Object `json:"eo,omitempty"`

	// object
	O *Object `json:"o,omitempty"`

	// complex object
	Co *ComplexObject `json:"co,omitempty"`

	// name starting with number
	N1s *string `json:"1s,omitempty"`
}

// GetHeaderParams defines parameters for GetHeader.
type GetHeaderParams struct {
	// primitive
	XPrimitive *int32 `json:"X-Primitive,omitempty"`

	// primitive
	XPrimitiveExploded *int32 `json:"X-Primitive-Exploded,omitempty"`

	// exploded array
	XArrayExploded []int32 `json:"X-Array-Exploded,omitempty"`

	// array
	XArray []int32 `json:"X-Array,omitempty"`

	// exploded object
	XObjectExploded *Object `json:"X-Object-Exploded,omitempty"`

	// object
	XObject *Object `json:"X-Object,omitempty"`

	// complex object
	XComplexObject *ComplexObject `json:"X-Complex-Object,omitempty"`

	// name starting with number
	N1StartingWithNumber *string `json:"1-Starting-With-Number,omitempty"`
}

// GetDeepObjectParams defines parameters for GetDeepObject.
type GetDeepObjectParams struct {
	// deep object
	DeepObj ComplexObject `json:"deepObj"`
}

// GetQueryFormParams defines parameters for GetQueryForm.
type GetQueryFormParams struct {
	// exploded array
	Ea []int32 `json:"ea,omitempty"`

	// array
	A []int32 `json:"a,omitempty"`

	// exploded object
	Eo *Object `json:"eo,omitempty"`

	// object
	O *Object `json:"o,omitempty"`

	// exploded primitive
	Ep *int32 `json:"ep,omitempty"`

	// primitive
	P *int32 `json:"p,omitempty"`

	// primitive string
	Ps *string `json:"ps,omitempty"`

	// complex object
	Co *ComplexObject `json:"co,omitempty"`

	// name starting with number
	N1s *string `json:"1s,omitempty"`
}

// Response is a common response struct for all the API calls.
// A Response object may be instantiated via functions for specific operation responses.
// It may also be instantiated directly, for the purpose of responding with a single status code.
type Response struct {
	body        interface{}
	Code        int
	contentType string
}

// Render implements the render.Renderer interface. It sets the Content-Type header
// and status code based on the response definition.
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", resp.contentType)
	render.Status(r, resp.Code)
	return nil
}

// Status is a builder method to override the default status code for a response.
func (resp *Response) Status(code int) *Response {
	resp.Code = code
	return resp
}

// ContentType is a builder method to override the default content type for a response.
func (resp *Response) ContentType(contentType string) *Response {
	resp.contentType = contentType
	return resp
}

// MarshalJSON implements the json.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(resp.body)
}

// MarshalXML implements the xml.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(resp.body)
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /contentObject/{param})
	GetContentObject(w http.ResponseWriter, r *http.Request, param ComplexObject) *Response

	// (GET /cookie)
	GetCookie(w http.ResponseWriter, r *http.Request, params GetCookieParams) *Response

	// (GET /header)
	GetHeader(w http.ResponseWriter, r *http.Request, params GetHeaderParams) *Response

	// (GET /labelExplodeArray/{.param*})
	GetLabelExplodeArray(w http.ResponseWriter, r *http.Request, param []int32) *Response

	// (GET /labelExplodeObject/{.param*})
	GetLabelExplodeObject(w http.ResponseWriter, r *http.Request, param Object) *Response

	// (GET /labelNoExplodeArray/{.param})
	GetLabelNoExplodeArray(w http.ResponseWriter, r *http.Request, param []int32) *Response

	// (GET /labelNoExplodeObject/{.param})
	GetLabelNoExplodeObject(w http.ResponseWriter, r *http.Request, param Object) *Response

	// (GET /matrixExplodeArray/{.id*})
	GetMatrixExplodeArray(w http.ResponseWriter, r *http.Request, id []int32) *Response

	// (GET /matrixExplodeObject/{.id*})
	GetMatrixExplodeObject(w http.ResponseWriter, r *http.Request, id Object) *Response

	// (GET /matrixNoExplodeArray/{.id})
	GetMatrixNoExplodeArray(w http.ResponseWriter, r *http.Request, id []int32) *Response

	// (GET /matrixNoExplodeObject/{.id})
	GetMatrixNoExplodeObject(w http.ResponseWriter, r *http.Request, id Object) *Response

	// (GET /passThrough/{param})
	GetPassThrough(w http.ResponseWriter, r *http.Request, param string) *Response

	// (GET /queryDeepObject)
	GetDeepObject(w http.ResponseWriter, r *http.Request, params GetDeepObjectParams) *Response

	// (GET /queryForm)
	GetQueryForm(w http.ResponseWriter, r *http.Request, params GetQueryFormParams) *Response

	// (GET /simpleExplodeArray/{param*})
	GetSimpleExplodeArray(w http.ResponseWriter, r *http.Request, param []int32) *Response

	// (GET /simpleExplodeObject/{param*})
	GetSimpleExplodeObject(w http.ResponseWriter, r *http.Request, param Object) *Response

	// (GET /simpleNoExplodeArray/{param})
	GetSimpleNoExplodeArray(w http.ResponseWriter, r *http.Request, param []int32) *Response

	// (GET /simpleNoExplodeObject/{param})
	GetSimpleNoExplodeObject(w http.ResponseWriter, r *http.Request, param Object) *Response

	// (GET /simplePrimitive/{param})
	GetSimplePrimitive(w http.ResponseWriter, r *http.Request, param int32) *Response

	// (GET /startingWithNumber/{1param})
	GetStartingWithNumber(w http.ResponseWriter, r *http.Request, n1param string) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// GetContentObject operation middleware
func (siw *ServerInterfaceWrapper) GetContentObject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param ComplexObject

	if err := json.Unmarshal([]byte(chi.URLParam(r, "param")), &param); err != nil {
		siw.ErrorHandlerFunc(w, r, &UnmarshalingParamError{err, "param"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetContentObject(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetCookie operation middleware
func (siw *ServerInterfaceWrapper) GetCookie(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCookieParams

	if cookie, err := r.Cookie("p"); err == nil {
		var value int32
		if err := runtime.BindStyledParameter("simple", false, "p", cookie.Value, &value); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "p"})
			return
		}
		params.P = &value

	}

	if cookie, err := r.Cookie("ep"); err == nil {
		var value int32
		if err := runtime.BindStyledParameter("simple", true, "ep", cookie.Value, &value); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "ep"})
			return
		}
		params.Ep = &value

	}

	if cookie, err := r.Cookie("ea"); err == nil {
		var value []int32
		if err := runtime.BindStyledParameter("simple", true, "ea", cookie.Value, &value); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "ea"})
			return
		}
		params.Ea = value

	}

	if cookie, err := r.Cookie("a"); err == nil {
		var value []int32
		if err := runtime.BindStyledParameter("simple", false, "a", cookie.Value, &value); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "a"})
			return
		}
		params.A = value

	}

	if cookie, err := r.Cookie("eo"); err == nil {
		var value Object
		if err := runtime.BindStyledParameter("simple", true, "eo", cookie.Value, &value); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "eo"})
			return
		}
		params.Eo = &value

	}

	if cookie, err := r.Cookie("o"); err == nil {
		var value Object
		if err := runtime.BindStyledParameter("simple", false, "o", cookie.Value, &value); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "o"})
			return
		}
		params.O = &value

	}

	if cookie, err := r.Cookie("co"); err == nil {
		var value ComplexObject
		var decoded string
		decoded, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &UnescapedCookieParamError{err, "co"})
			return
		}

		err = json.Unmarshal([]byte(decoded), &value)
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &UnmarshalingParamError{err, "co"})
			return
		}

		params.Co = &value

	}

	if cookie, err := r.Cookie("1s"); err == nil {
		var value string
		if err := runtime.BindStyledParameter("simple", true, "1s", cookie.Value, &value); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "1s"})
			return
		}
		params.N1s = &value

	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetCookie(w, r, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetHeader operation middleware
func (siw *ServerInterfaceWrapper) GetHeader(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params GetHeaderParams

	headers := r.Header

	// ------------- Optional header parameter "X-Primitive" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Primitive")]; found {
		var XPrimitive int32
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{n, "X-Primitive"})
			return
		}

		if err := runtime.BindStyledParameterWithLocation("simple", false, "X-Primitive", runtime.ParamLocationHeader, valueList[0], &XPrimitive); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "X-Primitive"})
			return
		}

		params.XPrimitive = &XPrimitive

	}

	// ------------- Optional header parameter "X-Primitive-Exploded" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Primitive-Exploded")]; found {
		var XPrimitiveExploded int32
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{n, "X-Primitive-Exploded"})
			return
		}

		if err := runtime.BindStyledParameterWithLocation("simple", true, "X-Primitive-Exploded", runtime.ParamLocationHeader, valueList[0], &XPrimitiveExploded); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "X-Primitive-Exploded"})
			return
		}

		params.XPrimitiveExploded = &XPrimitiveExploded

	}

	// ------------- Optional header parameter "X-Array-Exploded" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Array-Exploded")]; found {
		var XArrayExploded []int32
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{n, "X-Array-Exploded"})
			return
		}

		if err := runtime.BindStyledParameterWithLocation("simple", true, "X-Array-Exploded", runtime.ParamLocationHeader, valueList[0], &XArrayExploded); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "X-Array-Exploded"})
			return
		}

		params.XArrayExploded = XArrayExploded

	}

	// ------------- Optional header parameter "X-Array" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Array")]; found {
		var XArray []int32
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{n, "X-Array"})
			return
		}

		if err := runtime.BindStyledParameterWithLocation("simple", false, "X-Array", runtime.ParamLocationHeader, valueList[0], &XArray); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "X-Array"})
			return
		}

		params.XArray = XArray

	}

	// ------------- Optional header parameter "X-Object-Exploded" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Object-Exploded")]; found {
		var XObjectExploded Object
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{n, "X-Object-Exploded"})
			return
		}

		if err := runtime.BindStyledParameterWithLocation("simple", true, "X-Object-Exploded", runtime.ParamLocationHeader, valueList[0], &XObjectExploded); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "X-Object-Exploded"})
			return
		}

		params.XObjectExploded = &XObjectExploded

	}

	// ------------- Optional header parameter "X-Object" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Object")]; found {
		var XObject Object
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{n, "X-Object"})
			return
		}

		if err := runtime.BindStyledParameterWithLocation("simple", false, "X-Object", runtime.ParamLocationHeader, valueList[0], &XObject); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "X-Object"})
			return
		}

		params.XObject = &XObject

	}

	// ------------- Optional header parameter "X-Complex-Object" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Complex-Object")]; found {
		var XComplexObject ComplexObject
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{n, "X-Complex-Object"})
			return
		}

		if err := json.Unmarshal([]byte(valueList[0]), &XComplexObject); err != nil {
			siw.ErrorHandlerFunc(w, r, &UnmarshalingParamError{err, "X-Complex-Object"})
			return
		}

		params.XComplexObject = &XComplexObject

	}

	// ------------- Optional header parameter "1-Starting-With-Number" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("1-Starting-With-Number")]; found {
		var N1StartingWithNumber string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{n, "1-Starting-With-Number"})
			return
		}

		if err := runtime.BindStyledParameterWithLocation("simple", false, "1-Starting-With-Number", runtime.ParamLocationHeader, valueList[0], &N1StartingWithNumber); err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "1-Starting-With-Number"})
			return
		}

		params.N1StartingWithNumber = &N1StartingWithNumber

	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetHeader(w, r, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetLabelExplodeArray operation middleware
func (siw *ServerInterfaceWrapper) GetLabelExplodeArray(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param []int32

	if err := runtime.BindStyledParameter("label", true, "param", chi.URLParam(r, "param"), &param); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "param"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetLabelExplodeArray(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetLabelExplodeObject operation middleware
func (siw *ServerInterfaceWrapper) GetLabelExplodeObject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param Object

	if err := runtime.BindStyledParameter("label", true, "param", chi.URLParam(r, "param"), &param); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "param"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetLabelExplodeObject(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetLabelNoExplodeArray operation middleware
func (siw *ServerInterfaceWrapper) GetLabelNoExplodeArray(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param []int32

	if err := runtime.BindStyledParameter("label", false, "param", chi.URLParam(r, "param"), &param); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "param"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetLabelNoExplodeArray(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetLabelNoExplodeObject operation middleware
func (siw *ServerInterfaceWrapper) GetLabelNoExplodeObject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param Object

	if err := runtime.BindStyledParameter("label", false, "param", chi.URLParam(r, "param"), &param); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "param"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetLabelNoExplodeObject(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetMatrixExplodeArray operation middleware
func (siw *ServerInterfaceWrapper) GetMatrixExplodeArray(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id []int32

	if err := runtime.BindStyledParameter("matrix", true, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetMatrixExplodeArray(w, r, id)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetMatrixExplodeObject operation middleware
func (siw *ServerInterfaceWrapper) GetMatrixExplodeObject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id Object

	if err := runtime.BindStyledParameter("matrix", true, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetMatrixExplodeObject(w, r, id)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetMatrixNoExplodeArray operation middleware
func (siw *ServerInterfaceWrapper) GetMatrixNoExplodeArray(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id []int32

	if err := runtime.BindStyledParameter("matrix", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetMatrixNoExplodeArray(w, r, id)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetMatrixNoExplodeObject operation middleware
func (siw *ServerInterfaceWrapper) GetMatrixNoExplodeObject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id Object

	if err := runtime.BindStyledParameter("matrix", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetMatrixNoExplodeObject(w, r, id)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetPassThrough operation middleware
func (siw *ServerInterfaceWrapper) GetPassThrough(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param string

	param = chi.URLParam(r, "param")

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetPassThrough(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetDeepObject operation middleware
func (siw *ServerInterfaceWrapper) GetDeepObject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDeepObjectParams

	// ------------- Required query parameter "deepObj" -------------

	if err := runtime.BindQueryParameter("deepObject", true, true, "deepObj", r.URL.Query(), &params.DeepObj); err != nil {
		err = fmt.Errorf("invalid format for parameter deepObj: %w", err)
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "deepObj"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetDeepObject(w, r, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetQueryForm operation middleware
func (siw *ServerInterfaceWrapper) GetQueryForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params GetQueryFormParams

	// ------------- Optional query parameter "ea" -------------

	if err := runtime.BindQueryParameter("form", true, false, "ea", r.URL.Query(), &params.Ea); err != nil {
		err = fmt.Errorf("invalid format for parameter ea: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "ea"})
		return
	}

	// ------------- Optional query parameter "a" -------------

	if err := runtime.BindQueryParameter("form", false, false, "a", r.URL.Query(), &params.A); err != nil {
		err = fmt.Errorf("invalid format for parameter a: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "a"})
		return
	}

	// ------------- Optional query parameter "eo" -------------

	if err := runtime.BindQueryParameter("form", true, false, "eo", r.URL.Query(), &params.Eo); err != nil {
		err = fmt.Errorf("invalid format for parameter eo: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "eo"})
		return
	}

	// ------------- Optional query parameter "o" -------------

	if err := runtime.BindQueryParameter("form", false, false, "o", r.URL.Query(), &params.O); err != nil {
		err = fmt.Errorf("invalid format for parameter o: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "o"})
		return
	}

	// ------------- Optional query parameter "ep" -------------

	if err := runtime.BindQueryParameter("form", true, false, "ep", r.URL.Query(), &params.Ep); err != nil {
		err = fmt.Errorf("invalid format for parameter ep: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "ep"})
		return
	}

	// ------------- Optional query parameter "p" -------------

	if err := runtime.BindQueryParameter("form", false, false, "p", r.URL.Query(), &params.P); err != nil {
		err = fmt.Errorf("invalid format for parameter p: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "p"})
		return
	}

	// ------------- Optional query parameter "ps" -------------

	if err := runtime.BindQueryParameter("form", true, false, "ps", r.URL.Query(), &params.Ps); err != nil {
		err = fmt.Errorf("invalid format for parameter ps: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "ps"})
		return
	}

	// ------------- Optional query parameter "co" -------------

	if paramValue := r.URL.Query().Get("co"); paramValue != "" {

		var value ComplexObject
		if err := json.Unmarshal([]byte(paramValue), &value); err != nil {
			siw.ErrorHandlerFunc(w, r, &UnmarshalingParamError{err, "co"})
			return
		}
		params.Co = &value

	}

	// ------------- Optional query parameter "1s" -------------

	if err := runtime.BindQueryParameter("form", true, false, "1s", r.URL.Query(), &params.N1s); err != nil {
		err = fmt.Errorf("invalid format for parameter 1s: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "1s"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetQueryForm(w, r, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetSimpleExplodeArray operation middleware
func (siw *ServerInterfaceWrapper) GetSimpleExplodeArray(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param []int32

	if err := runtime.BindStyledParameter("simple", true, "param", chi.URLParam(r, "param"), &param); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "param"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetSimpleExplodeArray(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetSimpleExplodeObject operation middleware
func (siw *ServerInterfaceWrapper) GetSimpleExplodeObject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param Object

	if err := runtime.BindStyledParameter("simple", true, "param", chi.URLParam(r, "param"), &param); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "param"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetSimpleExplodeObject(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetSimpleNoExplodeArray operation middleware
func (siw *ServerInterfaceWrapper) GetSimpleNoExplodeArray(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param []int32

	if err := runtime.BindStyledParameter("simple", false, "param", chi.URLParam(r, "param"), &param); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "param"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetSimpleNoExplodeArray(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetSimpleNoExplodeObject operation middleware
func (siw *ServerInterfaceWrapper) GetSimpleNoExplodeObject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param Object

	if err := runtime.BindStyledParameter("simple", false, "param", chi.URLParam(r, "param"), &param); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "param"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetSimpleNoExplodeObject(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetSimplePrimitive operation middleware
func (siw *ServerInterfaceWrapper) GetSimplePrimitive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "param" -------------
	var param int32

	if err := runtime.BindStyledParameter("simple", false, "param", chi.URLParam(r, "param"), &param); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "param"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetSimplePrimitive(w, r, param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetStartingWithNumber operation middleware
func (siw *ServerInterfaceWrapper) GetStartingWithNumber(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "1param" -------------
	var n1param string

	n1param = chi.URLParam(r, "1param")

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetStartingWithNumber(w, r, n1param)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter %s: %v", err.paramName, err.err)
}

func (err UnescapedCookieParamError) Unwrap() error { return err.err }

type UnmarshalingParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnmarshalingParamError) Error() string {
	return fmt.Sprintf("error unmarshaling parameter %s as JSON: %v", err.paramName, err.err)
}

func (err UnmarshalingParamError) Unwrap() error { return err.err }

type RequiredParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err RequiredParamError) Error() string {
	if err.err == nil {
		return fmt.Sprintf("query parameter %s is required, but not found", err.paramName)
	} else {
		return fmt.Sprintf("query parameter %s is required, but errored: %s", err.paramName, err.err)
	}
}

func (err RequiredParamError) Unwrap() error { return err.err }

type RequiredHeaderError struct {
	paramName string
}

// Error implements error.
func (err RequiredHeaderError) Error() string {
	return fmt.Sprintf("header parameter %s is required, but not found", err.paramName)
}

type InvalidParamFormatError struct {
	err       error
	paramName string
}

// Error implements error.
func (err InvalidParamFormatError) Error() string {
	return fmt.Sprintf("invalid format for parameter %s: %v", err.paramName, err.err)
}

func (err InvalidParamFormatError) Unwrap() error { return err.err }

type TooManyValuesForParamError struct {
	NumValues int
	paramName string
}

// Error implements error.
func (err TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("expected one value for %s, got %d", err.paramName, err.NumValues)
}

// ParameterName is an interface that is implemented by error types that are
// relevant to a specific parameter.
type ParameterError interface {
	error
	// ParamName is the name of the parameter that the error is referring to.
	ParamName() string
}

func (err UnescapedCookieParamError) ParamName() string  { return err.paramName }
func (err UnmarshalingParamError) ParamName() string     { return err.paramName }
func (err RequiredParamError) ParamName() string         { return err.paramName }
func (err RequiredHeaderError) ParamName() string        { return err.paramName }
func (err InvalidParamFormatError) ParamName() string    { return err.paramName }
func (err TooManyValuesForParamError) ParamName() string { return err.paramName }

type ServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

type ServerOption func(*ServerOptions)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface, opts ...ServerOption) http.Handler {
	options := &ServerOptions{
		BaseURL:    "/",
		BaseRouter: chi.NewRouter(),
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
	}

	for _, f := range opts {
		f(options)
	}

	r := options.BaseRouter
	wrapper := ServerInterfaceWrapper{
		Handler:          si,
		ErrorHandlerFunc: options.ErrorHandlerFunc,
	}

	r.Route(options.BaseURL, func(r chi.Router) {
		r.Get("/contentObject/{param}", wrapper.GetContentObject)
		r.Get("/cookie", wrapper.GetCookie)
		r.Get("/header", wrapper.GetHeader)
		r.Get("/labelExplodeArray/{param}", wrapper.GetLabelExplodeArray)
		r.Get("/labelExplodeObject/{param}", wrapper.GetLabelExplodeObject)
		r.Get("/labelNoExplodeArray/{param}", wrapper.GetLabelNoExplodeArray)
		r.Get("/labelNoExplodeObject/{param}", wrapper.GetLabelNoExplodeObject)
		r.Get("/matrixExplodeArray/{id}", wrapper.GetMatrixExplodeArray)
		r.Get("/matrixExplodeObject/{id}", wrapper.GetMatrixExplodeObject)
		r.Get("/matrixNoExplodeArray/{id}", wrapper.GetMatrixNoExplodeArray)
		r.Get("/matrixNoExplodeObject/{id}", wrapper.GetMatrixNoExplodeObject)
		r.Get("/passThrough/{param}", wrapper.GetPassThrough)
		r.Get("/queryDeepObject", wrapper.GetDeepObject)
		r.Get("/queryForm", wrapper.GetQueryForm)
		r.Get("/simpleExplodeArray/{param}", wrapper.GetSimpleExplodeArray)
		r.Get("/simpleExplodeObject/{param}", wrapper.GetSimpleExplodeObject)
		r.Get("/simpleNoExplodeArray/{param}", wrapper.GetSimpleNoExplodeArray)
		r.Get("/simpleNoExplodeObject/{param}", wrapper.GetSimpleNoExplodeObject)
		r.Get("/simplePrimitive/{param}", wrapper.GetSimplePrimitive)
		r.Get("/startingWithNumber/{1param}", wrapper.GetStartingWithNumber)
	})
	return r
}

func WithRouter(r chi.Router) ServerOption {
	return func(s *ServerOptions) {
		s.BaseRouter = r
	}
}

func WithServerBaseURL(url string) ServerOption {
	return func(s *ServerOptions) {
		s.BaseURL = url
	}
}

func WithErrorHandler(handler func(w http.ResponseWriter, r *http.Request, err error)) ServerOption {
	return func(s *ServerOptions) {
		s.ErrorHandlerFunc = handler
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xa34+jNhD+V6JpnyoSkrs33lbXXyv19q7NSlfptA9emARfAftsZ5tVxP9e2UAAQwgk",
	"YTfbtwvMzDfzefydPewOfBZzlmCiJHg7ECg5SySaH0sa8wj/yh/pJz5LFCZK/1PhVrk8IjTRv6QfYkzM",
	"82eO4IFUgiZrSNPUgQClLyhXlCXgwc1EmriTAmvCHr+hr0CbZnEM+gemrbafspfeDrhgHIWiWXK3QQWN",
	"JgrXKCB14FbeBHGWVP7ykbEISaJflsF+FLgCD35wy/rdHNz9VOYj8PuGCgzA+1o4Oxq6xHmoha3nuKJC",
	"qjsSYwsxDggWtb2wUI2VUwn1YDilyYpp54j6mC9OYoDg4+29jq6o0uHhHqWaLFE8oQAHnlDIbBkWs/ls",
	"rg0Zx4RwCh68n81nC3CAExWa/N18vbP63B0ngsSpfrNGU64uluh11asBv6H6UHUwoQSJUaGQ4H2t9Q/h",
	"PKK+cXa/SWZ1Udfy1BsjZwM8kzY4BQ0GGapcKrHB9MGp9/i7+fwQ3t7OtTZCajBdn7F/KHazYSwaNNQ3",
	"BBc0poo+aUPc8ogFCN6KRBLzwvwiTFEaOBWqVkzERGWb4P07cBp7InV6IWp6DgDi2Yg5SjAhQpDnvrCk",
	"BksVxrIX/v5JhtaSTyONLr7HS2NPCys2TC9eWC2hflJmQzcRuyg4DXGs7V6vxM8MSg5bK/AZNEnQ7yZS",
	"EaFosp78S1U4STbxo5HK1igLWSPClu66uiSbKDJKESIJUHQpxe+ZxblKERZh8nT/nn6uuIyqGR3Q01/y",
	"Nn8RFWkmcqOt25N4MU05kNUrK0szq2ybtZM1htAcyuDN6U2zkDxQUdAJ6mPHXEyXufX0C1Xh9K6wHqxI",
	"EXnEKF9k04jubmak56fO490ftltTsdrarM/J7DIbwQGpns2511QIlzzvVTkrTsRDSTt0ML4Ea312yej8",
	"3LG2rjrOT92vg6CqePyP+mpff72zBhB3tLXOYe61eysmStCt1Vo06N54HxtOp2w8GozeU1l14xG276lB",
	"jJ2uVUcoG9ZMo5HTkCoa9CDnAkL1ljuqqVPDWDtDpa69qziR8j4UbLMO+4zKPpfmnYOyAYPWVxmDfd+g",
	"eP4ZkZdT0EMlV6yO3HQDRN59dTGwZZ1BFvrkDrFO/WWjBGXOhw7TJpVfmYi7av9zb3Sk9F6XXKv6i03K",
	"yrq1Kwy85FpZvVhS/S67NmfjT9EsxEsA7ks9No+xqx1naNxR7eUAJ7nGHcDpHsm98ljASva0KaQVZNAQ",
	"8ixtz77U1Y9JPW68y4bb9c4JshJhNNZq384G0HY9k4LRGLIP4MfPTMsWvyueFYzPXP8vs8s2x6uYFozG",
	"0v4DRH9+qp9LLGZOYqJH84xJQ/5/yheqwmxW7O4WPahouI14QVmMfEPRDJu/fsjy3ogIPAiV4p7r4pZo",
	"+5nPYkgf0v8CAAD//0duf5cNIwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
