-- ProjectEuler.net Solutions
-- Problem 7: What is the 10,001st prime number?


-- Function nthPrime
-- Input: Integer n
-- Returns: Integer
-- nthPrime determines the prime number of the given index
-- ex: nthPrime 3 returns 5, because 5 is the 3rd prime number
nthPrime :: Int -> Integer
nthPrime n = (findPrime [2] n)


-- Function findPrime
-- Input: Array of integers, integer
-- Returns: Integer
-- findPrime will generate the list of primes 
-- using the input array as a start, 
-- and then return the prime number at the index specified.
findPrime :: [Integer] -> Int -> Integer
findPrime n 0 = 1
findPrime (x:xs) 1 = x
findPrime (x:xs) m 
	| (length (x:xs)) == m 	= x
	| otherwise 			= findPrime ((nextPrime (x+1) xs):x:xs) m


-- Function nextPrime
-- Input: Integer n, list of integers m
-- Returns: Integer
-- nextPrime will return the next prime number 
-- in the list of primes, after the input integer n
nextPrime :: Integer -> [Integer] -> Integer
nextPrime n m
	| prime n m = n
	| even n 	= nextPrime (n+1) m
	| otherwise = nextPrime (n+2) m


-- Function prime
-- Input: integer n, list of integers xs
-- Returns: Boolean
-- prime takes in an integer value as input,
-- and uses the definition of a prime number
-- to determine if it is prime. If it is prime,
-- it returns true, otherwise it returns false
prime :: Integer -> [Integer] ->  Bool
prime n xs
	| n == 2										= True
	| (even n) 										= False
	| (length ([x | x <- xs, mod n x == 0]) == 0)	= True
	| otherwise 									= False