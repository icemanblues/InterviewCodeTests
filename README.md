# InterviewCodeTests

These are code tests that I've been asked over the years. Each question has a time limit of 1 hour

##ArrayFun
The questions can be found in `com.kluge.blues.interview.arrayfun.ArrayFunQuestion`

Some fun and interesting palidrome questions related to character arrays (Strings)

The interviewer provided JUnit tests to help the interviewee confirm their answers.

## Connected Cities

Suppose you are given a text file containing pairs of city names, one pair per line, with the names on each line separated by a comma. The file might look something like:
 
```
Philadelphia, Pittsburgh
Boston, New York
Philadelphia, New York
Los Angeles, San Diego
New York, Croton-Harmon
St. Petersburg, Tampa
```
 
Each line of the file indicates that it is possible to travel between the two cities named. (More formally, if we think of the cities as nodes in a graph, each line of the file specifies an edge between two nodes.) In the example above it is possible to travel between Boston and New York, and also between New York and Philadelphia and between Philadelphia and Pittsburgh, so it follows that Boston and Pittsburgh are connected. On the other hand, there is no path between Boston and Tampa, so they are not connected.
 
Write a Java program, called "Connected" that reads in such a file and outputs whether two specified cities are connected.
 
The program will be invoked from the command line as:
`java Connected <filename> <cityname1> <cityname2>`
 
and will output a single line stating "yes" or "no"
 
Here are some sample interactions, assuming the example file above is named "cities.txt"

```
> java Connected cities.txt "New York" Boston
yes
> java Connected cities.txt Boston Pittsburgh
yes
> java Connected cities.txt Boston Tampa
no
> java Connected cities.txt Boston Ypsilanti
no
```

*Notes*
 
* Commas will not appear within city names in the file. For example, "Washington, D.C." will not appear in the file as a city name.
* You can depend on command-line quoting for cities with spaces in their names, like "New York". This should, thus, require no extra code on your part.
* Your choice of algorithms and data structures should allow the program to handle arbitrarily large files reasonably efficiently. You can, however, assume that the file will fit in memory.
* The program is permitted to return any or no output when given a malformed input file or malformed command line.
If a city is not in the file, then it is connected to no other city.

## JumbleSort
You are to write a program that takes a list of strings containing
integers and words and returns a sorted version of the list.
 
The goal is to sort this list in such a way that all words are in
alphabetical order and all integers are in numerical order.
Furthermore, if the nth element in the list is an integer it must
remain an integer, and if it is a word it must remain a word.
 
*Input:*
 
The input will contain a single, possibly empty, line containing a
space-separated list of strings to be sorted. Words will not contain
spaces, will contain only the lower-case letters a-z. Integers will be
in the range -999999 to 999999, inclusive. The line will be at most 1000
characters long.
 
*Output:*
 
The program must output the list of strings, sorted per the requirements
above. Strings must be separated by a single space, with no leading
space at the beginning of the line or trailing space at the end of the
line.
 
 
*Constraints:*
 
The code you submit must take input from stdin and produce output to
stdout as specified above. No other output is permitted. You can
assume the input will be valid. Feel free to use standard libraries to
assist in sorting.
 
In the examples below, the text "Input:" and "Output:" are not part
of the output, and neither are the blank lines.
 
 
### Example 1:

*Input:*

`1`
 
*Output:*

`1`
 
 
### Example 2:

*Input:*

`car truck bus`
 
*Output:*

`bus car truck`
 
### Example 3:

*Input:*

`8 4 6 1 -2 9 5`
 
*Output:*
 
`-2 1 4 5 6 8 9`
 
 
### Example 4:

*Input:*
 
`car truck 8 4 bus 6 1`
 
*Output:*
 
`bus car 1 4 truck 6 8`

##CopyFile

Oh no! Disaster has struck some of ACME's redundant data centers. The
administrators have managed to restore backups, but some data sets are
still missing from some data centers. Fortunately, every data set can be
found at least once in one or more of the data centers. However, before
ACME can resume normal operations, it needs to ensure that each data
center has a copy of every data set.
 
Your goal is to help ACME resume normal operations by writing a program
to synchronize data sets between data centers using as few copies as
possible.
 
*Input:*
 
The first line of input will contain an integer between 0 and 999999
inclusive, representing the number of data centers.
 
Following that will be one line of input for each data center with a
space-separated list of data set ids currently present at that data
center. Data set ids are each an integer between 1 and 999999, inclusive.
Each line will be at most 1000 characters long.
 
A data center with no data sets is represented with a blank line. Data
set ids are not necessarily consecutive. The list of data sets will not
be in any particular order.
 
 
*Output:*
 
The program must output an optimal data set copy strategy to ensure that
every data center has a copy of every data set. Output one line for every
copy instruction.
 
A copy instruction is of the form `<data-set-id> <from> <to>`, where
`<data-set-id>` is the data set id, `<from>` is the index of the data center
the data set will be copied from (1 = the first data center), and `<to>`
is the index of the data center to copy the data set to.
 
When there are no more copy instructions, the program must output the
word "done" on a line by itself.
 
There is often more than one correct output with minimal number of
operations for a given input, and any output that satisfies the
requirements is valid.
 
 
*Constraints:*
 
The code you submit must take input from stdin and produce output to
stdout as specified above. No other output is permitted. You can
assume the input will be valid. In the examples below, the text
"Input:" and "Output:" (or "One Possible Correct Output:") are not
part of the output, and neither are the blank lines.
 
### Example 1:

*Input:*

``` 
4
1 3 4
1 2 3
1 3
1 4 2
```
 
*One Possible Correct Output:*
 
 ```
2 2 1
4 1 2
2 2 3
4 4 3
3 1 4
done
 ```
 
### Example 2:

*Input:*
 
 ```
2
1 2
2 1
```
 
*Output:*
 `
done
 `
 
### Example 3:

*Input:*
 
 ```
3
1 3 4 5 7
1 3
2
 ```
 
*One Possible Correct Output:*
 
 ```
2 3 2
2 3 1
1 1 3
4 1 2
5 1 3
5 3 2
4 2 3
3 1 3
7 1 2
7 1 3
done
```
