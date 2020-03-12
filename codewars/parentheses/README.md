# 括号校验

## 问题

写一个函数,有个入参是字符串,里面全是`(`或`)`,需要校验括号是否成对出现,
是就返回true,不是就返回false.

要求:字符串长度最多是100个

    "()"              =>  true
    ")(()))"          =>  false
    "("               =>  false
    "(())((()())())"  =>  true

## 性能测试结果

goos: linux
goarch: amd64
pkg: github.com/fight100year/go-ci-test-workflows-template/codewars/parentheses
BenchmarkValid-4      3681162   322 ns/op   16 B/op   4 allocs/op
BenchmarkValidFor-4   63761683  17.8 ns/op  0 B/op    0 allocs/op
BenchmarkValidRegexp-4 234153   4896 ns/op  4696 B/op 31 allocs/op
BenchmarkValidSwitch-4 56024958 21.6 ns/op  0 B/op    0 allocs/op
PASS
ok  github.com/fight100year/parentheses 2.670s

- valid函数是基于字符串一层层从里忘外剥开,性能较差
- validFor是基于for循环检查,效率还可以
- validRegexp 利用正则引擎来检查错误,写法最简单,效率最差
- validSwitch 稍微比if判断差那么一点点

起始还有一个发现:

    for _, x := range str {
      if x == '(' {
        level++
      } else {
        level--
      }
      if level < 0 {
        return false
      }
    }

如果将level是否小于0的判断放在level--下面,会发现性能并没有提高,反而降低了,
why? 循环优化
