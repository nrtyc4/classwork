# find pi to the nth digit using Machin's formula

from math import *

digits = raw_input('Enter number of digits to round Pi to: ')

# calculate and print pi using Machin formula, and nested string formatting
print ( '{0:.%df}' % min(30, int(digits))).format(4 * (4 * atan(1.0/5.0) - atan(1.0/239.0)))
