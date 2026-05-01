import random
import math
def lehman(p):
    for i in range(p-1):
        a = random.randint(1,p-1)
        x = math.pow(a,(p-1)/2) % p
        if(x!=1 and x!=p-1):
           return False
    return True, (1-(1/math.pow(2,p-1))) *100

print(lehman(17))
