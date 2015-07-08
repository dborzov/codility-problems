from sol import solution

def test_min_tape():
    assert solution([3,5,7,9]) == 6
    assert solution([2000,4000]) == 2000
    assert solution([1000, -1000]) == 2000
    assert solution([2,3]) == 1
    assert solution([3,1,2,4,3]) == 1
