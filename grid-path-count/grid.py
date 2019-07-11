def grid_count(n, m):
    return grid_count_helper(n, m, 0, 0)

def grid_count_helper(n, m, x, y):
    if x == n-1 and y == m-1:
        return 1
    
    count = 0
    if x < n-1:
        count += grid_count_helper(n, m, x+1, y)
    if y < m-1:
        count += grid_count_helper(n, m, x, y+1)

    return count

if __name__ == '__main__':
    print("2x2 => 2:")
    print(grid_count(2,2))

    print("2x3 => 3:")
    print(grid_count(2,3))

    print("3x3 => 6:")
    print(grid_count(3,3))

    print("4x4 => 20:")
    print(grid_count(4,4))
