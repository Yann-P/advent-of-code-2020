package main

import (
	"regexp"
	"strconv"
)

type Rule struct {
	from  string
	to    string
	count int
}

// shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
func ParseRule(ruleAsString string) []Rule {
	rules := []Rule{}

	if ruleAsString == "" {
		return rules
	}

	regexEmpty := regexp.MustCompile(`no other bags`)
	isEmptyBag := regexEmpty.MatchString(ruleAsString)

	if isEmptyBag {
		return rules
	}

	regexContainer := regexp.MustCompile(`(\w+ \w+) bags contain`)

	res := regexContainer.FindStringSubmatch(ruleAsString)

	if len(res) != 2 {
		panic("Can't extract container")
	}

	from := res[1] // shiny gold

	regexContents := regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)
	res2 := regexContents.FindAllStringSubmatch(ruleAsString, -1)

	for _, s := range res2 {
		match := s[1:]

		if len(res) != 2 {
			panic("Can't extract content")
		}

		count, err := strconv.Atoi(match[0]) // 1

		if err != nil {
			panic(err)
		}

		to := match[1] // dark olive

		rules = append(rules, Rule{
			from:  from,
			count: count,
			to:    to,
		})
	}

	return rules
}

type Node struct {
	next  []string
	count []int
}

func makeEmptyNode() *Node {
	return &Node{
		count: []int{},
		next:  []string{},
	}
}

func MakeRuleGraph(rules []Rule) map[string]*Node {
	nodes := make(map[string]*Node)

	for _, rule := range rules {
		if nodes[rule.from] == nil {
			nodes[rule.from] = makeEmptyNode()
		}
		if nodes[rule.to] == nil {
			nodes[rule.to] = makeEmptyNode()
		}
	}

	for _, rule := range rules {
		nodes[rule.from].count = append(nodes[rule.from].count, rule.count)
		nodes[rule.from].next = append(nodes[rule.from].next, rule.to)
	}

	return nodes
}

func MakeRuleGraphReversed(rules []Rule) map[string]*Node {
	nodes := make(map[string]*Node)

	for _, rule := range rules {
		if nodes[rule.from] == nil {
			nodes[rule.from] = makeEmptyNode()
		}
		if nodes[rule.to] == nil {
			nodes[rule.to] = makeEmptyNode()
		}
	}

	for _, rule := range rules {
		nodes[rule.to].count = append(nodes[rule.to].count, rule.count)
		nodes[rule.to].next = append(nodes[rule.to].next, rule.from)
	}

	return nodes
}

func SearchGraph(graph map[string]*Node, color string) []string {
	visited := make(map[string]bool)

	SearchGraphRec(graph, color, visited)

	v := []string{}

	for n := range visited {
		v = append(v, n)
	}

	return v
}

func SearchGraphRec(graph map[string]*Node, color string, visited map[string]bool) {
	visited[color] = true

	for _, node := range graph[color].next {
		if !visited[node] {
			SearchGraphRec(graph, node, visited)
		}
	}
}

func Count(graph map[string]*Node, color string) int {

	count := 0

	for i, neighbor := range graph[color].next {

		n := graph[color].count[i]

		count += n + n*Count(graph, neighbor)
	}

	return count

}
