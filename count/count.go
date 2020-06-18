package count

// Max x, y elements
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MaxOf3 x, y, z elements
func MaxOf3(x, y, z int) int {
	if y < x && x > z {
		return x
	} else if x < y && y > z {
		return y
	}
	return z
}

// MaxOfN n elements
func MaxOfN(list []int) int {
	max := list[0]
	for _, e := range list {
		if e > max {
			max = e
		}
	}
	return max
}
