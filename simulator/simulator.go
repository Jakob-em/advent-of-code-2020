package simulator

import (
	"strconv"
	"strings"
)

type Hook func(simulator Simulator, instruction Instruction) bool

type Simulator struct {
	InstructionCounter   int
	Accumulator          int
	Code                 []Instruction
	preExecution         Hook
	executedInstructions map[int]bool
}

type Instruction struct {
	Operation string
	Argument  int
}

func New(code []string, preExecutionHook Hook) (Simulator, error) {
	instructions, err := parseCode(code)
	if err != nil {
		return Simulator{}, err
	}

	return Simulator{
		InstructionCounter:   0,
		Accumulator:          0,
		Code:                 instructions,
		preExecution:         preExecutionHook,
		executedInstructions: map[int]bool{},
	}, nil
}

func parseCode(code []string) ([]Instruction, error) {
	var instructions []Instruction
	for _, line := range code {
		instruction, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, instruction)
	}
	return instructions, nil
}

func parseInstruction(line string) (Instruction, error) {
	cmd := strings.Split(line, " ")
	arg, err := strconv.Atoi(cmd[1])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{
		Operation: cmd[0],
		Argument:  arg,
	}, nil
}

func (s *Simulator) Run() {
	for ; s.InstructionCounter < len(s.Code); {
		_, alreadyExecuted := s.executedInstructions[s.InstructionCounter]
		s.executedInstructions[s.InstructionCounter] = true
		if alreadyExecuted {
			return
		}

		instruction := s.Code[s.InstructionCounter]

		if s.preExecution != nil {
			stop := s.preExecution(*s, instruction)
			if stop {
				return
			}
		}

		// pre execute could have modified the code
		instruction = s.Code[s.InstructionCounter]
		s.executeInstruction(instruction)
	}

}

func (s *Simulator) IsAtEnd() bool {
	return s.InstructionCounter == len(s.Code)
}

func (s *Simulator) executeInstruction(instruction Instruction) {
	if instruction.Operation == "acc" {
		s.Accumulator += instruction.Argument
	}

	if instruction.Operation == "jmp" {
		s.InstructionCounter += instruction.Argument

	} else {
		s.InstructionCounter++

	}

}
