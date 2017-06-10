import math

def func(n, a, b):
    return n ** 2 + a * n + b

def is_prime(n):
    if n < 2:
        return False
    for i in range(2, int(math.sqrt(n))+1):
        if n % i == 0:
            return False
    return True

def consecutive_prime_count(a, b):
    n = 0
    while is_prime(func(n, a, b)):
        n += 1
    return n

def get_max_pair(table):
    max_a, max_b = -999, -1000
    for key_a in table:
        for key_b in table[key_a]:
            if table[key_a][key_b] > table[max_a][max_b]:
                max_a, max_b = key_a, key_b
    return max_a, max_b

def main():
    table = {}
    for a in range(-999, 1000):
        table[a] = {}
        for b in range(-1000, 1001):
            table[a][b] = consecutive_prime_count(a, b)
    a, b = get_max_pair(table)
    print('a: {}\nb: {}\na * b: {}'.format(a, b, a * b))

if __name__ == '__main__':
    main()
