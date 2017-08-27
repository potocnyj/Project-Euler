-- ProjectEuler.net Solutions
-- Problem 2: By considering the terms in 
-- the Fibonacci sequence whose values do not exceed 
-- four million, find the sum of the even-valued terms.

-- Function: evenFibSum
-- Input: Integer
-- Returns: Integer
-- evenFibSum will compute and return the sum of the 
-- even fibonacci numbers, up to the input limit specified
evenFibSum :: Integer -> Integer
evenFibSum x = sum (filter even (limFib x))


-- Function: limFib
-- Input: Integer
-- Returns: List of Integers
-- limFib will run fibonacci, and limit to the input integer
limFib :: Integer -> [Integer]
limFib x = fibonacci 1 1 x


-- Function: fibonacci
-- Input: 3 integers, x n m
-- Returns: List of Integers
-- fibonacci will return the list of fibonacci numbers, 
-- up to the value of m.
fibonacci :: Integer -> Integer -> Integer -> [Integer]
fibonacci x n m 
	| x < m 	= x : fibonacci n (x+n) m
	| otherwise = [ ]