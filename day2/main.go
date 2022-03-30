package main

import (
	"aoc/helper"
	"fmt"
	"strconv"
	"strings"
)

type pos struct {
	hor   int
	depth int
	aim   int
}

type instruction struct {
	forward, delta int
}

func convInputToInstruction(inputs []string) (insts []instruction) {
	for _, input := range inputs {
		fields := strings.Split(input, " ")
		delta, _ := strconv.Atoi(fields[1])
		inst := instruction{}
		switch fields[0] {
		case "up":
			inst.delta = -delta
		case "down":
			inst.delta = delta
		case "forward":
			inst.forward = delta
		default:
			panic("Could not recognize command: " + fields[0])
		}
		insts = append(insts, inst)
	}
	return insts
}

func simpleControl(instructions []instruction) {
	pos := pos{
		hor:   0,
		depth: 0,
	}

	for _, instruction := range instructions {
		pos.hor += instruction.forward
		pos.depth += instruction.delta
	}
	fmt.Printf("Using simple controls final position is (%v, %v)\n", pos.hor, pos.depth)
	fmt.Printf("Multiplied that gives %v\n", pos.depth*pos.hor)
}

func complexControl(instructions []instruction) {
	pos := pos{
		hor:   0,
		depth: 0,
		aim:   0,
	}

	for _, instruction := range instructions {
		pos.hor += instruction.forward
		pos.aim += instruction.delta
		pos.depth += instruction.forward * pos.aim
	}

	fmt.Printf("Using complex control your final position is (%v, %v)\n", pos.hor, pos.depth)
	fmt.Printf("Multiplied that gives %v\n", pos.depth*pos.hor)
}

func main() {
	inputs := helper.ReadInput("input")
	instructions := convInputToInstruction(inputs)

	simpleControl(instructions)
	complexControl(instructions)
}
