import unittest
from reoccur import reoccur, reoccur_no_mem

class TestReoccur(unittest.TestCase):
    def test_reoccur_1(self):
        self.assertEqual(reoccur("abcdefabcdef"), 'a')
    
    def test_reoccur_2(self):
        self.assertEqual(reoccur("cc"), 'c')
    
    def test_reoccur_3(self):
        self.assertEqual(reoccur("xylem"), False)
    
    def test_reoccur_4(self):
        self.assertEqual(reoccur("cabbage"), 'b')
    
    def test_reoccur_5(self):
        self.assertEqual(reoccur(""), False)
    
    def test_reoccur_6(self):
        self.assertEqual(reoccur("12345 098765 blah"), '5')
    
    def test_reoccur_no_mem_1(self):
        self.assertEqual(reoccur_no_mem("abcdefabcdef"), 'a')
    
    def test_reoccur_no_mem_2(self):
        self.assertEqual(reoccur_no_mem("cc"), 'c')
    
    def test_reoccur_no_mem_3(self):
        self.assertEqual(reoccur_no_mem("xylem"), False)
    
    def test_reoccur_no_mem_4(self):
        self.assertEqual(reoccur_no_mem("cabbage"), 'b')
    
    def test_reoccur_no_mem_5(self):
        self.assertEqual(reoccur_no_mem(""), False)
    
    def test_reoccur_no_mem_6(self):
        self.assertEqual(reoccur_no_mem("12345 098765 blah"), '5')

if __name__ == "__main__":
    unittest.main()
