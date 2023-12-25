package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type module interface {
	sendSignal() ([]module, int, int)
	receiveSignal(id string, state bool)
	getDest() []module
	updateDest(newDest []module)
	getId() string
	getState() bool
}

type device struct {
	id    string
	state bool
	dest  []module
}

type flipFlop struct {
	device
	shouldPulse bool
}

type conjunction struct {
	device
	inputs map[string]bool
}

type broadcaster struct {
	device
}

type output struct {
	device
}

func (d *device) sendSignal() ([]module, int, int) {
	h, l := 0, 0
	for _, r := range d.dest {
		fmt.Printf("%v -%v-> %v\n", d.getId(), d.getState(), r.getId())
		if d.state {
			h++
		} else {
			l++
		}
		if r.getState() {

		}

	}
	return d.getDest(), h, l
}

func (d *device) receiveSignal(id string, state bool) {
	return
}

func (d *device) getDest() []module {
	return d.dest
}

func (d *device) updateDest(newDest []module) {
	d.dest = newDest
	return
}

func (d *device) getId() string {
	return d.id
}

func (d *device) getState() bool {
	return d.state
}

func newFlipFlop(id string) *flipFlop {
	return &flipFlop{device{id, false, []module{}}, true}
}

func newConjunction(id string) *conjunction {
	return &conjunction{device{id, false, []module{}}, map[string]bool{}}
}

func newBroadCaster(id string) *broadcaster {
	return &broadcaster{device{id, false, []module{}}}
}

func newOutput(id string) *output {
	return &output{device{id, false, []module{}}}
}

func (fl *flipFlop) receiveSignal(id string, state bool) {
	if state {
		fl.shouldPulse = false
		return
	}
	fl.shouldPulse = true
	fl.state = !fl.state
}

func (fl *flipFlop) sendSignal() ([]module, int, int) {
	if !fl.shouldPulse {
		return []module{}, 0, 0
	}
	h, l := 0, 0
	for _, r := range fl.dest {
		fmt.Printf("%v -%v-> %v\n", fl.getId(), fl.getState(), r.getId())
		if fl.state {
			h++
		} else {
			l++
		}
		if r.getState() {

		}
	}
	return fl.getDest(), h, l
}

func (cn *conjunction) receiveSignal(id string, state bool) {
	cn.inputs[id] = state

	for _, value := range cn.inputs {
		if !value {
			cn.state = true
			return
		}
	}
	cn.state = false
}

type caller struct {
	id   string
	mods []module
}

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	modules := parseInput(input[:len(input)-1])
	savedState := map[string]bool{}
	fmt.Println(modules)

	tH, tL := 0, 0
	for btnPush := 0; btnPush < 1000; btnPush++ {
		fmt.Printf("\n--  BUTTON PUSH %v --\n", btnPush+1)
		signalQueue := []module{modules["broadcaster"]}
		toReceive := []caller{}
		newH, newL := 0, 0
		for len(signalQueue) > 0 {
			newQueue := []module{}

			state := map[string]bool{}
			for _, clr := range toReceive {
				fmt.Println(modules[clr.id])
				if found, ok := modules[clr.id]; ok {
					state[clr.id] = found.getState()
				} else {
					state[clr.id] = false
				}
			}
			//slices.Reverse(toReceive)
			slices.Reverse(toReceive)
			fmt.Println(state)
			for _, clr := range toReceive {
				for _, mod := range clr.mods {
					mod.receiveSignal(modules[clr.id].getId(), state[clr.id])
				}
			}
			toReceive = []caller{}
			for _, mod := range signalQueue {
				dest, h, l := mod.sendSignal()
				toReceive = append(toReceive, caller{mod.getId(), dest})
				newH += h
				newL += l
				newQueue = append(newQueue, dest...)
			}
			signalQueue = newQueue
		}
		for _, clr := range toReceive {
			for _, mod := range clr.mods {
				mod.receiveSignal(modules[clr.id].getId(), modules[clr.id].getState())
			}
		}
		//allMatch := true
		//for key, value := range savedState {
		//	if modules[key].getState() != value {
		//		allMatch = false
		//		break
		//	}
		//}
		//if btnPush != 0 && allMatch {
		//	fmt.Println("KONEC", tL, tH)
		//	fmt.Println("VYSLEDOK: ", (tL*(1_000/(btnPush)))*(tH*(1_000/(btnPush))))
		//	break
		//}
		if btnPush == 0 {
			for _, mod := range modules {
				savedState[mod.getId()] = mod.getState()
			}
		}
		tH += newH
		tL += newL + 1
	}
	fmt.Println(tH * tL)
}

func parseInput(input []string) map[string]module {
	modules := make(map[string]module, len(input))
	dests := make(map[string]string, len(input))

	for _, line := range input {
		sp := strings.Split(line, " -> ")
		modName, dest := sp[0], sp[1]
		newModule := *new(module)

		switch name := modName; name {
		case "broadcaster":
			newModule = newBroadCaster(modName)
		default:
			switch prefix := modName[0]; prefix {
			case '%':
				modName = modName[1:]
				newModule = newFlipFlop(modName)
			case '&':
				modName = modName[1:]
				newModule = newConjunction(modName)
			default:
				panic("Couldn't parse input file.")
			}
		}
		modules[modName] = newModule
		dests[modName] = dest
	}
	for key, value := range modules {
		updateOutputs(value, modules, strings.Split(dests[key], ", "))
	}
	for _, value := range modules {
		if c, ok := value.(*conjunction); ok {
			updateInputs(c, modules)
		}
	}
	return modules
}

func updateOutputs(mod module, mods map[string]module, outs []string) {
	newDest := []module{}
	for _, o := range outs {
		found, ok := mods[o]
		if !ok {
			newDest = append(newDest, newOutput(o))
			continue
		}
		newDest = append(newDest, found)
	}
	//fmt.Println("outputs", mod, newDest)
	mod.updateDest(newDest)
}

func updateInputs(conj *conjunction, mods map[string]module) {
	for key, value := range mods {
		//fmt.Println(key, value)
		for _, d := range value.getDest() {
			if d.getId() == conj.getId() && key != conj.getId() {
				conj.inputs[key] = false
			}
		}
	}
	fmt.Println(conj.inputs)
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
}
