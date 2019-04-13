def load(file):
    triangle = []
    with open(file) as f:
        for line in f:
            line = line.strip()
            l = line.split()
            l = [int(i) for i in l]
            triangle.append(l)
    return triangle


def max_path(triangle, r, c):
    if r >= len(triangle):
        return 0
    
    return triangle[r][c] + max(max_path(triangle, r+1, c), max_path(triangle, r+1, c+1))

if __name__ == "__main__":
    print("hello triangle")
    tri = load("sample.txt")
    print(tri)
    print("solving triangle")
    s = max_path(tri, 0, 0)
    print(s)
    

            
