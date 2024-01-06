/*
Author: Marco Lackovic
Date: 2024-01-06

https://adventofcode.com/2023/day/10

--- Day 10: Pipe Maze ---

You use the hang glider to ride the hot air from Desert Island all the way up to the floating metal island. This island is surprisingly cold and there definitely aren't any thermals to glide on, so you leave your hang glider behind.

You wander around for a while, but you don't find any people or animals. However, you do occasionally find signposts labeled "Hot Springs" pointing in a seemingly consistent direction; maybe you can find someone at the hot springs and ask them where the desert-machine parts are made.

The landscape here is alien; even the flowers and trees are made of metal. As you stop to admire some metal grass, you notice something metallic scurry away in your peripheral vision and jump into a big pipe! It didn't look like any animal you've ever seen; if you want a better look, you'll need to get ahead of it.

Scanning the area, you discover that the entire field you're standing on is densely packed with pipes; it was hard to tell at first because they're the same metallic silver color as the "ground". You make a quick sketch of all of the surface pipes you can see (your puzzle input).

The pipes are arranged in a two-dimensional grid of tiles:

| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

Based on the acoustics of the animal's scurrying, you're confident the pipe that contains the animal is one large, continuous loop.

For example, here is a square loop of pipe:

.....
.F-7.
.|.|.
.L-J.
.....

If the animal had entered this loop in the northwest corner, the sketch would instead look like this:

.....
.S-7.
.|.|.
.L-J.
.....

In the above diagram, the S tile is still a 90-degree F bend: you can tell because of how the adjacent pipes connect to it.

Unfortunately, there are also many pipes that aren't connected to the loop! This sketch shows the same loop as above:

-L|F7
7S-7|
L|7||
-L-J|
L|-JF

In the above diagram, you can still figure out which pipes form the main loop: they're the ones connected to S, pipes those pipes connect to, pipes those pipes connect to, and so on. Every pipe in the main loop connects to its two neighbors (including S, which will have exactly two pipes connecting to it, and which is assumed to connect back to those two pipes).

Here is a sketch that contains a slightly more complex main loop:

..F7.
.FJ|.
SJ.L7
|F--J
LJ...

Here's the same example sketch with the extra, non-main-loop pipe tiles also shown:

7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ

If you want to get out ahead of the animal, you should find the tile in the loop that is farthest from the starting position. Because the animal is in the pipe, it doesn't make sense to measure this by direct distance. Instead, you need to find the tile that would take the longest number of steps along the loop to reach from the starting point - regardless of which way around the loop the animal went.

In the first example with the square loop:

.....
.S-7.
.|.|.
.L-J.
.....

You can count the distance each tile in the loop is from the starting point like this:

.....
.012.
.1.3.
.234.
.....

In this example, the farthest point from the start is 4 steps away.

Here's the more complex loop again:

..F7.
.FJ|.
SJ.L7
|F--J
LJ...

Here are the distances for each tile on that loop:

..45.
.236.
01.78
14567
23...

Find the single giant loop starting at S. How many steps along the loop does it take to get from the starting position to the point farthest from the starting position?

*/

package day10

import (
	"adventofcode2023/utils"
	"errors"
)

func Solve1(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, err
	}
	start, err := findStart(lines)
	if err != nil {
		return 0, err
	}
	steps := 0
	for dir := range directions {
		steps = explore(lines, start, dir)
		if steps > 0 {
			return steps / 2, nil
		}
	}
	return 0, errors.New("no path leads back to the start")
}

// define the mapping of directions to points
var directions = map[rune]Point{
	'N': {X: 0, Y: -1},
	'E': {X: 1, Y: 0},
	'S': {X: 0, Y: 1},
	'W': {X: -1, Y: 0},
}

// Find the starting position.
func findStart(lines []string) (Point, error) {
	var start Point
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				return Point{X: x, Y: y}, nil
			}
		}
	}
	return start, errors.New("starting position not found")
}

// define the Point struct
type Point struct {
	X, Y int
}

// explore explores the maze, returns the number of steps taken if it returns to the start, or 0 if it doesn't.
func explore(lines []string, start Point, dir rune) int {
	steps := 0
	current := start
	for {
		steps++
		current = current.add(directions[dir])
		if current.outOfBounds(lines) {
			return 0
		}
		nextChar := lines[current.Y][current.X]
		if nextChar == 'S' {
			return steps
		}
		if nextChar == '.' {
			return 0
		}
		dir = nextDir(dir, nextChar)
		if dir == ' ' {
			return 0
		}
	}
}

// add adds two points together.
func (p Point) add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

// outOfBounds returns true if the point is out of bounds.
func (p Point) outOfBounds(lines []string) bool {
	return p.Y < 0 || p.Y >= len(lines) || p.X < 0 || p.X >= len(lines[p.Y])
}

// nextDir returns the next direction to explore.
func nextDir(sourceDir rune, nextChar byte) rune {
	switch sourceDir {
	case 'N':
		switch nextChar {
		case '7':
			return 'W'
		case 'F':
			return 'E'
		case '|':
			return 'N'
		default:
			return ' '
		}
	case 'E':
		switch nextChar {
		case '-':
			return 'E'
		case 'J':
			return 'N'
		case '7':
			return 'S'
		default:
			return ' '
		}
	case 'S':
		switch nextChar {
		case '|':
			return 'S'
		case 'L':
			return 'E'
		case 'J':
			return 'W'
		default:
			return ' '
		}
	case 'W':
		switch nextChar {
		case 'F':
			return 'S'
		case '-':
			return 'W'
		case 'L':
			return 'N'
		default:
			return ' '
		}
	}
	return ' '
}
