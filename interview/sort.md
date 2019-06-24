## 经典排序算法
1. 排序算法在时间复杂度上分为三个档次：O(n)，O(nlgn)，O(n^2)
2. 排序算法的稳定性。如果待排序的列表中存在相同排序值的元素，在排序前后相同排序值的元素排序后相对位置不变。
3. 是否原地排序。也就是说算法是否需要额外空间。

这里的例子都是递减的排序，按时间复杂度分为了三个类别

### 1. O(n^2)
#### 1.1 冒泡排序
[冒泡排序原理](https://zh.wikipedia.org/wiki/%E5%86%92%E6%B3%A1%E6%8E%92%E5%BA%8F)  
[冒泡算法解析](https://www.geeksforgeeks.org/bubble-sort/)  
冒泡比较简单，每次选出子序列的最大值，放置最前端。  

最优时间复杂度 | 最坏时间复杂度 | 时间复杂度 | 空间复杂度 | 稳定排序 | 原地排序 
---  | --- | ---  | --- | --- | ---
O(n) | O(n^2) | O(n^2) | O(1) | ✅ | ✅ 

```go
package main

func bubbleSort(nums []int) {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] < nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
}
```

#### 1.2 插入排序
[插入排序原理](https://zh.wikipedia.org/wiki/%E6%8F%92%E5%85%A5%E6%8E%92%E5%BA%8F)  
[插入排序解析](https://www.geeksforgeeks.org/insertion-sort/)  
插入排序， 类似拿扑克的方式。0～i-1 区间保持有序，循环遍历将第i个元素插入有序区间。

最优时间复杂度 | 最坏时间复杂度 | 时间复杂度 | 空间复杂度 | 稳定排序 | 原地排序 
---  | --- | ---  | --- | --- | ---
O(n) | O(n^2) | O(n^2) | O(1) | ✅ | ✅ 

```go
package main

func insertSort(nums []int) {
    for i := 1; i < len(nums); i++ {
        key := nums[i]
        // 保证 0 ~ i-1 有序
        j := i - 1
        for j >= 0 && nums[j] < key {
            nums[j + 1] = nums[j]
            j--
        } 
        // 填坑 插入位置
        nums[j + 1] = key
    }
}
```

#### 1.3 选择排序
[选择排序原理](https://zh.wikipedia.org/wiki/%E9%80%89%E6%8B%A9%E6%8E%92%E5%BA%8F)  
[选择排序解析](https://www.geeksforgeeks.org/selection-sort/)  
选择排序跟插入排序的相似点在于也是要区分两个区间。选择排序是交换元素而不是移动。

最优时间复杂度 | 最坏时间复杂度 | 时间复杂度 | 空间复杂度 | 稳定排序 | 原地排序 
---  | --- | ---  | --- | --- | ---
O(n^2) | O(n^2) | O(n^2) | O(1) | ❌ | ✅ 

```go
package main

func selectSort(nums []int) {
    // start 为无序起始位置 max 为区间最大值的位置
    start, max := 0, 0
    for i := 0; i < len(nums); i++ {
        // 找出区间最大值 max
        for j := i; j < len(nums); j++ {
            if nums[j] > nums[max] {
                max = j
            }
        }
        // 筛出区间最大元素放入左边
        if nums[max] > nums[start] {
            nums[max], nums[start] = nums[start], nums[max]
        }
        max = start + 1
        start++
    }
}
```

#### 1.4 希尔排序
[希尔排序原理](https://zh.wikipedia.org/wiki/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F)  
[希尔排序解析](https://www.geeksforgeeks.org/shellsort/) 
改良版本的插入排序，把步长 step 替换为1，发现和插入排序一摸一样。

最优时间复杂度 | 最坏时间复杂度 | 时间复杂度 | 空间复杂度 | 稳定排序 | 原地排序 
---  | --- | ---  | --- | --- | ---
O(n) | O(n^2) | O(nlgn) | O(1) | ❌ | ✅ 

```go
package main

func shellSort(nums []int) {
    // step 为步长 每次对半分 ps: 按维基百科介绍有比较好的 step 公式，这里取一个比较简单的规则
    step := len(nums) >> 1
    for step > 0 {
        // 步长内插入排序 注意是从后到前
        for i := step; i < len(nums); i++ {
            // 每列最后一个元素
            key := nums[i]
            j := i - step
            // 按步长
            for j >= step - 1 && nums[j] < key {
                nums[j + step] = nums[j]
                j -= step
            }
            nums[j + step] = key
        }
        step = step >> 1
    }
}
```

### 2. O(nlgn)
#### 2.1 快速排序
[快速排序原理](https://zh.wikipedia.org/wiki/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F)  
[快速排序解析](https://www.geeksforgeeks.org/quick-sort/)  
快速排序运用的也是分治的思想。挑选基准值，然后把小于基准值的放左边，大于的放右边。最后迭代到1的时候就是排序结束。

最优时间复杂度 | 最坏时间复杂度| 平均时间复杂度 | 空间复杂度 | 稳定排序 | 原地排序 
---  | --- |---  | --- | --- | ---
O(nlgn) | O(n^2) | O(nlgn) | O(lgn) | ❌ | ✅ 

```go
package main

func partition(nums []int, p, q int) int {
    // 基准值的选择 有更优方式 这里简化
    // partition 的方式实现的有点巧妙，可以想象成如何用O(n) 的算法把序列按给定数字 分成大于和小于的两份。
    // 类似于选择排序 所以是不稳定的 采用双指针思想 i 记录分割点，j 遍历交换。
     base := nums[p]
     // i 记录按大小划分的位置
     i := p
     for j := p + 1; j <= q; j++ {
          if nums[j] > base {
               nums[j], nums[i] = nums[i], nums[j]
               i++
          }
     }
     nums[i], nums[q] = nums[q], nums[i]
     return i
}

func helper(nums []int, p, q int){
    if p >= q {
        return
    }
    r := partition(nums, p, q)
    helper(nums, p, r)
    helper(nums, r+1, q)
}

func quickSort(nums []int){
    helper(nums, 0, len(nums) - 1)
}
```

#### 2.2 归并排序
[归并排序原理](https://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F)  
[归并排序解析](https://www.geeksforgeeks.org/merge-sort/)  
归并排序使用了分治的思想。可以用递归实现也可以用迭代实现。
分：分解待排序的 n 个元素的序列成各具 n/2 个元素的两个子序列
解决：排序子序列
合：合并两个已排序的子序列

最优时间复杂度 | 最坏时间复杂度| 平均时间复杂度 | 空间复杂度 | 稳定排序 | 原地排序 
---  | --- |---  | --- | --- | ---
O(nlgn) |O(nlgn) |O(nlgn) | O(nlgn) | ✅ | ❌ 

```go
package main

func helper(nums []int, p, q int) {
    if q <= p {
        return 
    }
    r := (q + p) >> 1
    helper(nums, p, r)
    helper(nums, r+1, q)
    merge(nums, p, q, r)
}

func merge(nums []int, p, q, r int) {
    i, j := p, r+1 
    var res []int
    for i <= r && j <= q {
        if nums[i] > nums[j] {
            res = append(res, nums[i])
            i++
        } else {
            res = append(res, nums[j])
            j++
        }
    }
    for i <= r {
        res = append(res, nums[i])
        i++
    }
    for j <= q {
        res = append(res, nums[j])
        j++       
    }
    // 已排序区间
    for i := 0; i < q-p+1; i++ {
        nums[p+i] = res[i]
    }
}

func mergeSort(nums []int) {
    helper(nums, 0, len(nums)-1)
}
```

#### 2.3 堆排序
[堆排序原理](https://zh.wikipedia.org/zh-hans/%E5%A0%86%E6%8E%92%E5%BA%8F)  
[堆排序解析](https://www.geeksforgeeks.org/heap-sort/)  
主要需要了解堆这个数据结构，以及如何构建。由于我们的排序，都是降序，这里讨论小顶堆。

由于堆是完全二叉树，通常堆是通过一维数组来实现的。在数组起始位置为0的情形中：
父节点i的左子节点在位置 (2i+1)
父节点i的右子节点在位置 (2i+2)
子节点i的父节点在位置 floor(i/2)

大顶堆中的最大值总是位于根节点（在优先队列中使用堆的话堆中的最小值位于根节点）。堆中定义以下几种操作：
最大堆调整（Max Heapify）：将堆的末端子节点作调整，使得子节点永远小于父节点
创建最大堆（Build Max Heap）：将堆中的所有数据重新排序
堆排序（HeapSort）：移除位在第一个数据的根节点，并做最大堆调整的递归运算

核心点在于堆化的实现（heapify）。

最优时间复杂度 | 最坏时间复杂度| 平均时间复杂度 | 空间复杂度 | 稳定排序 | 原地排序 
---  | --- |---  | --- | --- | ---
O(nlgn) |O(nlgn) |O(nlgn) | O(1) | ❌ | ✅ 



```go
package main

// 自底向上 使每个点 满足堆条件
func siftUp(nums []int, i, end int) {
    smallest := i
    l := 2*i + 1
    r := 2*i + 2
    if end >= r && nums[i] > nums[r] {
        smallest = r
    }
    if end >= l && nums[smallest] > nums[l] {
        smallest = l
    }
    if i != smallest {
        nums[i], nums[smallest] = nums[smallest], nums[i]
        siftUp(nums, smallest, end)
    }
}

func heapify(nums []int) {
    base := len(nums) / 2 - 1
    // 从第一个非叶子节点开始 直到 root
    for i := base; i >= 0; i-- {
        siftUp(nums, i, len(nums) - 1)
    }
}

func heapSort(nums []int) {
    // 构建小顶堆 堆顶为最大值 0~i 依次取堆顶 即完成排序
    heapify(nums)
    for i := len(nums) - 1; i >= 0; i-- {
        // 删除元素 堆尾放回堆顶 重新构造 最小值放堆尾进行排序
        nums[i], nums[0] = nums[0], nums[i]
        siftUp(nums, 0, i-1)
    }
}

```

### 3. O(n) TODO
#### 3.1 基数排序
[基数排序原理](https://zh.wikipedia.org/wiki/%E5%9F%BA%E6%95%B0%E6%8E%92%E5%BA%8F)  
[基数排序解析](https://www.geeksforgeeks.org/radix-sort/)  

#### 3.2 计数排序
[计数排序解析](https://www.geeksforgeeks.org/counting-sort/)  

#### 3.3 桶排序
[桶排序解析](https://www.geeksforgeeks.org/bucket-sort-2/)  

## 总结
算法名称 | 最优时间复杂度 | 最坏时间复杂度 | 时间复杂度 | 空间复杂度 | 稳定排序 | 原地排序 
---  | ---  | --- | ---  | --- | --- | ---
冒泡排序 | O(n) | O(n^2) | O(n^2) | O(1) | ✅ | ✅ 
插入排序 | O(n) | O(n^2) | O(n^2) | O(1) | ✅ | ✅ 
选择排序 | O(n^2) | O(n^2) | O(n^2) | O(1) | ❌ | ✅ 
希尔排序 | O(n) | O(n^2) | O(nlgn) | O(1) | ❌ | ✅ 
快速排序 | O(nlgn) | O(n^2) | O(nlgn) | O(lgn) | ❌ | ✅ 
归并排序 | O(nlgn) |O(nlgn) |O(nlgn) | O(nlgn) | ✅ | ❌ 
堆排序 | O(nlgn) |O(nlgn) |O(nlgn) | O(1) | ❌ | ✅ 
基数排序 | O(k*n) | O(k*n) | O(k*n) | O(k+n) | ✅ | ✅ 