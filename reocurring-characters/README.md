# Reoccurring Characters

## Problem Statement

### Part 1

Give a string, detect the first reoccurring character. 
If there is no reoccurring character, then return null, nil, empty, etc

### Part 2

Do the above without using memory

## Examples

* `"abcdefghiabc"` => `a`
    ```
    abcdefghiabc
    ^        ^
    ```
* `"cc"` => `c`
    ```
    cc
    ^^
    ```
* `"cabbage"` => `b`
    ```
    cabbage
      ^^
    ```
* `""` => not found, nil, empty, null
