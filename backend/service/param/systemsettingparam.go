package param

type DisplayParam struct {
	TestUser          string             `xml:"TestUser"`
	TestAddress       string             `xml:"TestAddress"`
	TestCompany       string             `xml:"TestCompany"`
	ProbePipeSetlList []ProbePipeSetItem `xml:"ProbePipeSetItem"`
	TestDescription   string             `xml:"TestDescription"`
}

type TestInfo struct {
}

type SystemParam struct {
	PortName     string `xml:"PortName"`
	TimeInterval int32  `xml:"TimeInterval"`
	FilePrefix   string `xml:"FilePrefix"`
	WorkPath     string `xml:"WorkPath"`
}

type ProbePipeSetItem struct {
	ChannelName      string `xml:"ChannelName"`
	PowerStatus      string `xml:"PowerStatus"`
	VoltageThreshold int32  `xml:"VoltageThreshold"`
	CurrentThreshold int32  `xml:"CurrentThreshold"`
	PipeNumber       int32  `xml:"PipeNumber"`
}
