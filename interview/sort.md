## 经典排序算法
1. 排序算法在时间复杂度上分为三个档次：O(n), O(nlgn)，O(n^2)
2. 排序算法的稳定性。如果待排序的列表中存在相同排序值的元素，在排序前后相同排序值的元素排序后相对位置不变。
3. 是否原地排序。也就是说算法是否需要额外空间。

这里的例子都是递减的排序，按时间复杂度分为了三个类别

### 1. O(n^2)
#### 1.1 冒泡排序
[冒泡排序原理](https://zh.wikipedia.org/wiki/%E5%86%92%E6%B3%A1%E6%8E%92%E5%BA%8F)  
冒泡比较简单，每次选出子序列的最大值，放置最前端。  
时间复杂度: O(n^2)  
空间复杂度: O(1)  
是否稳定：是  
是否原地排序：是  

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
插入排序， 类似拿扑克的方式。0～i-1 区间保持有序，循环遍历将第i个元素插入有序区间。
时间复杂度: O(n^2)  
空间复杂度: O(1)  
是否稳定：是  
是否原地排序：是  
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
选择排序跟插入排序的相似点在于也是要区分两个区间。选择排序是交换元素而不是移动。
时间复杂度: O(n^2)  
空间复杂度: O(1)  
是否稳定：否 （交换过程无法保证有序）
是否原地排序：是  

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
[希尔排序](https://zh.wikipedia.org/wiki/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F)  
改良版本的插入排序
时间复杂度: O(n^2)  
空间复杂度: O(1)  
是否稳定：否  
是否原地排序：是  

```go
package main

func shellSort(nums []int) {
    
}

```

### 2. O(nlgn)
#### 2.1 快速排序

#### 2.2 归并排序

#### 2.3 堆排序

### 3. O(n)
#### 3.1 桶排序

#### 3.2 计数排序