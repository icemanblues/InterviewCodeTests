function targetSum(a, t) {
    const dict = {};
    for(let i=0; i< a.length; i++) {
        dict[a[i]] = i;

        const target = t - a[i];
        if(target in dict) {
            return [i, dict[target]];
        }
    }

    return [-1, -1];
}

const a = [1, 2, 5, 6, 7, 9, 11];
const t = 9;
const ans = targetSum(a,t);
console.log("(1,4) => ", ans);
