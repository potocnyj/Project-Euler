-- ProjectEuler.net Solutions
-- Problem 4: Find the largest palindrome
-- made from the product of two 3-digit numbers.


-- Function: isPalindrome
-- Input: List of Integers
-- Returns: Bool
-- isPalindrome takes a list of digits that make up a number, 
-- and checks to see if the number's digits make it a palindrome
isPalindrome :: [Integer] -> Bool
isPalindrome (x:xs)
	| length xs == 0 		= True
	| (length xs) == 1 		= (x == head xs)
	| (x == last xs) 		= (isPalindrome (init xs))
	| otherwise 			= False

-- Function: digits
-- Input: Integer
-- Returns: List of Integers
-- digits will break an Integer into an ordered list of its digits
digits :: Integer -> [Integer]
digits n = map (`mod` 10) $ reverse $ takeWhile (> 0) $ iterate (`div` 10) n