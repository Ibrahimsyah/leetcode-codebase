func minTimeToVisitAllPoints(points [][]int) int {
    if len(points) == 1 {
        return 0
    }

    start := points[0]
    destinationIndex := 1
    totalDistance := 0
    for destinationIndex < len(points) {
        distance := 0
        origin := start
        destination := points[destinationIndex]
        for origin[0] != destination[0] || origin[1] != destination[1] {
            origin[0], origin[1] = travel(origin[0], origin[1], destination[0], destination[1])
            distance ++
        }
        totalDistance += distance
        start = destination
        destinationIndex++
    }

    return totalDistance
}

func travel(xStart, yStart, xDest, yDest int) (int, int) {
    if xDest == xStart {
        if yDest > yStart {
            return xStart, yStart + 1
        }

        return xStart, yStart - 1
    }

    if xDest < xStart {
        if yDest > yStart {
            return xStart - 1, yStart + 1
        } else if yDest < yStart {
            return xStart - 1, yStart - 1
        }
        return xStart - 1, yStart
    }

    if yDest > yStart {
        return xStart + 1, yStart + 1
    } else if yDest < yStart {
        return xStart + 1, yStart - 1
    }
    return xStart + 1, yStart
} 