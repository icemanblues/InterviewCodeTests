def jumble_sort(line):
    elements = line.split()
    pos = []
    s = []
    n = []
    for e in elements:
        try:
            num = int(e)
            pos.append(True)
            n.append(num)
        except ValueError:
            pos.append(False)
            s.append(e)

    s.sort()
    n.sort()

    r = []
    for t in pos:
        if t:
            r.append(str(n.pop(0)))
        else:
            r.append(s.pop(0))

    return r

if __name__ == "__main__":
    print("Hello Jumble")
    js = jumble_sort("car truck 8 4 bus 6 1")
    print(js)
