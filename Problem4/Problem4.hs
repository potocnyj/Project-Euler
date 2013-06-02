-- ProjectEuler.net Solutions
-- Problem 4: Find the largest palindrome
-- made from the product of two 3-digit numbers.

-- Function: largestPalindrome
-- Input: Two lists of integers
-- Returns: Integer
-- Given two lists of integers, the largestPalindrome function 
-- will compute all of the possible palindrome numbers that can be produced 
-- by multiplying two numbers from each list, and returns the largest palindrome that can be made.
largestPalindrome :: [Integer] -> [Integer] -> Integer
largestPalindrome xs ns = foldr max 0 (findPalindromes xs ns)


-- Function: findPalindromes
-- Input: Two lists of integers
-- Returns: List of integers
-- findPalindromes will iterate through a list of palindromes, 
-- multiplying every entry in the list against every entry in the 
-- other list, checking whether it is a palindrome
findPalindromes :: [Integer] -> [Integer] -> [Integer]
findPalindromes (x:xs) ns
	| length xs == 0 			= (findPalindromes2 x ns)
	| otherwise					= (findPalindromes2 x ns) ++ (findPalindromes xs ns)


-- Function: findPalindromes2
-- Input: Two lists of integers
-- Returns: List of integers
-- findPalindromes2 will iterate through a list of palindromes, 
-- multiplying every entry in the list against every entry in the 
-- other list, checking whether it is a palindrome
findPalindromes2 :: Integer -> [Integer] -> [Integer]
findPalindromes2 x (n:ns) 
	| (length ns == 0) && (isPalindrome (digits (x*n))) == False	= [ ]
	| (length ns == 0) && (isPalindrome (digits (x*n)))				= [(x*n)]
	| isPalindrome (digits (x*n))									= (x*n):(findPalindromes2 x ns)
	| otherwise 													= findPalindromes2 x ns


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