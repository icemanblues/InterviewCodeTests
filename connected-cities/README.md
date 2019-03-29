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