package com.kluge.blues.interview.arrayfun;

import org.junit.Assert;
import org.junit.Test;

import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class ArrayFunTest {

    @Test
    public void testReverse() throws Exception {
        final char[] chars = "Novus".toCharArray();
        final char [] reversed = "suvoN".toCharArray();
        Assert.assertArrayEquals(reversed, ArrayFun.reverse(chars));
    }

    @Test
    public void testIsPalindrome() throws Exception {
        final char[] palindrome = "A Toyota! Race fast... safe car: a Toyota".toCharArray();
        final char[] notPalindrome = {'N','o','v','u','s'};
        Assert.assertTrue(ArrayFun.isPalindrome(palindrome));
        Assert.assertFalse(ArrayFun.isPalindrome(notPalindrome));
    }

    @Test
    public void testMostUsedCharacter() throws Exception{
        final char[] chars = "A Toyota! Race fast... safe car: a Toyota".toCharArray();
        final List<Character> output = ArrayFun.mostUsedCharacter(chars);
        Assert.assertEquals(output.size(), 1);
        Assert.assertEquals(output.get(0), new Character('a'));

        final char[] chars2 = "Abcba".toCharArray();
        final List<Character> output2 = ArrayFun.mostUsedCharacter(chars2);
        final Set<Character> resultSet = new HashSet<Character>(output2);
        Assert.assertEquals(output2.size(), 2);
        Assert.assertTrue(resultSet.contains('a'));
        Assert.assertTrue(resultSet.contains('b'));
        Assert.assertFalse(resultSet.contains('c'));
    }

    @Test
    public void testNumberOfPalindromes() throws Exception {
        char[] chars = "abbaccabbadda".toCharArray();
        int numPalindromes = ArrayFun.numberOfPalindromes(chars);
        Assert.assertEquals(numPalindromes, 9);
    }
}
