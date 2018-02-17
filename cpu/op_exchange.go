package cpu

func (cpu *CPU) opEB() {
	temp := cpu.getBRegister()
	cpu.setBRegister(cpu.getARegister())
	cpu.setARegister(temp)
	cpu.nFlag = temp&0x80 != 0
	cpu.zFlag = temp == 0
	cpu.cycles += 3
	cpu.PC++
}

func (cpu *CPU) opFB() {
	cpu.eFlag, cpu.cFlag = cpu.cFlag, cpu.eFlag
	cpu.cycles += 2
	cpu.PC++
}
