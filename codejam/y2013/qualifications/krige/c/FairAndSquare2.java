package codejam.y2013.qualifications.krige.c;

import java.io.BufferedInputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileWriter;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.Scanner;

import org.apache.commons.lang3.time.StopWatch;

/**
 * Problem
 * 
 * Little John likes palindromes, and thinks them to be fair (which is a fancy
 * word for nice). A palindrome is just an integer that reads the same backwards
 * and forwards - so 6, 11 and 121 are all palindromes, while 10, 12, 223 and
 * 2244 are not (even though 010=10, we don't consider leading zeroes when
 * determining whether a number is a palindrome).
 * 
 * He recently became interested in squares as well, and formed the definition
 * of a fair and square number - it is a number that is a palindrome and the
 * square of a palindrome at the same time. For instance, 1, 9 and 121 are fair
 * and square (being palindromes and squares, respectively, of 1, 3 and 11),
 * while 16, 22 and 676 are not fair and square: 16 is not a palindrome, 22 is
 * not a square, and while 676 is a palindrome and a square number, it is the
 * square of 26, which is not a palindrome.
 * 
 * Now he wants to search for bigger fair and square numbers. Your task is,
 * given an interval Little John is searching through, to tell him how many fair
 * and square numbers are there in the interval, so he knows when he has found
 * them all.
 * 
 * Solving this problem
 * 
 * Usually, Google Code Jam problems have 1 Small input and 1 Large input. This
 * problem has 1 Small input and 2 Large inputs. Once you have solved the Small
 * input, you will be able to download any of the two Large inputs. As usual,
 * you will be able to retry the Small input (with a time penalty), while you
 * will get only one chance at each of the Large inputs.
 * 
 * Input
 * 
 * The first line of the input gives the number of test cases, T. T lines
 * follow. Each line contains two integers, A and B - the endpoints of the
 * interval Little John is looking at.
 * 
 * Output
 * 
 * For each test case, output one line containing "Case #x: y", where x is the
 * case number (starting from 1) and y is the number of fair and square numbers
 * greater than or equal to A and smaller than or equal to B.
 * 
 * Limits
 * 
 * Small dataset
 * 
 * 1 ≤ T ≤ 100. 1 ≤ A ≤ B ≤ 1000.
 * 
 * First large dataset
 * 
 * 1 ≤ T ≤ 10000. 1 ≤ A ≤ B ≤ 1014.
 * 
 * Second large dataset
 * 
 * 1 ≤ T ≤ 1000. 1 ≤ A ≤ B ≤ 10100.
 * 
 * @author Marco Lackovic <marco.lackovic@gmail.com>
 * @version 2.0, April 21, 2013
 */
public class FairAndSquare2 {
    private static final String FILE_NAME = "C-large-2";

    private static final String INPUT_FILE_NAME = FILE_NAME + ".in";
    private static final String OUTPUT_FILE_NAME = FILE_NAME + ".out";

    private static final String ROOT_DIR = "dat";
    private static final String YEAR = "2013";
    private static final String ROUND = "qualifications";

    private static final String PATH = ROOT_DIR + File.separator + YEAR
            + File.separator + ROUND + File.separator;

    public static void main(String[] args) throws IOException {
        Scanner input = new Scanner(new BufferedInputStream(
                new FileInputStream(PATH + INPUT_FILE_NAME)));
        PrintWriter output = new PrintWriter(new FileWriter(PATH
                + OUTPUT_FILE_NAME));

        int numTestCases = input.nextInt();
        StopWatch t = new StopWatch();
        t.start();
        for (int tc = 1; tc <= numTestCases; tc++) {
            float a = input.nextFloat();
            float b = input.nextFloat();
            String outputLine = "Case #" + tc + ": " + solve(a, b);
            System.out.println(outputLine);
            output.println(outputLine);
        }
        t.stop();
        System.out.println(t);
        input.close();
        output.close();
    }

    private static long solve(float a, float b) {
        long num = 0;
        a = (float) Math.ceil(Math.sqrt(a));
        b = (float) Math.floor(Math.sqrt(b));

        for (long i = (long) a; i <= b; i++) {
            if (isPalindromeFast(i) && isPalindromeFast(i * i)) {
                num++;
            }
        }
        return num;
    }

    @SuppressWarnings("unused")
    private static boolean isPalindromeSlow(long n) {
        String number = Long.toString(n);
        int l = number.length();
        for (int i = 0; i < l / 2; i++) {
            if (number.charAt(i) != number.charAt(l - i - 1)) {
                return false;
            }
        }
        return true;
    }

    private static boolean isPalindromeFast(long n) {
        long original = n;
        long reversed = 0;
        for (int i = 0; i <= n; i++) {
            reversed = reversed * 10 + n % 10;
            n = n / 10;
        }
        return reversed == original;
    }

}
