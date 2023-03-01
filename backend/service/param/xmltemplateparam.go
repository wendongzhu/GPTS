package param

type SystemConfig struct {
	Version      string       `xml:"version,attr"`
	DisplayParam DisplayParam `xml:"DisplayParam"`
	TestInfo     TestInfo     `xml:"TestInfo"`
	SystemParam  SystemParam  `xml:"SystemParam"`
}
