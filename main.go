package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/gregmankes/solver/eval"
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")
	if _, err := os.Stat(flag.Arg(0)); err != nil {
		log.Fatalf("Error looking up the input file: %s", err)
	}
	bytes, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatalf("Error reading the input file: %s", err)
	}
	evaluatedEquations, err := eval.Run(string(bytes))
	if err != nil {
		log.Fatalf("Error importing and evaluating the file: %s", err)
	}
	printOutput(evaluatedEquations)
}

func printOutput(output map[string]int) {
	sortedVariables := []string{}
	for variable := range output {
		sortedVariables = append(sortedVariables, variable)
	}
	sort.Strings(sortedVariables)
	for _, variable := range sortedVariables {
		fmt.Printf("%s = %d\n", variable, output[variable])
	}
}
