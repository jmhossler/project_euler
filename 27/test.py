import pytest
import solution


primes = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59,
          61, 67, 71, 73, 79, 83, 89, 97]

not_prime = [i for i in range(0,101) if i not in primes]


@pytest.mark.parametrize('prime', primes)
def test_is_prime(prime):
    assert solution.is_prime(prime)

@pytest.mark.parametrize('number', not_prime)
def test_is_not_prime(number):
    assert not solution.is_prime(number)

@pytest.mark.parametrize('a, b, consecutive', [
    (1, 41, 40),
    (-79, 1601, 80),
    ])
def test_consecutive_prime_count(a, b, consecutive):
    assert solution.consecutive_prime_count(a, b) == consecutive
