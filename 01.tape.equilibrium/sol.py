def solution(A):
    S = sum(A)
    tape = 99999999999
    cur = 0
    for ai in A[:-1]:
        cur += ai
        cur_tape = abs(2*cur-S)
        if cur_tape <= tape:
            tape = cur_tape
    return tape
