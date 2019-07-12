#!/usr/bin/env python3

import unittest

def find_consecutive_runs(input_array):
    result = []
    
    idx = 0
    direction = None
    count = 1
    for i in range(1, len(input_array)):
        if count == 3:
            result.append(idx)
            direction = None
            idx = i
            count = 1
            continue

        diff = input_array[i] - input_array[i-1]
        if direction == None:
            idx=i-1
            if diff == 1 or diff == -1:
                direction = diff
                count += 1
        elif diff == direction:
            count += 1
        elif diff != direction:
            if count == 3:
                result.append(idx)
            direction = None
            idx = i
            count = 1
    
    if count == 3:
        result.append(idx)
    
    return result

class TestThis(unittest.TestCase):
    def test_simple(self):
        simple = [0, 1, 2]
        self.assertEqual(find_consecutive_runs(simple), [0])


    def test_given(self):
        given = [1, 2, 3, 5, 10, 9, 8, 9, 10, 11, 7]
        self.assertEqual(find_consecutive_runs(given), [0, 4, 6, 7])

unittest.main()
