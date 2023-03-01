package endinterface

import (
	"github.com/wendongzhu/GPTS/backend/service/param"
	"github.com/wendongzhu/GPTS/backend/service/tool"
)

func (b *Backend) DisplayAPI() (res param.DisplayParam) {
	path := "E:/workspace/go/src/GPTS/test/localfile/systemconfig.xml"
	ft := tool.FileTool{}

	res = ft.ReadXmlFile(path).DisplayParam
	return res
}

func (b *Backend) TestInfoAPI() (res param.TestInfo) {

	return res
}

func (b *Backend) SystemAPI() (res param.SystemParam) {
	path := "E:/workspace/go/src/GPTS/test/localfile/systemconfig.xml"
	ft := tool.FileTool{}

	res = ft.ReadXmlFile(path).SystemParam
	return res
}
