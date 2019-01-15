from itertools import product


_, m = input().split()
m = int(m)
nums = [int(x) for x in input().split()]
nums.sort()


def unique(iterable):
    seen = set()
    for x in iterable:
        if x in seen:
            continue
        yield x
        seen.add(x)


for seq in unique(product(nums, repeat=m)):
    print(' '.join(str(x) for x in seq))
