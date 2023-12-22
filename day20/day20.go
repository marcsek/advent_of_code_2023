package main

import (
	"fmt"
	"os"
	"strings"
)

type module interface {
	sendSignal() []module
	receiveSignal(sender module)
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

func (d *device) sendSignal() []module {
	for _, r := range d.dest {
		fmt.Printf("%v -%v-> %v\n", d.getId(), d.getState(), r.getId())
		r.receiveSignal(d)
	}
	return d.getDest()
}

func (d *device) receiveSignal(sender module) {
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
	return &flipFlop{device{id, false, []module{}}, false}
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

func (fl *flipFlop) receiveSignal(sender module) {
	if sender.getState() {
		fl.shouldPulse = false
		return
	}
	fl.shouldPulse = true
	fl.state = !fl.state
}

func (fl *flipFlop) sendSignal() []module {
	if !fl.shouldPulse {
		return []module{}
	}
	for _, r := range fl.dest {
		fmt.Printf("%v -%v-> %v\n", fl.getId(), fl.getState(), r.getId())
		r.receiveSignal(fl)
	}
	return fl.getDest()
}

func (cn *conjunction) receiveSignal(sender module) {
	cn.inputs[sender.getId()] = sender.getState()

	for _, value := range cn.inputs {
		if !value {
			cn.state = true
			return
		}
	}
	cn.state = false
}

// TODO: Prejst cez vseky hodnoty a najst ci sa zhoduju so zaciatkom.
func main() {
	input := strings.Split(readFile("sample.txt"), "\n")
	modules := parseInput(input[:len(input)-1])

	for btnPush := 0; btnPush < 5; btnPush++ {
		fmt.Printf("\n--  BUTTON PUSH %v --\n", btnPush+1)
		signalQueue := []module{modules["broadcaster"]}
		for len(signalQueue) > 0 {
			newQueue := []module{}

			for _, mod := range signalQueue {
				dest := mod.sendSignal()
				newQueue = append(newQueue, dest...)
			}
			signalQueue = newQueue
		}
	}
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
	mod.updateDest(newDest)
}

func updateInputs(conj *conjunction, mods map[string]module) {
	for key, value := range mods {
		for _, d := range value.getDest() {
			if d.getId() == conj.getId() {
				conj.inputs[key] = false
			}
		}
	}
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
}
