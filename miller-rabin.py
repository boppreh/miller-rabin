import sys, math, random

# Receive candidate number from arguments, or default to 221 for test purposes.
if len(sys.argv) > 1:
	candidate = eval(sys.argv[1])
else:
	candidate = 221

modulo = candidate - 1

# Write the modulo (candidate -1) number in the form
# 2^s * d.

s = 0
quotient = modulo
remainder = 0
while remainder == 0:
	quotient, remainder = divmod(quotient, 2)
	s += 1

# The last division failed, so we must decrement `s`.
s -= 1
# quotient here contains the leftover which we could not divide by two,
# and we have a 1 remaining from this last division. 
d = quotient * 2 + 1

# Here 10 is the precision. Every increment to this value decreases the
# chance of a false positive by 3/4.
for k in range(10):

	# Every witness may prove that the candidate is composite, or assert
	# nothing.
	witness = random.randint(2, modulo - 1)

	x = pow(witness, d, candidate)
	if x == 1 or x == modulo:
		continue

	for i in range(s - 1):
		x = pow(x, 2, candidate)

		if x == 1:
			print('Composite.')
			exit()

		if x == modulo:
			break

	if x != modulo:
		# We arrived here because the `i` loop ran its course naturally without
		# meeting the `x == modulo` break.
		print('Composite.')
		exit()

print('Prime.')