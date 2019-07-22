package main

/*
 解题思路：可以用分治的思维，求一个列表的最大和，要么在左半区，要么在右半区，要么在包含中点的区域。时间复杂度O(nlgn)
 还可以用动态规划的思想把时间复杂度降到O(n)。
*/

func maxSubArray(nums []int) int {
    return MaxSubArray(nums, 0, len(nums)-1)
}

func MaxSubArray(nums []int, low, high int) int {
    if low == high {
        return nums[low]
    }
    mid := (low + high) >> 1
    leftSum := MaxSubArray(nums, low, mid)
    rightSum := MaxSubArray(nums, mid+1, high)
    crossSum := FindCrossArray(nums, low, mid, high)
    if leftSum >= crossSum && leftSum >= rightSum {
        return leftSum
    } else if rightSum >= crossSum && rightSum >= leftSum {
        return rightSum
    } else {
        return crossSum
    }
}

func FindCrossArray(nums []int, low, mid, high int) int {
    var leftSum, rightSum int
    leftMaxSum := -999999
    for i := mid; i >= low; i-- {
        leftSum += nums[i]
        if leftSum > leftMaxSum {
            leftMaxSum = leftSum
        }
    }
    rightMaxSum := -999999
    for j := mid+1; j <= high; j++ {
        rightSum += nums[j]
        if rightSum > rightMaxSum {
            rightMaxSum = rightSum
        }
    }
    return leftMaxSum + rightMaxSum
}

/*
动态规划解法：F[i] = max(F[i-1]+nums[i], nums[i])
*/

func maxSubArray(nums []int) int {
    maxNum := nums[0]
    v := make([]int, len(nums))
    v[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        v[i] = max(nums[i], v[i-1] + nums[i])
        if v[i] > maxNum {
            maxNum = v[i]
        }
    }
    return maxNum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
