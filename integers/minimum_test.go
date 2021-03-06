package integers_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/integers"
)

func TestNewMinimumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition integers.MinimumValidatorDefinition
	}
	cases := []Case{
		{
			Message: "positive number and positive exclusive",
			Definition: integers.MinimumValidatorDefinition{
				Minimum:   10,
				Exclusive: true,
			},
		},
		{
			Message: "positive number and negative exclusive",
			Definition: integers.MinimumValidatorDefinition{
				Minimum:   10,
				Exclusive: false,
			},
		},
		{
			Message: "zero and positive exclusive",
			Definition: integers.MinimumValidatorDefinition{
				Minimum:   0.0,
				Exclusive: true,
			},
		},
		{
			Message: "zero number and negative exclusive",
			Definition: integers.MinimumValidatorDefinition{
				Minimum:   0.0,
				Exclusive: false,
			},
		},
		{
			Message: "negative number and positive exclusive",
			Definition: integers.MinimumValidatorDefinition{
				Minimum:   -10,
				Exclusive: true,
			},
		},
		{
			Message: "negative number and negative exclusive",
			Definition: integers.MinimumValidatorDefinition{
				Minimum:   -10,
				Exclusive: false,
			},
		},
	}
	for _, c := range cases {
		_, err := integers.NewMinimumValidator(c.Definition)
		if err != nil {
			t.Errorf("Test with %s: fail to NewMaximumValidator with error %v", c.Message, err)
		}
	}
}

type MinimumValidatorTestCase struct {
	Message string
	Input   int
	Error   error
}

func TestValidateOfMinimumValidatorWithPositiveNumberAndPositiveExclusive(t *testing.T) {
	def := integers.MinimumValidatorDefinition{
		Minimum:   10,
		Exclusive: true,
	}

	v, err := integers.NewMinimumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMinimumValidator: %s", err)
	}

	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   11,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   10,
			Error: &integers.MinimumValidationError{
				Input:      10,
				Definition: def,
			},
		},
		{
			Message: "less number",
			Input:   9,
			Error: &integers.MinimumValidationError{
				Input:      9,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMinimumValidatorWithPositiveNumberAndNegativeExclusive(t *testing.T) {
	def := integers.MinimumValidatorDefinition{
		Minimum:   10,
		Exclusive: false,
	}

	v, err := integers.NewMinimumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   11,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   10,
			Error:   nil,
		},
		{
			Message: "less number",
			Input:   9,
			Error: &integers.MinimumValidationError{
				Input:      9,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMinimumValidatorWithZeroAndPositiveExclusive(t *testing.T) {
	def := integers.MinimumValidatorDefinition{
		Minimum:   0.0,
		Exclusive: true,
	}

	v, err := integers.NewMinimumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMinimumValidator: %s", err)
	}

	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   0.0,
			Error: &integers.MinimumValidationError{
				Input:      0.0,
				Definition: def,
			},
		},
		{
			Message: "less number",
			Input:   -1,
			Error: &integers.MinimumValidationError{
				Input:      -1,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMinimumValidatorWithZeroAndNegativeExclusive(t *testing.T) {
	def := integers.MinimumValidatorDefinition{
		Minimum:   0.0,
		Exclusive: false,
	}

	v, err := integers.NewMinimumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   0.0,
			Error:   nil,
		},
		{
			Message: "less number",
			Input:   -1,
			Error: &integers.MinimumValidationError{
				Input:      -1,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMinimumValidatorWithNegativeNumberAndPositiveExclusive(t *testing.T) {
	def := integers.MinimumValidatorDefinition{
		Minimum:   -10,
		Exclusive: true,
	}

	v, err := integers.NewMinimumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMinimumValidator: %s", err)
	}

	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   -9,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   -10,
			Error: &integers.MinimumValidationError{
				Input:      -10,
				Definition: def,
			},
		},
		{
			Message: "less number",
			Input:   -11,
			Error: &integers.MinimumValidationError{
				Input:      -11,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMinimumValidatorWithNegativeNumberAndNegativeExclusive(t *testing.T) {
	def := integers.MinimumValidatorDefinition{
		Minimum:   -10,
		Exclusive: false,
	}

	v, err := integers.NewMinimumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   -9,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   -10,
			Error:   nil,
		},
		{
			Message: "less number",
			Input:   -11,
			Error: &integers.MinimumValidationError{
				Input:      -11,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}
