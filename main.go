package main

import (
	"fmt"
	"time"

	"github.com/amanchourasiya/golearning/basics"
	"github.com/amanchourasiya/golearning/features"
)

func main() {
	fmt.Println("Starting Go main program")
	//variablesDemoRun()
	//scopeDemo()
	//loopsDemo()
	//switchDemo()
	//functionDemoRun()
	//methodDemo()
	//interfaceDemo()
	//concurrencyDemo()
	//channelDemo()
	//bufferChannelDemo()
	//compositionDemo()
	//selectDemo()
}

func variablesDemoRun() {
	basics.VariablesDemo()

	name := basics.Name{}
	name.FirstName = "aman"
	name.LastName = "chourasiya"
}

func scopeDemo() {
	basics.ScopeTest()
}

func loopsDemo() {
	basics.LoopsDemo()
}

func switchDemo() {
	basics.SwitchTest(12)
	basics.SwitchTest(12.22)
	basics.SwitchTest("Some random string")
}

func functionDemoRun() {
	result := basics.FunctionDemo1(10, 23)
	fmt.Printf("FunctionDemo1: %d\n", result)

	//result1, isOk := basics.FunctionDemo2(11, 223)

}

func methodDemo() {
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
	features.InterfaceDemo()

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

func bufferChannelDemo() {
	features.BufferChannelDemo()
}

func compositionDemo() {
	features.CompositionDemo()
}

func selectDemo() {
	features.SelectDemo()
}
