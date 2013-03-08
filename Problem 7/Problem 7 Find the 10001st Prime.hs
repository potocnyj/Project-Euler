-- Returns the nth prime number in the set of integers
nthPrime :: Int -> Integer

nthPrime n = (findPrime [2] n)

findPrime :: [Integer] -> Int -> Integer
findPrime n 0 = 1
findPrime (x:xs) 1 = x
findPrime (x:xs) m 
	| (length (x:xs)) == m 	= x
	| otherwise 		= findPrime ((nextPrime (x+1) xs):x:xs) m


nextPrime :: Integer -> [Integer] -> Integer
nextPrime n m
	| prime n m 	= n
	| even n 	 	= nextPrime (n+1) m
	| otherwise 	= nextPrime (n+2) m


-- Function prime
-- Input: integer n, list of integers xs
-- Returns: Boolean
-- prime takes in an integer value as input,
-- and uses the definition of a prime number
-- to determine if it is prime. If it is prime,
-- it returns true, otherwise it returns false
prime :: Integer -> [Integer] ->  Bool

prime n xs
	| n == 2															= True
	| (even n) 															= False
	| (length ([x | x <- xs, mod n x == 0]) == 0)						= True
	| otherwise 														= False