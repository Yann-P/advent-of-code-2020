package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParseRules(t *testing.T) {

	rules := []Rule{
		{
			from:  "shiny gold",
			to:    "vibrant plum",
			count: 2,
		},
		{
			from:  "shiny gold",
			to:    "dark olive",
			count: 1,
		},
	}

	if reflect.DeepEqual(
		ParseRule("shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags."),
		rules,
	) {
		t.Error("1")
	}
}

func S(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), "\t", "")
}

func TestMakeRuleGraph(t *testing.T) {

	rules := []Rule{
		{
			from:  "shiny gold",
			to:    "vibrant plum",
			count: 2,
		},
		{
			from:  "shiny gold",
			to:    "dark olive",
			count: 1,
		},
		{
			from:  "dark olive",
			to:    "shiny gold",
			count: 3,
		},
	}

	g := MakeRuleGraph(rules)

	if fmt.Sprintf("%#v", g["shiny gold"]) != S(`&main.Node{
		next:[]string{"vibrant plum", "dark olive"}, 
		count:[]int{2, 1}
	}`) {
		t.Error(fmt.Sprintf("%#v", g["shiny gold"]))
	}

	if fmt.Sprintf("%#v", g["dark olive"]) != S(`&main.Node{
		next:[]string{"shiny gold"}, 
		count:[]int{3}
	}`) {
		t.Error(fmt.Sprintf("%#v", g["dark olive"]))
	}

	if fmt.Sprintf("%#v", g["vibrant plum"]) != S(`&main.Node{next:[]string{}, count:[]int{}}`) {
		t.Error("3")
	}

}

func TestMakeRuleGraphReversed(t *testing.T) {

	rules := []Rule{
		{
			from:  "shiny gold",
			to:    "vibrant plum",
			count: 2,
		},
		{
			from:  "shiny gold",
			to:    "dark olive",
			count: 1,
		},
		{
			from:  "dark olive",
			to:    "shiny gold",
			count: 3,
		},
	}

	g := MakeRuleGraphReversed(rules)

	if fmt.Sprintf("%#v", g["shiny gold"]) != S(`&main.Node{
		next:[]string{"dark olive"}, 
		count:[]int{3}
	}`) {
		t.Error("1")
	}

	if fmt.Sprintf("%#v", g["dark olive"]) != S(`&main.Node{
		next:[]string{"shiny gold"}, 
		count:[]int{1}
	}`) {
		t.Error("2")
	}

	if fmt.Sprintf("%#v", g["vibrant plum"]) != S(`&main.Node{next:[]string{"shiny gold"}, count:[]int{2}}`) {
		t.Error("3")
	}

}

func TestSearchGraph(t *testing.T) {

	g := map[string]*Node{
		// light red bags contain 1 bright white bag, 2 muted yellow bags.
		"lr": {count: []int{1, 2}, next: []string{"bw", "my"}},
		// dark orange bags contain 3 bright white bags, 4 muted yellow bags.
		"do": {count: []int{3, 4}, next: []string{"bw", "my"}},
		// bright white bags contain 1 shiny gold bag.
		"bw": {count: []int{1}, next: []string{"sg"}},
		// muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
		"my": {count: []int{2, 9}, next: []string{"sg", "fb"}},
		// shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
		"sg": {count: []int{1, 2}, next: []string{"dol", "vp"}},
		// dark olive bags contain 3 faded blue bags, 4 dotted black bags.
		"dol": {count: []int{3, 4}, next: []string{"fb", "db"}},
		// vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
		"vp": {count: []int{5, 6}, next: []string{"fb", "db"}},
		// faded blue bags contain no other bags.
		"fb": {count: []int{}, next: []string{}},
		// dotted black bags contain no other bags.
		"db": {count: []int{}, next: []string{}},
	}
	r := SearchGraph(g, "sg")
	if fmt.Sprintf("%#v", r) != `[]string{"sg", "dol", "fb", "db", "vp"}` {
		t.Error("1")
	}
}

func TestSearchGraphCount(t *testing.T) {

	rulesAsText := []string{"shiny gold bags contain 2 dark red bags.",
		"dark red bags contain 2 dark orange bags.",
		"dark orange bags contain 2 dark yellow bags.",
		"dark yellow bags contain 2 dark green bags.",
		"dark green bags contain 2 dark blue bags.",
		"dark blue bags contain 2 dark violet bags.",
		"dark violet bags contain no other bags.",
	}

	rules := []Rule{}

	for _, ruleAsText := range rulesAsText {
		parsedRules := ParseRule(ruleAsText)
		rules = append(rules, parsedRules...)
	}

	graph := MakeRuleGraph(rules)

	println(Count(graph, "shiny gold"))
}
