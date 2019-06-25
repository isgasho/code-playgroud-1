package heap

func siftUp(nums []int, i, end int) {
    l := 2*i + 1
    r := 2*i + 2
    largest := i
    if l <= end && nums[largest] < nums[l] {
        largest = l
    }
    if r <= end && nums[largest] < nums[r] {
        largest = r
    }
    if i != largest {
        nums[i], nums[largest] = nums[largest], nums[i]
        siftUp(nums, largest, end)
    }
}

func heapify(nums []int) {
    // first non-leaf node
    base := len(nums) / 2 - 1
    for i := base; i >=0; i-- {
        siftUp(nums, i, len(nums)-1)
    }
}

func findKthLargest(nums []int, k int) int {
    heapify(nums)
    for i := len(nums) - 1; i >= 0; i-- {
        k--
        if k == 0 {
            return nums[0]
        }
        nums[0], nums[i] = nums[i], nums[0]
        siftUp(nums, 0, i-1)
    }
    return nums[0]
}
