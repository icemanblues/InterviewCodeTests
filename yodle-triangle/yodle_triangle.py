# reads a file and returns the input triangle, as a list of lists
def load(file):
    triangle = []
    with open(file) as f:
        for line in f:
            line = line.strip()
            l = line.split()
            l = [int(i) for i in l]
            triangle.append(l)
    return triangle

# using top down recursion to solve for the max path
def max_path(triangle, r, c):
    if r >= len(triangle):
        return 0
    
    return triangle[r][c] + max(max_path(triangle, r+1, c), max_path(triangle, r+1, c+1))

# aggregating using a bottom up greedy method
def bottom_up(triangle):
    for r in range(len(triangle)-2, -1, -1):
        top = triangle[r]
        bot = triangle[r+1]
        for c in range(len(top)):
            top[c] += max(bot[c], bot[c+1])

    return triangle[0][0]


if __name__ == "__main__":
    print("hello triangle")
    sample_tri = load("sample.txt")
    print(sample_tri)

    print("solving simple top down")
    print(max_path(sample_tri, 0, 0))

    print("solving simple bottom up")
    print(bottom_up(sample_tri))

    print("solving HARD triangle bottom up")
    hard_tri = load("triangle.txt")
    print(bottom_up(hard_tri))
