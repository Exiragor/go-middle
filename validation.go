package main

import "reflect"

func MasterRegistrationValidate(master Master) string {
	// required fields in request
	fields := []string{"Phone", "Password", "Email"}
	errFields := validateFields(master, fields)

	return errFields
}

func validateFields(obj interface{}, requiredFields []string) string {
	v := reflect.ValueOf(obj)
	strIncompleteElems := ""
	// check required fields for empty val
	for _, elem := range requiredFields {
		value := v.FieldByName(elem).Interface()
		if value == "" {
			strIncompleteElems += elem + " is incorrect; "
		}
	}
	return strIncompleteElems
}
