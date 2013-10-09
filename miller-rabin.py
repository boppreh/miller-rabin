import sys, math, random

if len(sys.argv) > 1:
	candidate = eval(sys.argv[1])
else:
	candidate = 221

modulo = candidate - 1
s = -1
quotient = modulo
remainder = 0
while remainder == 0:
	quotient, remainder = divmod(quotient, 2)
	s += 1
d = quotient * 2 + 1

for k in range(10):
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
		print('Composite.')
		exit()
print('Prime.')