// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

func main() {
  fmt.Println("2x2: Expecting 2")
  fmt.Println(robotPaths(2,2))
  
  fmt.Println("2x3: Expecting 3")
  fmt.Println(robotPaths(2,3))
  
  fmt.Println("3x3: Expecting 6")
  fmt.Println(robotPaths(3,3))
  
  fmt.Println("4x4: Expecting 20")
  fmt.Println(robotPaths(4,4))
}

func robotPaths(n, m int) int {
  return robotPathsIter(n, m, 0, 0, 0)
}

func robotPathsIter(n, m, x, y, acc int) int {
  // are we at the goal
  if x == n-1 && y == m-1 {
    return acc+1
  }
  
  var count int
  // move right
  if x != n-1 {
    c := robotPathsIter(n,m, x+1, y, acc)
    count += c
  }
  // move down
  if y != m-1 {
    c := robotPathsIter(n,m, x, y+1, acc)
    count += c
  }
  
  return acc + count
}

/* 
 
 Given a grid and a robot that starts from the top left corner of the grid, count the number of unique paths for the robot to reach bottom right corner of grid
 
 - Robot can only move down and to the right
 
 Example:
 
 -------
 |  |  |
 -------
 |  |  |
 -------
 
 In a 2x2 grid, you have two unique paths
 - Right, down
 - Down, right
 
 ----------
 |  |  |  |
 ----------
 |  |  |  |
 ----------
 
 In a 2x3 grid, you have 3 unique paths
 - Right, Right, down
 - Down, right right
 - Right, down, right
 
*/
