console.log("hi");

function queens(n) {
    console.log("queens with n:", n);
    queens = Array(n);
    console.log("queens array with length:", queens.length);
    isSolved = queensIter(queens, 0);
    console.log("did we solve it?", isSolved)
    return queens;
}

// the array of queens, and the index i we're solving for
// return true if it find a solution, false if no solution could be determined
function queensIter(queens, i) {
    if( i >= queens.length ) {
        return true;
    }

    // try all possible values from 0 to n-1
    for(j=0; j<queens.length; j++) {
        v = isValid(queens, i, j);
        if(!v) {
            continue;
        }

        queens[i] = j;
        isSolved = queensIter(queens, i+1);
        if(isSolved) {
            return true;
        }
    }

    // tried everything and it just didn't work
    return false;
}

function isValid(queens, i, j) {
    // check column
    for(k=0;k<i;k++) {
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

// lets run it
console.log(queens(8));
