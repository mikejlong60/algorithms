package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
)

type InputWire struct {
	Id int
	//OutputJunctions       []*OutputWire //Not sure I need this yet. The complete list of junctions with every output wire for this input wire.
	OutputWirePreferences []*OutputWire //An array of OutputWire preferences. Every input wire must be in this list.
	//OutputRoute []*OutputWire //The last int is the actual destination output wire.  And it's preceeded in the list by the route to get there.
}

type OutputWire struct {
	Id int
	//InputWirePreferences []*InputWire //Not sure I need this yet. An array of InputWire preferences. Every input wire must be in this list. This is ordered the same way for every input wire, at least for now.
	InputJunctions []*InputWire
}

func (w InputWire) String() string {
	return fmt.Sprintf("Id:%v, OutputWirePreferences:%v", w.Id, w.OutputWirePreferences)
}
func (w OutputWire) String() string {
	return fmt.Sprintf("Id:%v, InputJunctions:%v", w.Id, w.InputJunctions)
}

func MakeSwitches(incompleteInputWires *linked_list.LinkedList[*InputWire]) []*OutputWire {
	fmt.Printf("Size of list:%v\n", linked_list.Len(incompleteInputWires))
	if linked_list.Len(incompleteInputWires) == 0 {
		return []*OutputWire{}
	}
	var allOutputWires = linked_list.Head(incompleteInputWires).OutputWirePreferences

	for incompleteInputWires != nil {
		iw := linked_list.Head(incompleteInputWires)
		for i, ow := range iw.OutputWirePreferences {
			if ow.InputJunctions[i] == nil {
				//iw.OutputJunctions[i] = ow
				ow.InputJunctions[i] = iw
			} else { //try earlier point in ow.InputJunctions array. If there are none try later point from current index(i)
				var foundEarlierOwSpot = false
				for j := i - 1; j > -1; j-- { //Work backwards until you find an empty junction point on output wire and put input wire there
					if ow.InputJunctions[j] == nil {
						ow.InputJunctions[j] = iw
						foundEarlierOwSpot = true
						break
					}
				}
				if !foundEarlierOwSpot {
					for j := i + 1; j < len(ow.InputJunctions); j++ { //Work forward until you find an empty junction point on output wire and put input wire there
						if ow.InputJunctions[j] == nil {
							ow.InputJunctions[j] = iw
							break
						}
					}
				}
			}
		} //end placing iw preferences on Outputwire
		incompleteInputWires, _ = linked_list.Tail(incompleteInputWires)
	} // end man for
	return allOutputWires
}
