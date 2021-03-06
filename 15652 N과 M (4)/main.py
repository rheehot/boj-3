from itertools import combinations_with_replacement


n, m = input().split()
n, m = int(n), int(m)


for seq in combinations_with_replacement(range(1, n+1), m):
    print(' '.join(str(x) for x in seq))
