package eval

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"testing"
)

func TestTokenizeEquation(t *testing.T) {
	type testcase struct {
		input  []string
		output map[string][]string
	}
	testcases := []testcase{
		testcase{
			input: []string{"foo = bar", "bar = baz + 1", "baz = 2"},
			output: map[string][]string{
				"foo": []string{"bar"},
				"bar": []string{"baz", "1"},
				"baz": []string{"2"},
			},
		},
		testcase{
			input: []string{"foo=bar", "bar=baz + 1", "baz=2"},
			output: map[string][]string{
				"foo": []string{"bar"},
				"bar": []string{"baz", "1"},
				"baz": []string{"2"},
			},
		},
	}
	for _, tc := range testcases {
		if !reflect.DeepEqual(tc.output, tokenizeEquation(tc.input)) {
			t.Fail()
		}
	}
}

func TestEvaluate(t *testing.T) {
	type testcase struct {
		input  map[string][]string
		output map[string]int
	}
	testcases := []testcase{
		testcase{
			input: map[string][]string{
				"foo": []string{"bar"},
				"bar": []string{"baz", "1"},
				"baz": []string{"2"},
			},
			output: map[string]int{
				"foo": 3,
				"bar": 3,
				"baz": 2,
			},
		},
	}
	for _, tc := range testcases {
		evaluatedEquations, err := evaluate(tc.input)
		if err != nil {
			log.Printf("Error while evaluating the equations")
			t.Fail()
		}
		fmt.Printf("Input: %v: Ouput: %v\n", tc.input, evaluatedEquations)
		if !reflect.DeepEqual(tc.output, evaluatedEquations) {
			log.Printf("Equations are not equal")
			t.Fail()
		}
	}
}

func TestIsNumber(t *testing.T) {
	type testcase struct {
		in  string
		out bool
	}
	testcases := []testcase{
		testcase{"1", true},
		testcase{fmt.Sprintf("%d", math.MaxInt32), true},
		testcase{"a", false},
		testcase{"+", false},
	}
	for _, tc := range testcases {
		if tc.out != isNumber(tc.in) {
			log.Printf("Incorrect response from isNumber")
			t.Fail()
		}
	}
}
