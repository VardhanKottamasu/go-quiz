package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Hello, World")
	csvFileName := flag.String("csv", "problems.csv", "A csv file in the format fo 'questions,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse csv")
	}
	problems:=parseLines(lines)
	// fmt.Println(problems)
	correct:=0
	for i,p:=range problems{
		fmt.Printf("%d. What is %s\t",i+1,p.q)
		var answer string
		fmt.Scanf("%s",&answer)
		if p.a!=answer{
			fmt.Println("Wrong Answer!")
			// correct++
		}else{
			fmt.Println("Correct")
			correct++
		}
	}
	fmt.Printf("Scored %d out of %d\n",correct,len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}
