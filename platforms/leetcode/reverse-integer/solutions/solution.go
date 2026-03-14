import "math"

func reverse(x int) int {
    r := 0
    mul := 1
    if x < 0 {
        x = -x
        mul = -1
    }
    for x > 0 {
        d := x % 10
        r = r*10+d
        x = x / 10
    }

    r = mul*r

    val2e31 := int(math.Pow(2, 31))

    if r < -val2e31 || r > val2e31-1 {
        return 0
    }

    return r
}