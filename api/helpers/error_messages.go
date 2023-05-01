package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func tag2Message(e validator.FieldError) string {
	tag := strings.Split(e.Tag(), "|")[0]
	switch tag {
	case "required":
		return "field required"
	case "required_without":
		return fmt.Sprintf("is required if %s is not supplied", e.Param())
	case "lt", "ltfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("must be less than %s", param)
	case "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("must be greater than %s", param)
	default:
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("en"); found {
			return e.Translate(translatorInstance)
		} else {
			return fmt.Errorf("%v", e).Error()
		}
	}
}

func getJsonField(s interface{}, fieldName string) string {
	if reflect.ValueOf(s).Kind() != reflect.Struct {
		field, ok := reflect.TypeOf(s).Elem().FieldByName(fieldName)
		if !ok {
			return fieldName
		}
		value := field.Tag.Get("json")
		if value == "" {
			return fieldName
		}
		return field.Tag.Get("json")
	}
	return fieldName
}

type ValidationError struct {
	Messages map[string][]string
}

func NewValidationError() *ValidationError {
	e := ValidationError{
		Messages: make(map[string][]string),
	}
	return &e
}

func (e *ValidationError) AddMessage(field string, message string) {
	_, exists := e.Messages[field]
	if !exists {
		e.Messages[field] = []string{message}
	} else {
		e.Messages[field] = append(e.Messages[field], message)
	}
}

func (e *ValidationError) ParseError(err error, schema interface{}) {

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range validationErrs {
			field := getJsonField(schema, ve.Field())
			e.AddMessage(field, tag2Message(ve))
		}
		return
	}
	if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok {
		e.AddMessage(marshallingErr.Field, fmt.Sprintf("must be a %s", marshallingErr.Type.String()))
		return
	}
	e.AddMessage("_schema", err.Error())
}
