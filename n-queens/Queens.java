import java.util.Arrays;

public class Queens {

    private static boolean queens(int[] q, int i) {
        if(i >= q.length) {
            return true;
        }

        // the values (rows) to try
        for(int j=0; j < q.length; j++) {
            // if not valid, keep trying
            boolean valid = isValid(q, i, j);
            if(!valid) {
                continue;
            }

            q[i] = j;

            // recursively solve for the next position
            // if good, then we solved it, so return true
            // if not, keep iterating the for loop to try to solve it
            boolean ok = queens(q, i+1);    
            if(ok) {
                return true;
            }
        }

        // tried everything and nothing worked :(
        return false;
    }

    private static boolean isValid(int[] q, int i, int j) {
        for(int k=0; k<i; k++) {
            // check rows
            if(j == q[k]) {
                return false;
            }

            // check diagonals
            if(Math.abs(i-k) == Math.abs(j-q[k])) {
                return false;
            }
        }

        return true;
    }

    public static void main(String[] args) throws Exception {
        System.out.println("Attempting to solve 8 queens");
        int[] q = new int[8];
        System.out.println("starting board: "+Arrays.toString(q));

        boolean ok = queens(q, 0);
        if(ok) {
            System.out.println("we solved it! "+Arrays.toString(q));
        } else {
            System.out.println("We were unable to solve it :(");
        }
    }
}