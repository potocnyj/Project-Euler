-- ProjectEuler.net Solutions
-- Problem 1: Find the sum of all the 
-- multiples of 3 or 5 below 1000.

-- Function sumOfMultiples
-- Inputs: Integer
-- Returns: Integer
-- sumOfMultiples will return the sum of 
-- the multiples of 3 and 5, below the input limit value
sumOfMultiples :: Integer -> Integer
sumOfMultiples x = sum (multOfThreeAndFive x)

-- Function: multOfThreeAndFive
-- Inputs: Integer
-- Returns: List of Integers
-- multOfThreeAndFive will return a list of all of 
-- the numbers below the input limit value that are 
-- a multiple of three or 5.
multOfThreeAndFive :: Integer -> [Integer]
multOfThreeAndFive x = [n | n <- [1 .. x-1], mod n 3 == 0 || mod n 5 == 0]