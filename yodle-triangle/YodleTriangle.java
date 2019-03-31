import java.io.BufferedReader;
import java.io.FileReader;
import java.util.Arrays;

/**
 * @author <a href="mailto:roland.kluge@gmail.com">Roland Kluge</a>
 */
public class YodleTriangle {
	private final int numRows;
	private final String filename;
	private final int[][] triangle;
	
	public YodleTriangle(final int numRows, final String filename) {
		this.numRows = numRows;
		this.filename = filename;
		
		this.triangle = buildTriangle(numRows, filename);
	}
	
	private int[][] buildTriangle(final int numRows, final String filename) {
		final int[][] triangle = new int[numRows][];
		
		int lineNum = 1; // start at 1 because it is the number of elements in the line
		try(BufferedReader br = new BufferedReader(new FileReader(filename));) {
			String line = br.readLine();
			while(line != null) {
				String[] words = line.split("\\s+");
				
				if(words.length != lineNum) {
					throw new IllegalArgumentException("Expecting "+ lineNum + " elements, found " + words.length + " instead");
				}
				
				int[] numbers = Arrays.stream(words).mapToInt(Integer::parseInt).toArray();
				triangle[lineNum-1] = numbers;
				// increment the while loop
				lineNum++;
				line = br.readLine();
			}
		}
		catch(Exception e) {
			System.err.println("Error parsing file at line number: "+lineNum);
			e.printStackTrace();
		}
		
		return triangle;
	}
	
	private int[][] buildMemoGrid() {
		final int[][] memo = new int[triangle.length][];
		for(int i=0; i< triangle.length; i++) {
			memo[i] = new int[triangle[i].length];
		}
		
		return memo;
	}
	
	public void printTriangle() {
		for(int i=0; i<triangle.length; i++) {
			final int expectedLength = i+1;
			System.out.println("line "+expectedLength+": "+Arrays.toString(triangle[i]));
		}
	}
	
	public boolean validateTriangle() {
		for(int i=0; i<triangle.length; i++) {
			final int expectedLength = i+1;
			if(triangle[i].length != expectedLength) {
				return false;
			}
		}
		
		return true;
	}
	
	/**
	 * 
	 * @param triangle the triangle
	 * @param row the previous row used
	 * @param col the previous col used
	 * @param sum the current sum so far
	 * @return
	 */
	private int maxSum(final int row, final int col, final int sum) {
		final int nextRow = row+1;
		
		if(nextRow >= triangle.length) {
			return sum;
		}
		
		// nextCol could be col or col+1
		final int left = maxSum(nextRow, col, sum+triangle[nextRow][col]);
		final int right= maxSum(nextRow, col+1, sum+triangle[nextRow][col+1]);
		return Math.max(left, right);
	}
	
	public int solveDFS() {
		return maxSum(0, 0, triangle[0][0]);
	}
	
	public int solve() {
		final int[][] memo = buildMemoGrid();
		
		// initialize the memo grid with the bottom values
		// then bubble up with the max value
		
		// FIXME: this should be System.ArrayCopy
		final int botRow = triangle.length-1;
		for(int i=0; i< triangle[botRow].length; i++) {
			memo[botRow][i] = triangle[botRow][i];
		}
		
		for(int i=botRow-1; i>=0; i--) {
			for(int j=0; j<triangle[i].length; j++) {
				// this is okay, since the row after i (i+1) will have (j+1)
				final int left = memo[i+1][j];
				final int right = memo[i+1][j+1];
				memo[i][j] = triangle[i][j] + Math.max(left, right);
			}
		}
		
		return memo[0][0];
	}
	
	public static void main(String[] args) {
		final int numRows = Integer.parseInt(args[0]);
		final String filename = args[1];
		final YodleTriangle triangle = new YodleTriangle(numRows, filename);
		
		final boolean isValid = triangle.validateTriangle();
		if(!isValid) {
			throw new IllegalArgumentException("This triangle is not valid");
		}
		
		triangle.printTriangle();
		
		final int solution = triangle.solve();
		System.out.println("memo: "+solution);
		
		final int old = triangle.solveDFS();
		System.out.println("dfs: "+old);
	}
}
