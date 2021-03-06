# 最长回文子串
[返回首页](../../../README.md)

### 题目
给定一个字符串`s`，找到`s`中最长的回文子串。你可以假设`s`的最大长度为`1000`。

貌似就是最长对称子串。。。

### 示例1
```
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
```

### 示例2
```
输入: "cbbd"
输出: "bb"
```

### 代码实现
- [GO](../../../golang/lc00001-lc00100/lc00005/longest-palindromic-substring.go) **176 / 176** 个通过测试用例；执行用时: **0 ms**；内存消耗: **2.6 MB**；
- [PHP](../../../php/lc00001-lc00100/lc00005/LongestPalindromicSubstring.php) **176 / 176** 个通过测试用例；执行用时: **44 ms**；内存消耗: **15.3 MB**；


### 心里路程

##### 初版
```
声明：给定的字符串定义为母串
1. 母串对称，直接返回
2. 母串去掉右边的字符串产生的子串L对称，直接返回
3. 母串去掉左边的字符串产生的子串R对称，直接返回
4. 以子串L作为母串求最长回文子串
5. 以子串R作为母串求最长回文子串
6. 返回4、5中大的值
可优化：求子串最大回文子串时，可指定一个值，小于该值直接返回该值。如，求子串R时，指定子串L的最大子串长度，这样，避免无效的查找。

写在后面：1000个字符的串就会超过30秒，弃掉。原因，该思路也是暴力计算，但是比暴力计算多了很多重复的查找，所以会超时...
```

##### 第二版
```
1. 反转母串得到新串
2. 假设最长回文子串长度为n，索引从0开始，以长度为n右移，判断长度为n的子串是否在新串中，是的话返回，否则n--

PS：该版还是暴力查找...无非是用空间【反转新串】换了时间【判断是否回文】...
```

##### 中心扩散版
```
1. 遍历母串
2. 当前索引i，向两边扩展，直到两边不相等为止，判断是否最大回文子串，继续下一个

优化：为了解决奇偶性问题，我们可以判断i+1【别名r】是否等于i，如果相等，那就是偶数回文串，顺便r++，如果不等就是奇数回文串，扩散的时候直接r与i-1进行比较，看方法[centerSpread]
```

##### Manacher[马拉车]算法
```
emmmm，这个算法真是...虽然是用空间换时间，你也得知道怎么换呀
这个算法，在这篇文章里说不清，直接看代码吧...
```
