public class Fib {

	public static int fib(int x) {
		if (x < 0) {
			System.out.println("ERR: You passed in a non-positive integer: " + x);
			return x;
		}
		if (x == 1 || x == 0) {
			return 1;
		}

		return fib(x - 1) + fib(x - 2);
	}

	public static void main(String[] args) throws Exception {
		System.out.println("Starting fib tests");
		for (int x : new int[] { 1, 2, 3, 4, 5, 6, 7, 8, 10, -5 }) {
			System.out.println("fib x " + x + " : " + fib(x));
		}

		System.out.println("Ending fib tests");
	}
}
