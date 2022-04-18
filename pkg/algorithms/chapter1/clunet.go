package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
)

type InputWire struct {
	Id                    int
	OutputWirePreferences []*OutputWire //An array of OutputWire preferences. Every input wire must be in this list.
}

type OutputWire struct {
	Id             int
	InputJunctions []*InputWire
}

func (w InputWire) String() string {
	var owIds = make([]int, len(w.OutputWirePreferences), len(w.OutputWirePreferences))
	for i, ow := range w.OutputWirePreferences {
		owIds[i] = ow.Id
	}
	return fmt.Sprintf("InputWire{Id:%v, OutputWirePreferences:%v}", w.Id, owIds)
}

func (w OutputWire) String() string {
	var iwIds = make([]int, len(w.InputJunctions), len(w.InputJunctions))
	for i, iw := range w.InputJunctions {
		iwIds[i] = iw.Id
	}
	return fmt.Sprintf("OutputWire{Id:%v, InputJunctions:%v}", w.Id, iwIds)
}

func MakeSwitches(incompleteInputWires *linked_list.LinkedList[*InputWire]) []*OutputWire {
	fmt.Printf("Size of list:%v\n", linked_list.Len(incompleteInputWires))
	if linked_list.Len(incompleteInputWires) == 0 {
		return []*OutputWire{}
	}
	var allOutputWires = linked_list.Head(incompleteInputWires).OutputWirePreferences

	for incompleteInputWires != nil {
		//fmt.Println(linked_list.Len(incompleteInputWires))
		iw := linked_list.Head(incompleteInputWires)
		//fmt.Println(iw.OutputWirePreferences)
		fmt.Println("==================")

		for i, ow := range iw.OutputWirePreferences {
			if ow.InputJunctions[i] == nil {
				ow.InputJunctions[i] = iw
				fmt.Printf("found empty spot for first InputWire:%v\n", iw)
				//fmt.Printf("all output wires for first InputWire:%v\n", allOutputWires)

			} else { //try earlier point in ow.InputJunctions array. If there are none try later point from current index(i)
				//Stick input wire at first empty spot on output wire
				var foundEmptySpotOnOw = false
				for j := 0; j < len(ow.InputJunctions); j++ {
					if ow.InputJunctions[j] == nil {
						ow.InputJunctions[j] = iw
						foundEmptySpotOnOw = true
						fmt.Printf("found empty spot for InputWire:%v\n", iw)
						//fmt.Printf("all output wires for after empty spot InputWire:%v\n", allOutputWires)
						break
					}
				}
				if !foundEmptySpotOnOw {
					fmt.Println("could not find empty spot")
				}
				//var foundEarlierOwSpot = false
				//for j := i; j >= 0; j-- { //Work backwards until you find an empty junction point on output wire and put input wire there
				//	if ow.InputJunctions[j] == nil {
				//		ow.InputJunctions[j] = iw
				//		foundEarlierOwSpot = true
				//		break
				//	}
				//}
				//if !foundEarlierOwSpot {
				//	for j := i; j < len(ow.InputJunctions); j++ { //Work forward until you find an empty junction point on output wire and put input wire there
				//		if ow.InputJunctions[j] == nil {
				//			ow.InputJunctions[j] = iw
				//			break
				//		}
				//	}
				//}
			}
		} //end placing iw preferences on Outputwire
		incompleteInputWires, _ = linked_list.Tail(incompleteInputWires)
	} // end InputWire for
	return allOutputWires
}
