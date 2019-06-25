## 二分搜索

## 简介
[原理解析](https://www.geeksforgeeks.org/binary-search/) 
二分搜索也叫折半搜索，是一种在有序数组中查找某特定元素的搜索算法。搜索从中间开始，如果该元素大于目标值，则目标值在该元素左边的区间。

## 步骤
1. 令 left 为 0，right 为 len(nums) - 1
2. 如果 left > right，则搜索结束
3. mid = (left + right) / 2，为防止溢出可以(left + (right - left) / 2)
4. 如果nums[mid] < target，则说明在 mid 右边的区间，left = mid + 1，跳步骤2
5. 如果nums[mid] > target，则说明在 mid 左边的区间，right = mid - 1，跳步骤2
5. nums[mid] == target，搜索终止，返回 mid

## 实现
时间复杂度: O(lgn)
空间复杂度: O(1)

```go
package main

func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left < right {
        mid := (left + right) >> 1
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}
``` 

## 变种
二分搜索法不仅仅在有序的数组可以用，回转有序数组（即有序数组向左或者右偏移）也可以用来搜索。

回转数组示例：
```go
// 有序数组
a = [1, 2, 3, 4, 5, 6, 7] 

// 以下均为 回转 数组
a = [2, 3, 4, 5, 6, 7, 1]
a = [5, 6, 7, 1, 2, 3, 4]
```

具体思路：由于回转数组是有序数组的一个偏移，我们可以把题目转换成求回转数组的最小值，然后索引还原，就跟普通二分搜索一样了。
假设为回转数组，步骤：
1. 令 left 为 0，right 为 len(nums) - 1
2. 如果 left > right，则搜索结束
3. mid = (left + right) / 2，为防止溢出可以(left + (right - left) / 2)
4. 如果nums[mid] > nums[right]，说明最小值在右边区间，left = mid + 1。因为子区间都是单调递增的，所以数组的最左值肯定大于右值，那么最小值在右边。
5. 如果nums[mid] <= nums[right]，right = mid。
6. left 即最小值索引

## 变种实现
```go
package main

// 保证 nums 是回转数组
func minIndex(nums []int) int {
    left, right := 0, len(nums) - 1
    for left < right {
        mid := (left + right) >> 1
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func binarySearch(nums []int, target int) int {
    minIdx := minIndex(nums)
    left, right := 0, len(nums) - 1
    for left < right {
        mid := (left + right) >> 1
        // 中位数的真实位置
        realMid := (minIdx + mid) % len(nums)
        if nums[realMid] == target {
            return realMid
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}
```