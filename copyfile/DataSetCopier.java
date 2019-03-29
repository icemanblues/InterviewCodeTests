/**
 *
 */
package com.kluge.blues.interview.copyfile;

import java.util.Scanner;
import java.util.Set;

/**
 * @author roland
 *
 */
public class DataSetCopier {

	public static void main(String[] args) throws Exception {
		Scanner scanner = new Scanner(System.in);

		String firstLine = scanner.nextLine();
		int numberOfDataCenters = Integer.valueOf(firstLine);
		DataSetManager manager = new DataSetManager(numberOfDataCenters);

		int count = 1;
		while (count < numberOfDataCenters+1) {
			String line = scanner.nextLine();
			manager.addDataCenter(count, line);
			count++;
		}
		scanner.close();

		Set<Copy> copyResults = manager.getCopies();
		for (Copy copy: copyResults) {
			System.out.println(copy);
		}
		System.out.println("done");
	}
}
