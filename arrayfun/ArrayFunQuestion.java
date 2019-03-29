import java.util.List;

/**
 * A template file with basic JUnit tests has been provided.
 * Note that your program will be evaluated by a more comprehensive test suite.
 *
 * Please consider the runtime and space complexity of your algorithms and comment on those trade offs.
 *
 * Good luck!
 */

public class ArrayFunQuestion {

    /***
     *  @param chars an array of characters
     *  @return the same array with the character order reversed
     */
    public static char[] reverse(char[] chars){
        //TODO - implement me!
        return null;
    }

    /***
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
        //TODO - implement me!
        return false;
    }

    /***
     * Find the most used char in the following char[]
     * If a few chars share the top count, return them all.
     *
     * Caveat 1: exclude the space (' ') character from your count.
     * Caveat 2: treat all uppercase characters as their lowercase counterpart
     *
     * @return the characters which appear most frequently
     */
    public static List<Character> mostUsedCharacter(char [] chars){
        //TODO - implement me!
        return null;
    }

    /***
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
    public static int numberOfPalindromes(char[] chars){
        //TODO - implement me!
        return 0;
    }

}
