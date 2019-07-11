# Grid Path Count

## Problem Statement

Given a grid and a robot that starts from the top left corner of the grid, count the number of unique paths for the robot to reach bottom right corner of grid

- Robot can only move down and to the right

## Examples

### 2x2

In a 2x2 grid, you have two unique paths:

```
-------
|  |  |
-------
|  |  |
-------
```

- Right, down
- Down, right

### 2x3

In a 2x3 grid, you have 3 unique paths:

```
----------
|  |  |  |
----------
|  |  |  |
----------
```

- Right, Right, down
- Down, right right
- Right, down, right

## Additional Test Cases

- 3x3 => 6
- 4x4 => 20