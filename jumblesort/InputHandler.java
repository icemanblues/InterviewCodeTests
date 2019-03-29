/**
 *
 */
package com.kluge.blues.interview.jumblesort;

import java.util.Collections;
import java.util.LinkedList;
import java.util.List;

/**
 * @author roland
 *
 */
public class InputHandler {
	private final String line;
	private boolean[] isString;

	private LinkedList<String> wordsToSort;
	private LinkedList<Integer> numbersToSort;

	public InputHandler(String line) {
		this.line = line;
		wordsToSort = new LinkedList<String>();
		numbersToSort = new LinkedList<Integer>();
	}

	private void initialize() {
		String[] tokens = line.split("\\s");
		isString = new boolean[tokens.length];

		for(int i=0; i<tokens.length; i++) {
			String s = tokens[i];
			boolean isAWord = isWord(s);
			isString[i] = isAWord;

			if(isAWord) {
				wordsToSort.add(s);
			} else {
				int number = Integer.valueOf(s);
				numbersToSort.add(number);
			}
		}
	}

	public List<String> sort() {
		List<String> result;
		if(line == null || line.trim().equals("")) {
			 result = new LinkedList<String>();
		} else {
			initialize();

			sortNumbers();
			sortWords();

			result = merge();
		}

		return result;
	}

	private boolean isWord(String s) {
		return s.matches("[a-z]+");
	}

	private void sortNumbers() {
		 Collections.sort(numbersToSort);
	}

	private void sortWords() {
		Collections.sort(wordsToSort);
	}

	private List<String> merge() {
		List<String> sortedResults = new LinkedList<String>();

		for(boolean isAWord: isString) {
			if(isAWord) {
				String word = wordsToSort.poll();
				sortedResults.add(word);
			} else { // must be an integer
				int number = numbersToSort.poll();
				String s = String.valueOf(number);
				sortedResults.add(s);
			}
		}

		return sortedResults;
	}
}
