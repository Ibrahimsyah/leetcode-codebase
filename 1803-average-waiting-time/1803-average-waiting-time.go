func averageWaitingTime(customers [][]int) float64 {
    totalWaiting := 0
    startTime := 0
    for _, order := range customers {
        arrive := order[0]
        process := order[1]
        if arrive > startTime {
            startTime = arrive
        }
        finish := startTime + process
        totalWaiting += finish - arrive
        startTime = finish

    }

    return float64(totalWaiting) / float64(len(customers))
}