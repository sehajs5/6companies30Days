package main

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	x1 -= xCenter
	x2 -= xCenter
	y1 -= yCenter
	y2 -= yCenter
	minX := x1 * x2
	if minX > 0 {
		minX = min(x1*x1, x2*x2)
	} else {
		minX = 0
	}

	minY := y1 * y2
	if minY > 0 {
		minY = min(y1*y1, y2*y2)
	} else {
		minY = 0
	}
	return minY+minX <= radius*radius
}
