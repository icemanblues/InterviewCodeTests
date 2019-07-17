keypad = {
    '0': ['0'],
    '1': ['1'],
    '2': ['a','b','c'],
    '3': ['d','e','f'],
    '4': ['g','h','i'],
    '5': ['j','k','l'],
    '6': ['m','n','o'],
    '7': ['p','q','r','s'],
    '8': ['t','u','v'],
    '9': ['w','x','y','z'],
}

def t9(number):
    return t9_helper(number, 0, [], [])

def t9_helper(number, idx, l, acc):
    if idx >= len(number):
        acc.append([x for x in l])
    else:
        possibles = keypad[number[idx]]
        for p in possibles:
            l.append(p)
            t9_helper(number, idx+1, l, acc)
            l.pop()
    return acc


if __name__ == '__main__':
    print('23 =>', t9('23'))
