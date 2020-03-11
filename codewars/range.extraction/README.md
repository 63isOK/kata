# range extraction

## 题目

A format for expressing an ordered list of integers
is to use a comma separated list of either individual integers
Or a range of integers denoted by the starting integer
separated from the end integer in the range by a dash, '-'.
(The range includes all integers in the interval including both endpoints)

The range syntax is to be used only for,
and for every range that expands to more than two values.

Example

The list of integers:

-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20

Is accurately expressed by the range expression:

-6,-3-1,3-5,7-11,14,15,17-20

(And vice-versa).

Task

Create a function that takes a list of integers in increasing order
and returns a correctly formatted string in the range format;
Use the function to compute and print the range formatted version
of the following ordered list of integers.
(The correct answer is: `0-2,4,6-8,11,12,14-25,27-33,35-39`).

     0,  1,  2,  4,  6,  7,  8, 11, 12, 14,
    15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
    25, 27, 28, 29, 30, 31, 32, 33, 35, 36,
    37, 38, 39

Show the output of your program.

## 题目解读

有序整数列表的表示方法,要么是单个整数,要么是`范围整数`的写法,
所谓范围整数的写法,就是超过3个连续的整数可写成a-b,a是最小数,
b是最大数.不管是单个整数还是范围整数,都是用逗号隔开

## 测试结果

    goos: linux
    goarch: amd64
    pkg: github.com/fight100year/go-ci-test-workflows-template/codewars/range.extraction
    BenchmarkSolution-4   936351  1126 ns/op  400 B/op  22 allocs/op
    BenchmarkSolution2-4  675604  1513 ns/op  272 B/op  24 allocs/op
    PASS
    ok github.com/fight100year/codewars/range.extraction  2.112s

第二种解法使用了fmt.Sprintf来格式化输出,效率比strconv库差太多了,
其次,这类问题使用普通的for语句真心好处理
