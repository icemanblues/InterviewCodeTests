import java.io.BufferedReader;
import java.io.FileReader;
import java.util.Arrays;
import java.util.List;
import java.util.ArrayList;
import java.util.stream.Collectors;

/**
 * @author <a href="mailto:roland.kluge@gmail.com">Roland Kluge</a>
 */
public class YodleTriangle {
	final List<List<Integer>> triangle = new ArrayList<>();

	public List<List<Integer>> buildTriangle(final String filename) {
		try(BufferedReader br = new BufferedReader(new FileReader(filename));) {
			String line = br.readLine();
			while(line != null) {
				String[] words = line.split("\\s+");
				
				List<Integer> row = Arrays.stream(words)
					.map(Integer::valueOf)
					.collect(Collectors.toList());
				triangle.add(row);

				// increment the while loop
				line = br.readLine();
			}
		}
		catch(Exception e) {
			System.err.println("Error parsing file: "+filename);
			e.printStackTrace();
		}
		
		return triangle;
	}
	
	public void printTriangle() {
		for(int i=0; i<triangle.size(); i++) {
			final int expectedLength = i+1;
			System.out.println("line "+expectedLength+": "+triangle.get(i));
		}
	}
	
	public boolean validateTriangle() {
		for(int i=0; i<triangle.size(); i++) {
			final int expectedLength = i+1;
			if(triangle.get(i).size() != expectedLength) {
				return false;
			}
		}
		
		return true;
	}

	private int maxSum(final int row, final int col, final int sum) {
		final int nextRow = row+1;
		
		if(nextRow >= triangle.size()) {
			return sum;
		}
		
		// nextCol could be col or col+1
		final int left = maxSum(nextRow, col, sum+triangle.get(nextRow).get(col));
		final int right= maxSum(nextRow, col+1, sum+triangle.get(nextRow).get(col+1));
		return Math.max(left, right);
	}
	
	public int topDown() {
		return maxSum(0, 0, triangle.get(0).get(0));
	}
	
	private int[][] buildMemoGrid() {
		final int[][] memo = new int[triangle.size()][];
		for(int i=0; i< triangle.size(); i++) {
			memo[i] = new int[triangle.get(i).size()];
		}
		
		return memo;
	}

	public int bottomUp() {
		// initialize the memo grid with the bottom values
		// then bubble up with the max value
		final int[][] memo = buildMemoGrid();
		
		final int botRow = triangle.size()-1;
		for(int i=0; i< triangle.get(botRow).size(); i++) {
			memo[botRow][i] = triangle.get(botRow).get(i);
		}
		
		for(int i=botRow-1; i>=0; i--) {
			for(int j=0; j<triangle.get(i).size(); j++) {
				// this is okay, since the row after i (i+1) will have (j+1)
				final int left = memo[i+1][j];
				final int right = memo[i+1][j+1];
				memo[i][j] = triangle.get(i).get(j) + Math.max(left, right);
			}
		}
		
		return memo[0][0];
	}
	
	public static void main(String[] args) {
		final YodleTriangle sample = new YodleTriangle();
		sample.buildTriangle("sample.txt");
		
		if(! sample.validateTriangle()) {
			throw new IllegalArgumentException("This triangle is not valid");
		}
		sample.printTriangle();
		
		System.out.println("sample bottomUp: "+sample.bottomUp());
		System.out.println("sample  topDown: "+sample.topDown());

		final YodleTriangle hard = new YodleTriangle();
		hard.buildTriangle("triangle.txt");
		System.out.println("hard bottomUp: "+hard.bottomUp());
	}
}
