public class Grid {
    public static void main(String[] args) {
        System.out.println("2x2 => 2");
        System.out.println(gridCount(2,2));

        System.out.println("2x3 => 3");
        System.out.println(gridCount(2,3));

        System.out.println("3x3 => 6");
        System.out.println(gridCount(3,3));

        System.out.println("4x4 => 20");
        System.out.println(gridCount(4,4));
    }

    public static int gridCount(int n, int m) {
        return gridCountIter(n, m, 0, 0);
    }

    private static int gridCountIter(int n, int m, int x, int y) {
        if(x == n-1 && y == m-1) {
            return 1;
        }

        int count = 0;
        if(x < n-1) {
            count += gridCountIter(n, m, x+1, y);
        }
        if(y < m-1) {
            count += gridCountIter(n, m, x, y+1);
        }

        return count;
    }
}
