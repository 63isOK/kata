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
BenchmarkValid-4      3681162   322 ns/op   16 B/op 4 allocs/op
BenchmarkValidFor-4   63761683  17.8 ns/op  0 B/op  0 allocs/op
PASS
ok  github.com/fight100year/parentheses 2.670s

valid函数是基于字符串一层层从里忘外剥开,性能最差,
validFor是基于for循环检查,效率还可以
