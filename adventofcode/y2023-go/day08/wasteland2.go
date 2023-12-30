/*
Author: Marco Lackovic
Date: 2023-12-23

https://adventofcode.com/2023/day/8

--- Part Two ---

The sandstorm is upon you and you aren't any closer to escaping the wasteland. You had the camel follow the instructions, but you've barely left your starting position. It's going to take significantly more steps to escape!

What if the map isn't for people - what if the map is for ghosts? Are ghosts even bound by the laws of spacetime? Only one way to find out.

After examining the maps a bit longer, your attention is drawn to a curious fact: the number of nodes with names ending in A is equal to the number ending in Z! If you were a ghost, you'd probably just start at every node that ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z.

For example:

LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)

Here, there are two starting nodes, 11A and 22A (because they both end with A). As you follow each left/right instruction, use that instruction to simultaneously navigate away from both nodes you're currently on. Repeat this process until all of the nodes you're currently on end with Z. (If only some of the nodes you're on end with Z, they act like any other node and you continue as normal.) In this example, you would proceed as follows:

Step 0: You are at 11A and 22A.
Step 1: You choose all of the left paths, leading you to 11B and 22B.
Step 2: You choose all of the right paths, leading you to 11Z and 22C.
Step 3: You choose all of the left paths, leading you to 11B and 22Z.
Step 4: You choose all of the right paths, leading you to 11Z and 22B.
Step 5: You choose all of the left paths, leading you to 11B and 22C.
Step 6: You choose all of the right paths, leading you to 11Z and 22Z.

So, in this example, you end up entirely on nodes that end in Z after 6 steps.

Simultaneously start on every node that ends with A. How many steps does it take before you're only on nodes that end with Z?

*/

package day08

import (
	"adventofcode2023/utils"
	"strings"
)

func Solve2(filename string) (int64, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}
	directions := parseDirectionsAsBits(lines[0])
	nodes := createNodes(lines)
	startingNodes := findStartingNodes(nodes)
	endsWithZCh := make(chan bool)
	directionCh := make(chan uint8)
	for _, startNode := range startingNodes {
		go func(node []string) {
			currentNode := node
			for {
				nextNodeKey := currentNode[<-directionCh]
				currentNode = nodes[nextNodeKey]
				endsWithZCh <- strings.HasSuffix(nextNodeKey, "Z")
			}
		}(startNode)
	}
	steps := int64(0)
	max := 0
	for i := 0; ; i = (i + 1) % len(directions) {
		steps++
		if steps%1000000 == 0 {
			println(steps)
		}
		for range startingNodes {
			directionCh <- directions[i]
		}
		endsWithZNo := 0
		for range startingNodes {
			if <-endsWithZCh {
				endsWithZNo++
			}
		}
		if endsWithZNo == len(startingNodes) {
			break
		} else if endsWithZNo > max {
			max = endsWithZNo
			println("max = ", max, " out of ", len(startingNodes))
		}
	}
	return steps, nil
}

// findStartingNodes returns all nodes that end with A
func findStartingNodes(nodes map[string][]string) map[string][]string {
	startingNodes := make(map[string][]string)
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			startingNodes[node] = nodes[node]
		}
	}
	return startingNodes
}
