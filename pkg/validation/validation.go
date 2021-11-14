package validation

import ozvalidation "github.com/go-ozzo/ozzo-validation"

func requiredIf(cond bool) ozvalidation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return ozvalidation.Validate(value, ozvalidation.Required)
		}
		return nil
	}
}

func repeatPassword(pass string, repeatPass string) ozvalidation.RuleFunc {
	return func(value interface{}) error {
		if pass != repeatPass {
			return ozvalidation.Validate(value, ozvalidation.Required)
		}
		return nil
	}
}
