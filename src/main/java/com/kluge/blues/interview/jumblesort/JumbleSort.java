/**
 *
 */
package com.kluge.blues.interview.jumblesort;

import java.util.List;
import java.util.Scanner;

/**
 * @author roland
 *
 */
public class JumbleSort {

	/**
	 * Use ctrl+c to terminate
	 *
	 * @param args
	 * @throws Exception
	 */
	public static void main(String args[]) throws Exception {
		Scanner scanner = new Scanner(System.in);
		String line = scanner.nextLine();
		scanner.close();

		InputHandler handler = new InputHandler(line);
		List<String> sorted = handler.sort();
		String formattedResults = formatResults(sorted);

		System.out.println(formattedResults);
	}

	private static String formatResults(List<String> list) {
		StringBuilder sb = new StringBuilder();
		for (String s : list) {
			sb.append(s);
			sb.append(" ");
		}
		return sb.toString().trim();
	}

}
