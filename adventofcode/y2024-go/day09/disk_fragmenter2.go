/*
Author: Marco Lackovic
Date: 2024-12-15

https://adventofcode.com/2024/day/09

--- Day 9, part 2: Disk Fragmenter ---

Upon completion, two things immediately become clear. First, the disk definitely has a lot more contiguous free space, just like the amphipod hoped. Second, the computer is running much more slowly! Maybe introducing all of that file system fragmentation was a bad idea?

The eager amphipod already has a new plan: rather than move individual blocks, he'd like to try compacting the files on his disk by moving whole files instead.

This time, attempt to move whole files to the leftmost span of free space blocks that could fit the file. Attempt to move each file exactly once in order of decreasing file ID number starting with the file with the highest file ID number. If there is no span of free space to the left of a file that is large enough to fit the file, the file does not move.

The first example from above now proceeds differently:

00...111...2...333.44.5555.6666.777.888899
0099.111...2...333.44.5555.6666.777.8888..
0099.1117772...333.44.5555.6666.....8888..
0099.111777244.333....5555.6666.....8888..
00992111777.44.333....5555.6666.....8888..

The process of updating the filesystem checksum is the same; now, this example's checksum would be 2858.

Start over, now compacting the amphipod's hard drive using this new method instead. What is the resulting filesystem checksum?

*/

package day09

import (
	"adventofcode2024/utils"
	"fmt"
)

func Part2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	diskMap := lines[0]
	individualBlocks := diskMapToIndividualBlocks(diskMap)
	defragment2(individualBlocks)
	total := 0
	for i, block := range individualBlocks {
		if block == -1 {
			continue
		}
		total += i * block
	}
	return total, nil
}

// defragment2 moves whole files to the leftmost span of free space blocks that could fit the file
func defragment2(individualBlocks []int) {
	for j := len(individualBlocks) - 1; j > 0; {
		if individualBlocks[j] == -1 {
			j--
			continue
		}
		size := computeSize(individualBlocks, j)
		i := findFirstSlot(individualBlocks, size, j)
		if i != -1 {
			for k := 0; k < size; k++ {
				individualBlocks[i+k] = individualBlocks[j]
			}
			for k := j; k > j-size; k-- {
				individualBlocks[k] = -1
			}
		}
		j = j - size
	}
}

// computeSize computes the size of the file starting at the given block going to the left
func computeSize(individualBlocks []int, j int) int {
	size := 1
	for i := j - 1; i >= 0; i-- {
		if individualBlocks[i] != individualBlocks[j] {
			break
		}
		size++
	}
	return size
}

// findFirstSlot returns the index of the first slot of free space that could fit a file of the given size, otherwise returns -1
func findFirstSlot(individualBlocks []int, size, max int) int {
	for i := 0; i < max-size; i++ {
		if individualBlocks[i] == -1 {
			if isLargeEnough(individualBlocks, i, size) {
				return i
			}
			for j := i + 1; j < i+size; j++ {
				if individualBlocks[j] != -1 {
					i = j - 1
					break
				}
			}
		}
	}
	return -1
}

// isLargeEnough returns true if the span of free space is large enough to fit the file, otherwise returns false
func isLargeEnough(individualBlocks []int, i, size int) bool {
	for j := i; j < i+size; j++ {
		if individualBlocks[j] != -1 {
			return false
		}
	}
	return true
}
