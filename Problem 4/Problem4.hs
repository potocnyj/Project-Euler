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
	| n == 1				= False
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
divisors n = n : [x | x <- [1, 2 .. (root n)], mod n x == 0]


-- Function root
-- Input: Integer [Integral] n
-- Returns: Integral of same type input
-- root will calculate the square root of the input number, 
-- and then return it as an Integral type
root :: Integral n => n -> n
root =  floor . sqrt . fromIntegral

-- A sieve filter for primes
sieve :: [Integer] -> [Integer]
sieve  (x:xs) = x : (sieve [y | y <- xs, mod y x > 0])

-- Infinite list of primes!
primes :: [Integer]
primes = sieve [2 .. ]

-- All of the primes, up to the input number
primesTo :: Integer -> [Integer]
primesTo x = sieve [2 .. x]