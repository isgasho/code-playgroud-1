## 分治算法
最近看到《算法导论》的分治策略一节，看到的一个题目可以优化引申出来多种解法，同时也可以帮助理解分治策略的一些化整为零的思维。

### 最大子数组问题
 最大子数组：数组 A 中的和最大的非空连续子数组。

#### 暴力解法 O(n^2)
这个问题可以用暴力解法，两层循环遍历，时间复杂度为 O(n^2)，当然最容易想到的并不是最好的解法。

```go
package main

import (
    "fmt"
)

func main() {
    A := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
    left, right, sum := FindMaxSubArray(A)
    fmt.Println(left, right, sum)
}

func FindMaxSubArray(A []int) (left, right, sum int){
    // 这里先不考虑 max 取值的问题，可以取切片的第一个元素或者 int 的最小值。
    var max int
    for i := 0; i < len(A); i++ {
        sum = 0
        for j := i; j < len(A); j++ {
            sum += A[j]
            if sum > max {
                max = sum
                left, right = i, j
            }
        }
    }
    return left, right, sum
}
```

可以得知最大的和为 43，即下标 7, 10 之间的子数组。

#### 分治解法 O(nlgn)
既然这一节是讲分治策略，那么怎么用分治的思想来优化呢。这个解法确实比较难懂，如果让脑袋去跑一遍递归，真的有点累。那么分治本来就是一种局部整体的思想，我们把切片分成三组，左，中，右。那么我们只需要得出，这三个子集的最大值即可。然后再不断分化下去，最后把最大值冒上来。分治解法的关键就是如何用整体局部的思想把问题抽象化。

![image.png](https://upload-images.jianshu.io/upload_images/8573331-43269bf9a510d4fe.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


```go
package main

import (
    "fmt"
)

func main() {
    A := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
    left, right, sum := FindMaxSubArray(A, 0, len(A)-1)
    fmt.Println(left, right, sum)
}

func FindMaxCrossingSubArray(A []int, low, mid, high int) (int, int, int){
    // 这个函数为什么从中间分开算呢？因为得出的结果必须要跟 mid 位相关
    var maxLeft, maxRight int
    // 取负无穷大就行，这里简化处理
    leftSum := -99999
    sum := 0
    for i := mid; i > low; i-- {
        sum += A[i]
        if sum > leftSum {
            leftSum = sum
            maxLeft = i
        }
    }
    rightSum := -99999
    sum = 0
    for j := mid + 1; j < high; j++ {
        sum += A[j]
        if sum > rightSum {
            rightSum = sum
            maxRight = j
        }
    }
    return maxLeft, maxRight, leftSum+rightSum
}

func FindMaxSubArray(A []int, low, high int) (int, int, int) {
    if low == high {
        return low, high, A[low]
    } else {
        mid := (low + high) >> 1
        // 求左半区子数组最大值
        leftLow, leftHigh, leftSum := FindMaxSubArray(A, low, mid)
        // 求右半区子数组最大值
        rightLow, rightHigh, rightSum := FindMaxSubArray(A, mid+1, high)
        // 求包含中位数区子数组的最大值
        crossLow, crossHigh, crossSum := FindMaxCrossingSubArray(A, low, mid, high)
        if leftSum >= rightSum  && leftSum >= crossSum {
            return leftLow, leftHigh, leftSum
        } else if rightSum >= leftSum && rightSum >= crossSum {
            return rightLow, rightHigh, rightSum
        } else {
            return crossLow, crossHigh, crossSum
        }
    }
}

```

时间复杂度分析：不用推理的方式，

可以将时间复杂度降低到 O(n) 吗？贪心算法？

#### 线性解法 O(n)


## 题目
[53 最大子序和]([https://leetcode-cn.com/problems/maximum-subarray/](https://leetcode-cn.com/problems/maximum-subarray/)
)
