import unittest

def find_consecutive_runs(input_array):
    result = []
    
    for i in range(len(input_array)-2):
        diff_1 = input_array[i] - input_array[i+1]
        diff_2 = input_array[i+1] - input_array[i+2]

        if diff_1 == diff_2 and abs(diff_1) == 1:
            result.append(i)
            
    return result


class TestThis(unittest.TestCase):
    def test_simple(self):
        simple = [0, 1, 2]
        self.assertEqual(find_consecutive_runs(simple), [0])


    def test_given(self):
        given = [1, 2, 3, 5, 10, 9, 8, 9, 10, 11, 7]
        self.assertEqual(find_consecutive_runs(given), [0, 4, 6, 7])

    def test_none(self):
        none = []
        self.assertEqual(find_consecutive_runs(none), [])


unittest.main()
