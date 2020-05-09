package main

import (
	"fmt"
	"time"

	"github.com/amanchourasiya/learning/basics"
)

func main() {
	variablesDemoRun()
	//loopsDemo()
	//functionDemoRun()
	//methodDemo()
	//interfaceDemo()
	//concurrencyDemo()
	//channelDemo()
}

func variablesDemoRun() {
	basics.VariablesDemo()

	//name := basics.Name{}
	//name.firstName = "aman"
	//name.LastName = "chourasiya"
}

func loopsDemo() {
	basics.LoopsDemo()
}

func functionDemoRun() {
	result := basics.FunctionDemo1(10, 23)
	fmt.Printf("FunctionDemo1: %d\n", result)

	//result1, isOk := basics.FunctionDemo2(11, 223)

}

func methodaDemo() {
	name := &basics.FullName{}
	name.SetFirstName()

}

type SBC7k struct {
	isCloud bool
}

func (sbc *SBC7k) GetSbcType() string {
	if sbc.isCloud == true {
		return "cloud"
	} else {
		return "H/W"
	}
}

func SBCDeployment(sbc basics.SbcType) {
	sbc_type := sbc.GetSbcType()
	fmt.Printf("Sbc type: %s\n", sbc_type)
}

func interfaceDemo() {
	swe := &SBC7k{
		isCloud: false,
	}
	SBCDeployment(swe)

}

// Concurrency Demo

func concurrencyDemo() {
	basics.ConcurrencyDemo()
	time.Sleep(20 * time.Second)
}

// Channel Demo

func channelDemo() {
	basics.ChannelDemo()
}
