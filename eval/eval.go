package eval

import (
	"log"
	"math"
	"strconv"
	"strings"
)

const notCalculated = math.MaxInt32

// Run runs the import and evaluate of system of equations
func Run(file string) (map[string]int, error) {
	lines := strings.Split(file, "\n")
	tokenizedEquation := tokenizeEquation(lines)
	return evaluate(tokenizedEquation)
}

func formatTokens(input string) []string {
	tokens := []string{}
	for _, token := range strings.Fields(input) {
		if !strings.Contains(token, "+") {
			tokens = append(tokens, token)
		}
	}
	return tokens
}

func tokenizeEquation(lines []string) map[string][]string {
	tokenizedEquation := make(map[string][]string)
	for _, line := range lines {
		sides := strings.Split(line, "=")
		tokenizedEquation[strings.TrimSpace(sides[0])] = formatTokens(strings.TrimSpace(sides[1]))
	}
	return tokenizedEquation
}

func isNumber(input string) bool {
	if _, err := strconv.Atoi(input); err != nil {
		return false
	}
	return true
}

func evaluate(tokenizedEquation map[string][]string) (map[string]int, error) {
	memo := make(map[string]int)
	for variable := range tokenizedEquation {
		memo[variable] = notCalculated
	}
	for variable := range tokenizedEquation {
		_, err := evaluateHelper(variable, tokenizedEquation, memo)
		if err != nil {
			return nil, err
		}
	}
	return memo, nil
}

func evaluateHelper(input string, tokenizedEquation map[string][]string, memo map[string]int) (int, error) {
	if isNumber(input) {
		// we made it to a base equation (a = 1) or similar
		val, err := strconv.Atoi(input)
		if err != nil {
			log.Printf("Error converting the string to int: %s", err)
			return val, err
		}
		return val, err
	} else if memo[input] != notCalculated {
		// we have calcuated this value before, use the memo
		return memo[input], nil
	}
	// the RHS of the equation has variables and we haven't see the LHS before
	// get the RHS tokens
	tokens := tokenizedEquation[input]
	sum := 0
	// for each RHS token, recursively call evaluateHelper and add the value to the sum
	for _, token := range tokens {
		val, err := evaluateHelper(token, tokenizedEquation, memo)
		if err != nil {
			return sum, err
		}
		sum += val
	}
	// store the sum for this input to save time
	memo[input] = sum
	return sum, nil
}
