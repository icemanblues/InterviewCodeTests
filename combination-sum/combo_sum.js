function comboSum(a, t) {
		acc = [];
		comboSumRec(a, 0, t, [], acc);
		return acc;
}

function comboSumRec(a, idx, t, l, acc) {
		if(idx >= a.length) {
				return;
		}

		for(let i=idx; i<a.length; i++) {
				const target = t-a[i];

				// this doesn't work, move on
				if(target < 0) {
						continue;
				} else if (target == 0) {
						// found it
						l.push(a[i]);
						acc.push([...l]);
						l.pop();
				} else {
						// getting closer, continue to look for it
						l.push(a[i]);
						comboSumRec(a, i, target, l, acc);
						l.pop();
				}
		}
}

const a1 = [2, 3, 6, 7];
const t1 = 7;
const test1 = comboSum(a1, t1);
console.log("Test1:",test1);

const a2 = [2, 3, 5];
const t2 = 8;
const test2 = comboSum(a2, t2);
console.log("Test2:",test2);

const a3 = [7, 3, 2];
const t3 = 18;
const test3 = comboSum(a3, t3);
console.log("Test3:",test3);
