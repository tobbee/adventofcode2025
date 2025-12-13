import unittest
from part2 import part2


class TestPart2(unittest.TestCase):
    def test_part2_with_testinput(self):
        result = part2("dec10/testinput")
        self.assertIsInstance(result, int)
        self.assertGreater(result, 0)


if __name__ == "__main__":
    unittest.main()
