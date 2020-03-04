//
//  Practicing Redis
//
//  Copyright © 2016. All rights reserved.
//

package form

import (
	"time"
)

type KeyValueForm struct {
	Key    string        `json:"key"`
	Value  interface{}   `json:"value"`
	Expiry time.Duration `json:"expiry"`
}

// Validate represent the validation method from KeyValueForm
func (v *KeyValueForm) Validate() []string {
	errs := []string{}
	if len(v.Key) < 1 {
		errs = append(errs, "Key can't be empty")
	}

	if v.Value == nil {
		errs = append(errs, "Value can't be empty")
	}

	return errs
}

type KeyValuesForm struct {
	Key    string        `json:"key"`
	Values []interface{} `json:"values"`
}

// Validate represent the validation method from KeyValuesForm
func (v *KeyValuesForm) Validate() []string {
	errs := []string{}
	if len(v.Key) < 1 {
		errs = append(errs, "Key can't be empty")
	}

	if len(v.Values) < 1 {
		errs = append(errs, "Values can't be empty")
	}

	for _, val := range v.Values {
		if val == nil {
			errs = append(errs, "Value can't be empty")
		}
	}

	return errs
}

type KeyFieldValueForm struct {
	Key   string `json:"key"`
	Field string `json:"field"`
	Value string `json:"value"`
}

// Validate represent the validation method from KeyValueForm
func (v *KeyFieldValueForm) Validate() []string {
	errs := []string{}
	if len(v.Key) < 1 {
		errs = append(errs, "Key can't be empty")
	}

	if len(v.Field) < 1 {
		errs = append(errs, "Field can't be empty")
	}

	if len(v.Value) < 1 {
		errs = append(errs, "Value can't be empty")
	}

	return errs
}
