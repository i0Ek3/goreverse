/*
 * Package goreverse is a simple tool to reverse given data. 
 * We will offer you a method to reverse any type value in Go.
 *
 */

package goreverse


// Reverse reverses []byte by interval left and right.
func Reverse(b []byte, left, right int) string {
    for left < right {
        b[left], b[right] = b[right], b[left]
        left++
        right--
    }
    return string(b)
}

/*
// ReverseAny reverses any type value with interval left and right.
func ReverseAny(v interface{}, left, right int) string {
    for left < right {
        // TODO
        v[left], v[right] = v[right], v[left]
        left++
        right--
    }
    return v.(string)
}
*/