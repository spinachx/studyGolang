package main

import (
     "fmt"
)

func main() {
     var s string = "hello,世界"
     //s[0] = 'w' 不能修改字符串的值
     fmt.Println("s len: ", len(s))
     //把string转为rune数组
     r := []rune(s)
     fmt.Println("r len: ", len(r))
     //按索引修改rune数组中的值
     r[0] = 'w'
     //直接打印的是unicode编码
     fmt.Println("r: ", r)
     //转为打印字符
     for _, v := range r {
          fmt.Print(string(v))
     }
}
