### bracketsmatch 括号匹配

##### ((({(sdafsaf){}))



循环遍历字符串 {
    if str[i]是左括号 {
        push
    } else {
        if 栈为空 {
            匹配失败
            结束
        } else {
            弹出栈顶的左括号
        }
    }
}
if 栈为空 {
    yes
} else {
    no
}