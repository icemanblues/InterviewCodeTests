# Run of 3 consecutive integers

## Problem Statement

Given an array of unsorted positive integers, write a function that finds runs of 3 consecutive numbers (ascending or descending) and returns the
indices where such runs begin. 

Increasing run is [1, 2, 3]  
Decreasing run is [3, 2, 1]

If no such runs are found, return null.

Implement the `find_consecutive_runs` function:
Input: array of numbers
Output: array of numbers

## Examples

### Simple

Simple Input: [0, 1, 2]
Returns: [0]

```
Run starting at index 0
[0, 1, 2]
 ^  ^  ^
```

### Given

Given Input: [1, 2, 3, 5, 10, 9, 8, 9, 10, 11, 7]
Returns: [0, 4, 6, 7]

```
Run starting at index 0
[1, 2, 3, 5, 10, 9, 8, 9, 10, 11, 7]
 ^  ^  ^
 
Run starting at index 4
[1, 2, 3, 5, 10, 9, 8, 9, 10, 11, 7]
              ^  ^  ^

Run starting at index 6
[1, 2, 3, 5, 10, 9, 8, 9, 10, 11, 7]
                    ^  ^   ^

Run starting at index 7
[1, 2, 3, 5, 10, 9, 8, 9, 10, 11, 7]
                       ^   ^   ^
```
