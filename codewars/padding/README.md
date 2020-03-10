# padding 填充游戏

## 题目

设计一个函数,对于一个正整数,输出一个包含正整数的字符数组.
要求:如果不够N位,前面填充0;如果N为-1,原样输出

正整数 | 位数 | 输出
---|---|---
2019|2|"2019"
2019|4|"2019"
2019|5|"02019"
3|-1|"3"

## 两种实现的测试结果

    ➜  padding git:(padding) ✗ go test -bench="." -benchmem
    goos: linux
    goarch: amd64
    pkg: github.com/fight100year/go-ci-test-workflows-template/codewars/padding
    BenchmarkPadding-4         	  927896   	      1163 ns/op	     120 B/op	      15 allocs/op
    BenchmarkPaddingNoStd-4    	11957875	        98.1 ns/op	      16 B/op	       3 allocs/op
    BenchmarkPaddingNoStd2-4   	11913781	        98.9 ns/op	      16 B/op	       3 allocs/op
    BenchmarkPaddingStd-4      	12566420	        90.9 ns/op	      16 B/op	       3 allocs/op

- padding是基于标准库实现的,性能最差,和其他实现相差一个数量级
- std是标准库中的实现,NoStd是手写的
  - 两者的差别是取模运算%,这儿点差别就是10%左右的性能差别
  - 标准库是利用多个变量来实现,原理是 `a - (a/10)*10`
  - 猜测是循环对多变量的简单计算有优化
- NoStd有两个版本,一个是将for循环的循环因子放一起,一个是分成两个for
  - 从结果上看,性能基本差不多
  - 从重构的角度看,相似代码越少越好
