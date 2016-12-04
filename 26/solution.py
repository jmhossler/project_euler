import math

def main():
    for i in range(999,1,-1):
        if is_prime(i):
            period = 1
            while (10 ** period) % i != 1:
                period += 1
            if period == i - 1:
                print(i, "with cycle", period)
                break


def is_prime(v):
    for i in range(2,int(math.sqrt(v))):
        if v % i == 0:
            return False
    return True

main()
