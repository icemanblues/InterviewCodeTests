import unittest
from jumble import jumble_sort

class TestJumbleSort(unittest.TestCase):

    def test_example_1(self):
        self.assertEqual(jumble_sort("1"), ["1"])
    
    def test_example_2(self):
        self.assertEqual(jumble_sort("car truck bus"), ["bus", "car", "truck"])
    
    def test_example_3(self):
        self.assertEqual(jumble_sort("8 4 6 1 -2 9 5"), ["-2", "1", "4", "5", "6", "8", "9"])

    def test_example_4(self):
        self.assertEqual(jumble_sort("car truck 8 4 bus 6 1"), ["bus", "car", "1", "4", "truck", "6", "8"])

if __name__ == "__main__":
    unittest.main()
