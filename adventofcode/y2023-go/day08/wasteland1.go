/*
Author: Marco Lackovic
Date: 2023-12-23

https://adventofcode.com/2023/day/8

--- Day 8: Haunted Wasteland ---

You're still riding a camel across Desert Island when you spot a sandstorm quickly approaching. When you turn to warn the Elf, she disappears before your eyes! To be fair, she had just finished warning you about ghosts a few minutes ago.

One of the camel's pouches is labeled "maps" - sure enough, it's full of documents (your puzzle input) about how to navigate the desert. At least, you're pretty sure that's what they are; one of the documents contains a list of left/right instructions, and the rest of the documents seem to describe some kind of network of labeled nodes.

It seems like you're meant to use the left/right instructions to navigate the network. Perhaps if you have the camel follow the same instructions, you can escape the haunted wasteland!

After examining the maps for a bit, two nodes stick out: AAA and ZZZ. You feel like AAA is where you are now, and you have to follow the left/right instructions until you reach ZZZ.

This format defines each node of the network individually. For example:

RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)

Starting with AAA, you need to look up the next element based on the next left/right instruction in your input. In this example, start with AAA and go right (R) by choosing the right element of AAA, CCC. Then, L means to choose the left element of CCC, ZZZ. By following the left/right instructions, you reach ZZZ in 2 steps.

Of course, you might not find ZZZ right away. If you run out of left/right instructions, repeat the whole sequence of instructions as necessary: RL really means RLRLRLRLRLRLRLRL... and so on. For example, here is a situation that takes 6 steps to reach ZZZ:

LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)

Starting at AAA, follow the left/right instructions. How many steps are required to reach ZZZ?

*/

package day08

import (
	"adventofcode2023/utils"
	"strings"
)

func Solve1(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}
	directions := parseDirectionsAsBits(lines[0])
	nodes := createNodes(lines)
	currentNode := nodes["AAA"]
	steps := 1
	for next := 0; ; steps++ {
		if currentNode[directions[next]] == "ZZZ" {
			break
		}
		currentNode = nodes[currentNode[directions[next]]]
		next = (next + 1) % len(directions)
	}
	return steps, nil
}

// parseDirectionsAsBits converts a string of directions as L and R to an array of bits
func parseDirectionsAsBits(directions string) []uint8 {
	result := make([]uint8, len(directions))
	for i, char := range directions {
		if char == 'L' {
			result[i] = 0
		} else {
			result[i] = 1
		}
	}
	return result
}

func createNodes(lines []string) map[string][]string {
	nodes := make(map[string][]string)
	for _, line := range lines[1:] {
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		node := strings.Split(line, "=")[0]
		children := strings.Split(line, "=")[1]
		leftChild := strings.Split(children, ",")[0]
		rightChild := strings.Split(children, ",")[1]
		nodes[node] = []string{leftChild, rightChild}
	}
	return nodes
}
