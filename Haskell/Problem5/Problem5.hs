-- ProjectEuler.net Solutions
-- Problem 5: What is the smallest 
-- positive number that is evenly divisible 
-- by all of the numbers from 1 to 20?

-- Function: smallestFactorable
-- Inputs: Integer n
-- Returns: Integer
-- smallestFactorable finds the smallest 
-- integer that is evenly divisible by every number from 1 to n.
smallestFactorable :: Integer -> Integer

smallestFactorable n = findSmallestFactorable n [1,2 .. n]

-- Function: findSmallestFactorable
-- Inputs: Integer x, list of integers xs
-- Returns: Integer
-- findSmallestFactorable recursively checks 
-- for whether the input integer x is evenly divisible 
-- by all of the integers in the input list xs. 
-- If it is not, then the largest value of xs is 
-- added to x and the function is called again. 
findSmallestFactorable :: Integer -> [Integer] -> Integer

findSmallestFactorable x xs
	| [y | y <- xs, mod x y == 0] == xs	= x
	| otherwise							= findSmallestFactorable (x + last xs) xs