// returns a solution to the placing n queens on a n x n board so that
// no queen is in direct threat of another
function queens(n) {
    const queens = Array(n);
    queensIter(queens, 0);
    return queens;
}

// the array of queens, and the index i we're solving for
// return true if it find a solution, false if no solution could be determined
function queensIter(queens, i) {
    if(i >= queens.length) {
        return true;
    }

    // try all possible values from 0 to n-1
    for(let j=0; j<queens.length; j++) {
        const v = isValid(queens, i, j);
        if(!v) {
            continue;
        }

        queens[i] = j;
        const isSolved = queensIter(queens, i+1);
        if(isSolved) {
            return true;
        }
    }

    // tried everything and it just didn't work
    return false;
}

function isValid(queens, i, j) {
    // check column
    for(let k=0;k<i;k++) {
        // check row
        if(queens[k] === j) {
            return false;
        }

        // check diagonal
        if(Math.abs(k-i) === Math.abs(queens[k] - j)) {
            return false;
        }
    }
    
    return true;
}

// main
console.log(queens(8));
