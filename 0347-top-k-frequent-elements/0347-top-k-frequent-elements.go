func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for _, num := range nums {
        m[num]++
    }

    // create array of 0..n
    // where n is the len of nums + 1
    // to store the num frequency.
    // The best case is all num in nums is same
    // so that arr[len(nums)-1] = nums
    arr := make([][]int, len(nums) + 1)
    for k, v := range m {
        arr[v] = append(arr[v], k)
    }

    // get top k frequent
    res := []int{}
    for i := len(arr) - 1; i >= 0; i-- {
        for _, num := range arr[i] {
            res = append(res, num)
            if len(res) == k {
                return res
            }
        }
    }

    return res
}