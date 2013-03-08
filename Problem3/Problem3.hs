-- ProjectEuler.net Solutions
-- Problem 3: What is the largest 
-- prime factor of the number 600851475143?

-- Function: largestPrimeFactor
-- Inputs: Integer n
-- Returns: Integer
-- largestPrimeFactor finds the largest 
-- prime factor of the input integer.
largestPrimeFactor :: Integer -> Integer

largestPrimeFactor n = foldr max 0 (primeFactors n)


-- Function: primeFactors
-- Inputs: Integer x
-- Returns: list of Integers
-- primeFactors builds a list of the prime factors 
-- of the input integer x.
primeFactors :: Integer -> [Integer]

primeFactors x = filter prime (divisors x)


-- Function prime
-- Input: integer n
-- Returns: Boolean
-- prime takes in an integer value as input,
-- and uses the definition of a prime number
-- to determine if it is prime. If it is prime,
-- it returns true, otherwise it returns false
prime :: Integer -> Bool

prime n
	| (divisors n) == [n,1]	= True
	| otherwise 			= False


-- Function divisors
-- Input: integer n
-- Returns: List of integers
-- divisors looks at an input number, and 
-- checks every number from 1 to the input number value to 
-- determine if it is a divisor of the input. All divisors of the 
-- input are returned in a list
divisors :: Integer -> [Integer]

divisors n
	| even n 	= n : [x | x <- [1, 2 .. (div n 2)], mod n x == 0]
	| otherwise = n : [x | x <- [1, 3 .. (div n 2)], mod n x == 0]