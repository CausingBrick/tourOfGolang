//Package exercise population count
//TODO need to add test function
package exercise

//pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//PopCount returns the population count of x. By using range method
// Exercise2.3
func PopCount(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>uint(i*8))])
	}
	return count
}

// PopCountEB returns the population count of x. By range each bit of x.
// Exercise2.4
func PopCountEB(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if (x>>uint(i))&uint64(1) == 1 {
			count++
		}
	}
	return count
}

// PopCountRev returns the population count of x. By algorithm of x&(x-1).
// Exercise2.5
func PopCountRev(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x&(x-1) == 0 {
			count++
			x = x >> 1
		}
	}
	return count
}
