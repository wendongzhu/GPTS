package main

import (
	"encoding/xml"
	"fmt"
	"github.com/wendongzhu/GPTS/backend/service/param"
	"io/ioutil"
)

func main() {
	path := "E:/workspace/go/src/GPTS/test/localfile/systemconfig.xml"
	sc := readXml(path)
	writeXml(path, sc)
	readXml(path)
}

func readXml(path string) param.SystemConfig {
	readData, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读文件出错！", err)

	}

	sc := param.SystemConfig{}
	err = xml.Unmarshal(readData, &sc)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println("DisplayParam : ", sc.DisplayParam)
	fmt.Println("SystemParam : ", sc.SystemParam)
	for i, item := range sc.DisplayParam.ProbePipeSetlList {
		fmt.Println("ProbePipeSetItem", i, item)
	}
	return sc
}

func writeXml(path string, sc param.SystemConfig) {
	sc.SystemParam.TimeInterval = 50
	sc.DisplayParam.TestCompany = "wuhan"
	b, _ := xml.MarshalIndent(sc, "", "	")
	b = append([]byte(xml.Header), b...)
	ioutil.WriteFile(path, b, 0666)
	fmt.Println("程序结束")

}
