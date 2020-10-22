package models

import validation "github.com/go-ozzo/ozzo-validation"

func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}

func repeatPassword(pass string, repeatPass string) validation.RuleFunc {
	return func(value interface{}) error {
		if pass != repeatPass {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
