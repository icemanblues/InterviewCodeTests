
def target_sum_sorted(a, t):
    i, j = 0, len(a)-1
    sum = a[i] + a[j]
    while sum != t:
        sum = a[i] + a[j]
        if sum < t:
            i += 1
        elif sum > t:
            j -= 1
    return i, j


def target_sum(a, t):
    d = {}
    for i in range(len(a)):
        d[a[i]] = i
        target = t - a[i]
        if target in d:
            return i, d[target]


if __name__ == '__main__':
    a = [1,2,4,6,7,9,11]
    t = 9

    print('sorted:', target_sum_sorted(a, t))
    print('nosort:', target_sum(a, t))
    
