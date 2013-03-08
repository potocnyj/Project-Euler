-- ProjectEuler.net Solutions
-- Problem 6: Find the difference 
-- between the sum of the squares 
-- of the first one hundred natural numbers 
-- and the square of the sum.

-- Function: diffOfSums
-- Input: Integer n
-- Returns: Integer
-- Computes the difference between the sum of the squares of the first n natural numbers,
-- and the square of the sum of the first n natural numbers.
diffOfSums :: Integer -> Integer

diffOfSums n = (squareOfSum n) - (sumOfSquares n)


-- Function: squareOfSum
-- Input: Integer x
-- Returns: Integer
-- squareOfSum adds the first x natural numbers,
-- and computes the square of the resulting sum.
squareOfSum :: Integer -> Integer

squareOfSum x = (foldr (+) 0 [1,2 .. x])^2


-- Function: sumOfSquares
-- Input: Integer y
-- Returns: Integer
-- sumOfSquares adds the square 
-- of the first y natural numbers.
sumOfSquares :: Integer -> Integer

sumOfSquares y = foldr (+) 0 [n^2 | n <- [1,2 .. y]]