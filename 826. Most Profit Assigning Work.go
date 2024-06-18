type Job struct {
    profit, diff int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    jobs := []Job{}

    for i := 0; i < len(profit); i++ {
        jobs = append(jobs, Job{profit[i], difficulty[i]})
    }

    sort.Slice(jobs, func(a, b int) bool {
        return jobs[a].profit > jobs[b].profit
    })

    sort.Slice(worker, func(a, b int) bool {
        return worker[a] > worker[b]
    })

    result := 0
    nextWorker := 0
    for i := 0; i < len(jobs); i++ {
        for nextWorker < len(worker) && worker[nextWorker] >= jobs[i].diff {
            result += jobs[i].profit
            nextWorker++
        }
    }

    return result
}
