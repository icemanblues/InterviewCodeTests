function longestPrefix(words) {
    if(words.length === 0) {
        return "";
    } else if(words.length === 1) {
        return words[0];
    }

    // 2 or more
    let prefix = undefined;
    for(let i=0; i<words.length-1; i++) {
        if(prefix) {
            prefix = cmp(prefix, words[i+1]);
        } else {
            prefix = cmp(words[i], words[i+1]);
        }
    }

    return prefix;
}

function cmp(a, b) {
    let i=0;
    while(i<a.length && i<b.length && a[i] === b[i]) {
        i++;
    }
    return a.substring(0,i);
}

a = ['hello'];
b = [`floor`, `florida`, `fluid`];
c = [`doodle`, `double`, `dave`, `murmur`];
console.log("hello => ",longestPrefix(a));
console.log("fl => ",longestPrefix(b));
console.log("'' => ",longestPrefix(c));