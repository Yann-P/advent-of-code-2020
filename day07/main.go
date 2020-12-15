package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	rulesAsText := strings.Split(string(content), "\n")

	rules := []Rule{}

	for _, ruleAsText := range rulesAsText {
		parsedRules := ParseRule(ruleAsText)
		rules = append(rules, parsedRules...)
	}

	graph := MakeRuleGraphReversed(rules)

	search := SearchGraph(graph, "shiny gold")

	println("Part 1")
	println(len(search) - 1)

	println("Part 2")
	println(Count(MakeRuleGraph(rules), "shiny gold"))
}
