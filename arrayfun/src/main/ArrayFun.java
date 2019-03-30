import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Set;

/**
 * A template file with basic JUnit tests has been provided.
 * Note that your program will be evaluated by a more comprehensive test suite.
 *
 * Please consider the runtime and space complexity of your algorithms and comment on those trade offs.
 *
 * Good luck!
 */

public class ArrayFun {

    /***
     * This runs in O(n) - only one pass of the array. Technically, only half the array,
     * but that is still O(n)
     * No extra memory is used, aside from the temp variable, which is constant
     *
     *  @param chars an array of characters
     *  @return the same array with the character order reversed
     */
    public static char[] reverse(char[] chars){
    	if(chars != null) {
    		int midpoint = (chars.length -1 ) / 2;
	    	for(int i=0; i<midpoint; i++) {
	    		int index = chars.length-1-i;
	    		char temp = chars[i];
	    		chars[i] = chars[index];
	    		chars[index] = temp;
	    	}
    	}

    	return chars;
    }

    /***
     * This runs in O(n), technically O(n/2) which is O(n)
     * It uses constant space based on the primitives it requires
     *
     *
     *  A palindrome is a word, phrase, number, or other sequence of units that may be
     *  read the same way in either direction.
     *
     *  Caveat 1 : punctuation and spaces do not count as contributing characters
     *    (hint) use the static method Character.isLetter(char c)
     *
     *  Caveat 2 : upper case letters should still match with lower case
     *    (hint) use the static method Character.toLowerCase(char c)
     *
     *  Examples of valid palindromes :
     *
     *  "Never odd or even"
     *  "No trace; not one carton"
     *  "A Toyota! Race fast... safe car: a Toyota"
     *  "Sums are not set as a test on Erasmus"
     *
     *  @return true if chars form a valid palindrome, otherwise false
     */
    public static boolean isPalindrome(char[] chars){
    	boolean result = true;

    	if(chars != null) {
    		int front = 0;
    		int end = chars.length-1;
    		while(front < end) {
    			char firstChar = Character.toLowerCase(chars[front]);
    			if(!Character.isLetter(firstChar)) {
    				front++;
    				continue;
    			}

    			char endChar = Character.toLowerCase(chars[end]);
    			if(!Character.isLetter(endChar)) {
    				end--;
    				continue;
    			}

    			if(firstChar != endChar) {
    				result = false;
    				break;
    			}
    			else {
    				front++;
    				end--;
    			}
    		}
    	}
    	else {
    		result = false;
    	}

    	return result;
    }

    /***
     * This runs in O(n) since it has to check each element of chars
     * It uses space of sice O(n) since we have a map that contains each
     * elem of chars. Its O(n) space since each character can be unique
     *
     * Find the most used char in the following char[]
     * If a few chars share the top count, return them all.
     *
     * Caveat 1: exclude the space (' ') character from your count.
     * Caveat 2: treat all uppercase characters as their lowercase counterpart
     *
     * @return the characters which appear most frequently
     */
    public static List<Character> mostUsedCharacter(char [] chars){
    	if(chars == null) {
    		return null;
    	}

    	final char space = ' ';
    	final List<Character> results = new LinkedList<Character>();
    	final Map<Character, Integer> characterCounter = new HashMap<Character, Integer>();
    	int highestOccurrence = Integer.MIN_VALUE;
    	for(char ch: chars) {
    		char c = Character.toLowerCase(ch);
    		if (c != space) {
				if (!characterCounter.containsKey(c)) {
					characterCounter.put(c, 1);
					if(highestOccurrence < 1) {
						highestOccurrence = 1;
						results.add(ch);
					}
				} else {
					int count = characterCounter.get(c);
					characterCounter.put(c, ++count);
					if(count > highestOccurrence) {
						highestOccurrence = count;
						results.clear();
						results.add(ch);
					}
					else if(count == highestOccurrence) {
						results.add(ch);
					}
				}
			}
    	}

    	return results;
    }

    /***
     * The performance is O(n^3) since we have to do a loop within a loop when we are
     * looking at all of the possible sub words that can be palidromes, and the call to
     * isPalidromes is another loop in those.
     *
     * I assume Set contains is constant, as is System.Arraycopy (this may not be the case)
     *
     * Space is order (n^2) due to the possible number of sub words
     *
     * Find the number of unique palindromes in the char array
     *
     * Example: for the array {'a', 'b', 'b', 'a', 'c', 'c', 'a', 'b' 'b', 'a', 'd', 'd', 'a'}
     *  the palindromes are
     *      abbaccabba
     *      bbaccabb
     *      baccab
     *      acca
     *      cc
     *      abba
     *      bb
     *      adda
     *      dd
     *
     * @return the number of unique palindromes in the character array
     */
	public static int numberOfPalindromes(char[] chars) {
		Set<String> knownPalindromes = new HashSet<String>();

		if (chars != null) {
			int start = 0;
			while (start < chars.length - 1) { // min size is 2 for a palindrome here
				for (int end = start + 1; end < chars.length; end++) {
					int wordLength = end - start + 1;
					char[] word = new char[wordLength];
					System.arraycopy(chars, start, word, 0, wordLength);

					// I don't really need to check if it already in the set, but
					// then it if does, it saves us the O(n) check to isPalidrome
					String s = new String(word);
					if (!knownPalindromes.contains(s) && isPalindrome(word)) {
						knownPalindromes.add(s);
					}
				}
				start++;
			}
		}

		return knownPalindromes.size();
	}
}