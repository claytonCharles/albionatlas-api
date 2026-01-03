package validator

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func FormValidate(r *http.Request, stt any) map[string][]string {
	errors := make(map[string][]string)

	if err := r.ParseForm(); err != nil {
		errors["form"] = []string{"Formulário Invalido!"}
		return errors
	}

	v := reflect.ValueOf(stt).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fType := t.Field(i)

		formKey := fType.Tag.Get("form")
		rules := fType.Tag.Get("validate")

		value := r.FormValue(formKey)
		field.SetString(value)
		validateField(formKey, value, rules, errors)
	}

	return errors
}

func validateField(key string, value string, rules string, errors map[string][]string) {
	if rules == "" {
		return
	}

	for rule := range strings.SplitSeq(rules, "|") {
		switch {
		case rule == "required":
			if strings.TrimSpace(value) == "" {
				errors[key] = append(errors[key], "Field is required!")
			}
		case rule == "mail":
			if value != "" && !regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`).MatchString(value) {
				errors[key] = append(errors[key], "Invalid mail!")
			}
		case strings.HasPrefix(rule, "min="):
			min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
			if len(value) < min {
				errors[key] = append(errors[key], fmt.Sprintf("Min %d characters", min))
			}
		}
	}
}
