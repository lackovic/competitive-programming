/*
Author: Marco Lackovic
Date: 2024-12-08

https://adventofcode.com/2024/day/06

--- Day 6, part 2: Guard Gallivant ---

While The Historians begin working around the guard's patrol route, you borrow their fancy device and step outside the lab. From the safety of a supply closet, you time travel through the last few months and record the nightly status of the lab's guard post on the walls of the closet.

Returning after what seems like only a few seconds to The Historians, they explain that the guard's patrol area is simply too large for them to safely search the lab without getting caught.

Fortunately, they are pretty sure that adding a single new obstruction won't cause a time paradox. They'd like to place the new obstruction in such a way that the guard will get stuck in a loop, making the rest of the lab safe to search.

To have the lowest chance of creating a time paradox, The Historians would like to know all of the possible positions for such an obstruction. The new obstruction can't be placed at the guard's starting position - the guard is there right now and would notice.

In the above example, there are only 6 different positions where a new obstruction would cause the guard to get stuck in a loop. The diagrams of these six situations use O to mark the new obstruction, | to show a position where the guard moves up/down, - to show a position where the guard moves left/right, and + to show a position where the guard moves both up/down and left/right.

Option one, put a printing press next to the guard's starting position:

....#.....
....+---+#
....|...|.
..#.|...|.
....|..#|.
....|...|.
.#.O^---+.
........#.
#.........
......#...

Option two, put a stack of failed suit prototypes in the bottom right quadrant of the mapped area:


....#.....
....+---+#
....|...|.
..#.|...|.
..+-+-+#|.
..|.|.|.|.
.#+-^-+-+.
......O.#.
#.........
......#...

Option three, put a crate of chimney-squeeze prototype fabric next to the standing desk in the bottom right quadrant:

....#.....
....+---+#
....|...|.
..#.|...|.
..+-+-+#|.
..|.|.|.|.
.#+-^-+-+.
.+----+O#.
#+----+...
......#...

Option four, put an alchemical retroencabulator near the bottom left corner:

....#.....
....+---+#
....|...|.
..#.|...|.
..+-+-+#|.
..|.|.|.|.
.#+-^-+-+.
..|...|.#.
#O+---+...
......#...

Option five, put the alchemical retroencabulator a bit to the right instead:

....#.....
....+---+#
....|...|.
..#.|...|.
..+-+-+#|.
..|.|.|.|.
.#+-^-+-+.
....|.|.#.
#..O+-+...
......#...

Option six, put a tank of sovereign glue right next to the tank of universal solvent:

....#.....
....+---+#
....|...|.
..#.|...|.
..+-+-+#|.
..|.|.|.|.
.#+-^-+-+.
.+----++#.
#+----++..
......#O..

It doesn't really matter what you choose to use as an obstacle so long as you and The Historians can put it into position without the guard noticing. The important thing is having enough options that you can find one that minimizes time paradoxes, and in this example, there are 6 different positions you could choose.

You need to get the guard stuck in a loop by adding a single new obstruction. How many different positions could you choose for this obstruction?

*/

package day06

import (
	"adventofcode2024/utils"
	"fmt"
)

func Part2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	setLabAndGuard(lines)
	total := 0
	startI, startJ, startDirection := i, j, direction
	for {
		nextI := i + dirToPosChange[direction][0]
		nextJ := j + dirToPosChange[direction][1]
		lab[i][j] = 'X'
		if guardIsMovingOut(nextI, nextJ) {
			break
		}
		if lab[nextI][nextJ] != '#' && lab[nextI][nextJ] != 'X' {
			lab[nextI][nextJ] = '#'
			total += isStuckInLoop(startI, startJ, startDirection)
			lab[nextI][nextJ] = '.'
		}
		if guardIsFacingObstacle(nextI, nextJ) {
			direction = dirToDirChange[direction]
		} else {
			i = nextI
			j = nextJ
		}
	}
	return total, nil
}

// Returns 1 if the guard gets stuck in a loop, 0 otherwise.
func isStuckInLoop(i, j int, direction rune) int {
	visited := make([][]int, len(lab))
	for i := range lab {
		visited[i] = make([]int, len(lab[i]))
	}
	for {
		nextI := i + dirToPosChange[direction][0]
		nextJ := j + dirToPosChange[direction][1]
		if guardIsMovingOut(nextI, nextJ) {
			return 0
		}
		if visited[nextI][nextJ] == 4 {
			return 1
		}
		visited[i][j]++
		if guardIsFacingObstacle(nextI, nextJ) {
			direction = dirToDirChange[direction]
		} else {
			i = nextI
			j = nextJ
		}
	}
}
