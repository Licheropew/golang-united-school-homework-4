package string_sum

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

type cases struct {
	inputString  string
	outputString string
	expectedErr  error
	numErr       bool
}

func TestStringSum(t *testing.T) {
	testCases := map[string]cases{
		"positive operands case":              {inputString: "3+5", outputString: "8", expectedErr: nil},
		"first negative operand case":         {inputString: "-3+5", outputString: "2", expectedErr: nil},
		"second negative operand case":        {inputString: "3-5", outputString: "-2", expectedErr: nil},
		"both negative operands case":         {inputString: "-3-5", outputString: "-8", expectedErr: nil},
		"empty string first case":             {inputString: "", outputString: "", expectedErr: errorEmptyInput},
		"empty string second case":            {inputString: "            ", outputString: "", expectedErr: errorEmptyInput},
		"whitespaces case":                    {inputString: " -3 - 5 ", outputString: "-8", expectedErr: nil},
		"one positive operand first case":     {inputString: "7", outputString: "", expectedErr: errorNotTwoOperands},
		"one positive operand second case":    {inputString: "+8", outputString: "", expectedErr: errorNotTwoOperands},
		"one positive operand third case":     {inputString: "8+", outputString: "", expectedErr: errorNotTwoOperands},
		"one negative operand first case":     {inputString: "-5", outputString: "", expectedErr: errorNotTwoOperands},
		"one negative operand second case":    {inputString: "5-", outputString: "", expectedErr: errorNotTwoOperands},
		"one negative operand third case":     {inputString: "-5-", outputString: "", expectedErr: errorNotTwoOperands},
		"three operands first case":           {inputString: "5+3+2", outputString: "", expectedErr: errorNotTwoOperands},
		"three operands seconde case":         {inputString: "5+3-", outputString: "", expectedErr: errorNotTwoOperands},
		"first operand is not a number case":  {inputString: "3f+5", outputString: "", expectedErr: &strconv.NumError{Func: "Atoi", Num: "3f", Err: strconv.ErrSyntax}, numErr: true},
		"second operand is not a number case": {inputString: "3+5f", outputString: "", expectedErr: &strconv.NumError{Func: "Atoi", Num: "5f", Err: strconv.ErrSyntax}, numErr: true},
	}

	for caseName, testCase := range testCases {
		t.Run(caseName, func(t *testing.T) {
			output, err := StringSum(testCase.inputString)
			if testCase.expectedErr != nil {
				if testCase.numErr {
					e := errors.Unwrap(err)
					if numerr, ok := e.(*strconv.NumError); !ok {
						t.Errorf("%s:\n wrong type of error is wrapped into the returned error: got %s, want %s", caseName, reflect.TypeOf(e), reflect.TypeOf(numerr))
					}
					if !errors.As(err, &testCase.expectedErr) {
						t.Errorf("%s:\n wrong error type is used in the return: got %T, want %T", caseName, err, testCase.expectedErr)
					}
				} else {
					if err == testCase.expectedErr {
						t.Errorf("%s:\n returned error must be wrapped", caseName)
					}
					if !errors.Is(err, testCase.expectedErr) {
						t.Errorf("%s:\n wrong error is used in the return: got %s, want %s", caseName, err.Error(), testCase.expectedErr.Error())
					}
				}
			} else {
				if err != nil {
					t.Errorf("error should be nil: got %s", err)
				}
			}

			if output != testCase.outputString {
				t.Errorf("error in the sum output: got %s, want %s", output, testCase.outputString)
			}

		})
	}
}
