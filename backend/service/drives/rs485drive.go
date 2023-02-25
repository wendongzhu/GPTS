package drives

import (
	"github.com/tarm/serial"
	"log"
)

type RS485 struct {
}

type SerialConfig struct {
	Name     string
	Baud     int
	Size     byte
	Parity   serial.Parity
	StopBits serial.StopBits
}

func (rs *RS485) Start() {

	//1.打开串口
	//打开串口的参数，name:串口名称，baud:波特率，size:数据位,parity:校验位，stopBit:停止位
	//例：串口名称：/dev/ttyS1 波特率9600，数据位8位，停止位1位，无校验
	// Name: "COM3", Baud: 115200, Size: 8, StopBits: 1
	sc := SerialConfig{}

	cfg := &serial.Config{
		Name:     sc.Name,
		Baud:     sc.Baud,
		Size:     sc.Size,
		Parity:   sc.Parity,
		StopBits: sc.StopBits,
	}

	//打开串口
	conn, err := serial.OpenPort(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() //要关闭

}

func (rs *RS485) handlerConnect() {}

func (rs *RS485) Write() {}
func (rs *RS485) Read()  {}
