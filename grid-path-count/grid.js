function gridCount(n, m) {
    return gridCountIter(n, m, 0, 0);
}

function gridCountIter(n, m, x, y) {
    if(x == n-1 && y == m-1) {
        return 1;
    }

    let count = 0;
    if(x < n-1) {
        count += gridCountIter(n, m, x+1, y)
    }
    if(y < m-1) {
        count += gridCountIter(n, m, x, y+1)
    }

    return count;
}

console.log("2x2 => 2")
console.log(gridCount(2,2))

console.log("2x3 => 3")
console.log(gridCount(2,3))

console.log("3x3 => 6")
console.log(gridCount(3,3))

console.log("4x4 => 20:")
console.log(gridCount(4,4))
