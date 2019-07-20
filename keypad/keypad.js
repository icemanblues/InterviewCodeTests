const keypad = {
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

function t9(number) {
    acc = [];
    t9rec(number, 0, [], acc);
    return acc;
}

function t9rec(number, idx, l, acc) {
    if(idx >= number.length) {
        acc.push([...l]);
    } else {
        const c = number.charAt(idx);
        const possibles = keypad[c];
        for(let i=0; i<possibles.length; i++) {
            l.push(possibles[i]);
            t9rec(number, idx+1, l, acc);
            l.pop();
        }
    }
}

console.log("23 => ", t9("23"));
