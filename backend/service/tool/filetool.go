package tool

import (
	"encoding/xml"
	"fmt"
	"github.com/wendongzhu/GPTS/backend/service/param"
	"io/ioutil"
)

type FileTool struct {
}

func (f *FileTool) ReadXmlFile(path string) (sc param.SystemConfig) {
	readData, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读文件出错！", err)
	}

	err = xml.Unmarshal(readData, &sc)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return sc
}

func (f *FileTool) WriteXmlFile(path string, sc param.SystemConfig) {
	b, _ := xml.MarshalIndent(sc, "", "	")
	b = append([]byte(xml.Header), b...)
	ioutil.WriteFile(path, b, 0666)
	fmt.Println("xml 写入完成")
}
