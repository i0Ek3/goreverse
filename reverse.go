/*
 * Package goreverse is a simple tool to reverse given data. 
 * We will offer you a method to reverse any type value in Go.
 *
 */

package goreverse


// ReverseWithInterval reverses []byte by interval left and right.
func ReverseWithInterval(b []byte, left, right int) string {
    for left < right {
        b[left], b[right] = b[right], b[left]
        left++
        right--
    }
    return string(b)
}

// Reverse reverses given res
func Reverse(res []int) []int {
    for i := 0; i < len(res)/2; i++ {
        res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
    }
    return res
}

/*
// ReverseAny reverses any type value with interval left and right.
func ReverseAny[T any](v T, left, right int) string {
    for left < right {
        v[left], v[right] = v[right], v[left]
        left++
        right--
    }
    return v.(string)
}
*/
