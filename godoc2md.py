#!/usr/bin/env python
#coding=utf-8

"""
根据翻译的格式约定
这个脚本用于半自动化生成待翻译的文档
例如要翻译bufio包
python godoc2md.py bufio > bufio.md

开始编辑翻译bufio.md吧!

by getwe

TODO:
    还没搞定在哪个地方插入 "##内容".暂时先手动加一下

"""

import os
import sys
import string
import re

def getPackage(package):
    cmdStr = "cd $GOROOT/src/pkg; godoc " + package
    handle = os.popen(cmdStr)

    # 结构体定义
    structPattern = re.compile("^type\s(.*?)\s\w+\s{")

    # 普通函数
    funcPattern = re.compile("^func\s(\w+)\(.*?\)")

    # 方法
    methodPattern = re.compile("^func\s\(\w.*?\s(.*?)\)\s(\w+)")

    print "#%s"%package
    print ""
    print "import \"%s\""%package
    print "\n##简介\n\n##概览"

    for line in handle.readlines():
        line = line.strip('\n')
        
        if line == "CONSTANTS":
            print "###常量"
            continue
        if line == "VARIABLES":
            print "###变量"
            continue

        # 结构体
        match = structPattern.match(line)
        if match:
            name = match.group(1)
            print "###type %s"%(name)
            print "```go"
            print line
            #结构体定义有可能是多行,还不能加下面这行
            #print "```"
            continue

        # 尝试匹配普通函数
        match = funcPattern.match(line)
        if match:
            className = match.group(1)
            print "###func %s"%(className)
            print "```go"
            print line
            print "```"
            continue

        # 尝试匹配类方法
        match = methodPattern.match(line)
        if match:
            className = match.group(1)
            methodName = match.group(2)
            print "###func (%s) %s"%(className,methodName)
            print "```go"
            print line
            print "```"
            continue

        # 常量变量
        if line == "const (" or line == "var (":
            print "```go"
            print line
            continue

        # 多行定义的结束
        if line == ")" or line == "}":
            print line
            print "```"
            continue


        # 无法识别的行直接输出
        print line


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print "Usage %s packageName"%sys.argv[0]
        sys.exit(1)

    getPackage(sys.argv[1])
