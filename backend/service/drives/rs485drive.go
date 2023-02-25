package drives

import (
	"fmt"
	_ "github.com/tarm/goserial"
	"github.com/tarm/serial"
)

type RS485 struct {
}

func (rs *RS485) Start() {
	//1.打开串口
	//打开串口的参数，name:串口名称，baud:波特率，size:数据位,parity:校验位，stopBit:停止位
	//例：串口名称：/dev/ttyS1 波特率9600，数据位8位，停止位1位，无校验
	cfg := &serial.Config{
		Name:     comName,
		Baud:     GetBaudrate(baudRate),
		Size:     GetDataBits(dataBit),
		Parity:   GetParityBits(checkBit),
		StopBits: GetStopBits(stopBit),
	}
	readPort, portErr := serial.OpenPort(cfg)
	if portErr != nil {
		fmt.printf("打开串口错误：%v", portErr)
		return
	}
	defer readPort //要关闭

	//2.生成读指令
	//根据实际情况，该函数就是根据具体的协议生成读指令
	readByteArr, generateErr := GenerateRead()
	if generateErr != nil {
		fmt.Printf("generate cmdByte err: %v", generateErr)
		return
	}

	//3.发送读指令
	wn, writeErr := readPort.Write(readByteArr)
	if writeErr != nil {
		fmt.Printf("write err: %v", writeErr)
		return

	}
	//4.接收并验证返回指令
	//计算返回命令的长度
	resLength := countResponseLength(cmdByteArr)

	//接收返回指令
	msg, msgErr := receiveData(iorwc, resLength)
	if msgErr != nil {
		fmt.Printf("receiveData err: %v", msgErr)
		return
	}

	//5.解析返回指令
	res, resBool := ParseData(msg)
	if !resBool {
		println("parse error")
		return
	}
	println("res data:", res)

}
