package main

import "sort"

// Problem: You have to make a change of an amount using the smallest possible number of coins.
// Amount: $18

// Available coins are
//   $5 coin
//   $2 coin
//   $1 coin
// There is no limit to the number of each coin you can use.

func solutionAbove() {}

// 452. Minimum Number of Arrows to Burst Balloons

// There are some spherical balloons taped onto a flat wall that represents the XY-plane. The balloons are represented as a 2D integer array points where points[i] = [xstart, xend] denotes a balloon whose horizontal diameter stretches between xstart and xend. You do not know the exact y-coordinates of the balloons.

// Arrows can be shot up directly vertically (in the positive y-direction) from different points along the x-axis. A balloon with xstart and xend is burst by an arrow shot at x if xstart <= x <= xend. There is no limit to the number of arrows that can be shot. A shot arrow keeps traveling up infinitely, bursting any balloons in its path.

// Given the array points, return the minimum number of arrows that must be shot to burst all balloons.

// Example 1:

// Input: points = [[10,16],[2,8],[1,6],[7,12]]
// Output: 2
// Explanation: The balloons can be burst by 2 arrows:
// - Shoot an arrow at x = 6, bursting the balloons [2,8] and [1,6].
// - Shoot an arrow at x = 11, bursting the balloons [10,16] and [7,12].

// Example 2:

// Input: points = [[1,2],[3,4],[5,6],[7,8]]
// Output: 4
// Explanation: One arrow needs to be shot for each balloon for a total of 4 arrows.

// Example 3:

// Input: points = [[1,2],[2,3],[3,4],[4,5]]
// Output: 2
// Explanation: The balloons can be burst by 2 arrows:
// - Shoot an arrow at x = 2, bursting the balloons [1,2] and [2,3].
// - Shoot an arrow at x = 4, bursting the balloons [3,4] and [4,5].

// Constraints:

//     1 <= points.length <= 105
//     points[i].length == 2
//     -231 <= xstart < xend <= 231 - 1

func findMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	// sort the points by left, order by right
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})
	// store the current target range (left, right) for arrow
	cnt := 0
	left := points[0][0]
	right := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > right {
			// if the current ballon is not in the range of (left, right)
			// fire an arrow to the target range to hit the group of ballons within the range (left, right)
			cnt++
			// update the range for the next arrow
			left = points[i][0]
			right = points[i][1]
		} else {
			// if current ballon is within the range of (left, right)
			// update the range if needed
			left = findMax(left, points[i][0])
			right = findMin(right, points[i][1])
		}
	}
	// +1: at the end, fire the last arrow to hit the last group of ballons
	return cnt + 1
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
