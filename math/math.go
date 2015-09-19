package math

func Max(x float64, y ...float64) float64 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(x float64, y ...float64) float64 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

func Round(x float64) float64 {
	if x > 0.0 {
		return float64(int64(x + 0.5))
	}
	return float64(int64(x - 0.5))
}
