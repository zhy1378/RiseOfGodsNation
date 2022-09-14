package main

import (
	"fmt"
	"io/ioutil"
    "log"
    "strings"
    //"bytes"
    "unicode/utf8"
    "time"
)

// 将要脱去的字符
const punctions = "，。！“”…‘’§￥%&、（=？·）—*+# -_,.;:`!<>?\r\n\t"
const DIRECTORY = "/Users/zhy1378/Documents/RiseOfGodsNation/1.Im-AndMortal/300/01/"

func main() {
    //var x = 38+11+22+10+27+65+22+42+17+25+23+24+25+24+26+22+26+30+16+86+13+37+54+66+56+37+70+100+113+19+84+30+100+100+34+100+46+100+38+100+36+100+68+100+10+40+135+42+63+74+52
    //fmt.Printf("%d", x)

    // 遍历文件夹
	files, err := ioutil.ReadDir(DIRECTORY)
	if err != nil {
		log.Fatal(err)
    }

    var totalLength, totalNumRunes, totalDryNumRunes = 0, 0, 0
    var filePath string

	for _, file := range files {
        // 只计算md文件
        if(strings.HasSuffix(file.Name(), ".md")){
            filePath = DIRECTORY + file.Name()
            //fmt.Println(filePath)
            length, numRunes, dryRunes := CountMd(filePath)
            totalLength+= length
            totalNumRunes += numRunes
            totalDryNumRunes += dryRunes
        }
    }

    t := time.Now()
    fmt.Printf(t.Format("2006-01-02 15:04:05 "))
    fmt.Printf("文件字节数：%d，字符总数:%d，脱水后：%d\n", totalLength, totalNumRunes, totalDryNumRunes)
}

/**
 * 统计md文件中的字符数量。返回值为：
 * 整体字节数
 * 脱水前字符总数
 * 脱水后字符总数
 */
func CountMd(file string)(int, int, int){
    content, err := ioutil.ReadFile(file) // 读取文件
	if err != nil {
		log.Fatal(err)
    }
    
    var length = len(content) // 长度
    var numRunes = utf8.RuneCount(content) // 字符总数
    var dryRunes = 0

    for i, w := 0, 0; i < len(content); i += w {
        // 取出一个rune，也即utf8字符
        runeValue, width := utf8.DecodeRune(content[i:])
        var s = string(runeValue)
        // 是否应该被脱掉的字符
        if(!strings.Contains(punctions, s)){
            dryRunes++
            //fmt.Printf(s)
        }
        w = width
    }

    //fmt.Printf("文件字节数：%d，字符总数:%d，脱水后：%d\n", length, numRunes, dryRunes)

    return length, numRunes, dryRunes
}