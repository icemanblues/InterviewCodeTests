def longest_prefix(a):
    sum = None
    for i in range(len(a)-1):
        if not sum:
            sum = cmp(a[i], a[i+1])
        else:
            sum = cmp(sum, a[i+1])
    return sum


def cmp(a,b):
    i = 0
    while i<len(a) and i < len(b) and a[i] == b[i]:
        i += 1
    return a[:i]


if __name__ == '__main__':
    a = ['flight', 'flood', 'flute']
    b = ['dood', 'double', 'dave', 'meow']
    print("fl =>", longest_prefix(a))
    print("'' =>", longest_prefix(b))
