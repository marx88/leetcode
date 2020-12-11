# 寻找两个正序数组的中位数
[返回首页](../../../README.md)

### 题目
给定两个大小为`m`和`n`的正序（从小到大）数组`nums1`和`nums2`。请你找出并返回这两个正序数组的中位数。

**进阶**：你能设计一个时间复杂度为`O(log(m+n))`的算法解决此问题吗？

### 示例1
```
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
```

### 示例2
```
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
```

### 示例3
```
输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
```

### 示例4
```
输入：nums1 = [], nums2 = [1]
输出：1.00000
```

### 示例5
```
输入：nums1 = [2], nums2 = []
输出：2.00000
```

### 提示
```
1. nums1.length == m
2. nums2.length == n
3. 0 <= m <= 1000
4. 0 <= n <= 1000
5. 1 <= m + n <= 2000
6. -106 <= nums1[i], nums2[i] <= 106
```

### 代码实现
- [GO](../../../golang/lc00001-lc00100/lc00004/median_of_two_sorted_arrays.go) **2091 / 2091** 个通过测试用例；执行用时: **16 ms**；内存消耗: **5.3 MB**；
- [PHP](../../../php/lc00001-lc00100/lc00004/MedianOfTwoSortedArrays.php) **2091 / 2091** 个通过测试用例；执行用时: **16 ms**；内存消耗: **15.3 MB**；


### 心里路程

##### 初版
```
合并两个数组，直到列出(m + n) / 2位置的数字即可break
```

##### 能想到的优化
```
1. 两个数组，长度长的为M、长度短的为N，长度m、n
2. 判断M[m/2] 和 N[n/2]的大小
  - 如果相等中位数就是M、N中位数求均值
  - 如果不等，假设M[m/2]大于N[n/2]，这里的假设反过来也成立继续下面的步骤
    M:      _ _ _ _ _ ...M[m/2]... _ _ _ _ _
    N: _ _ _ ...N[n/2]... _ _ _
3. M[m/2]右侧的数据都会进入新数组中位数的右侧，然后再对M[m/2]左侧和N同时遍历，与M[m/2]比较
  - 遍历值大，直接放入新数组右侧
  - 遍历值小，M[m/2]重新赋值这个小的
  - 一直遍历，直到放入n/2个
```

##### 百度版
```
也叫【第K小数算法】。本题的K等于(m + n) / 2
1. 常识，如果第K小，那么num[k - 1] <= num[k] <= num[k+1]
2. 原理：
   M和N各存在且仅存在一个索引i、j，在M和N合并后满足：M[i-1] <= num[k] <= M[i] 且 N[j-1] <= num[k] <= N[j] 且 M[i-1] <= N[j] 且 N[j-1] <= M[i]。
   话句话说就是：以新数组中位数为界，向左找原先在M中最大的数，向右找原先在M中最小的数，这两个数在M中相邻，同理，N也有这么两个数，假设M中相邻两数的位置为i-1、i，N中相邻两数位置为j-1、j。
   该算法就是找到i和j，终结条件就是M[i-1] <= N[j] 且 N[j-1] <= M[i] 且 i+j=(m+n)/2
3. 循环时的边界：
   上面的i、i-1、j、j-1可能溢出，溢出就表示该数组所有的数全在新数组中位数的某一边，所以不用纠结，因为
   i和j同时加减，所以一直满足i+j=(m+n)/2
   如果i-1或j溢出，那么不用判断M[i-1] <= N[j]，肯定是成立的，i和j-1同理
4. 计算中位数时的边界：
   如果总长度为偶数，那么M[i-1]、N[j-1]两个值取大的，M[i]、N[j]两个值取小的，求均值即可，i、i-1、j、j-1溢出问题也不用纠结，因为
   如果M[i-1]溢出，那么自然可能是获取N[j-1]，如果N[j-1]也溢出，那么肯定是返回0，M[i]、N[j]同理
   如果总长度为奇数，那么M[i]、N[j]两个值取小的，就是中位数，边界问题同上
```

##### 写在后面
`百度版`是看了百度才直到【第K小数算法】，然后花了这4天才写出最终版，中间尝试过两种方式取处理边界问题，`emmmmmm`...

 从这个题开始，可以的话，开始用单元测试，我也不知道这四天往leetcode提交了多少错误代码，`emmmmmm`...