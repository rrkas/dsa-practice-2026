func mySqrt(x int) int {
    if x < 2 {
        return x
    }

    l, r := 0, x
    ans := 0

    for l <= r {
        mid := l + (r-l)/2

        if mid <= x/mid { // avoids overflow
            ans = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    return ans
}