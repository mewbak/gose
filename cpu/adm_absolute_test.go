package cpu

import (
	"testing"

	"github.com/snes-emu/gose/memory"
)

func TestAbsolute(t *testing.T) {
	memory := memory.New()
	memory.SetByteBank(0xFF, 0x7E, 0x00001)
	memory.SetByteBank(0xFF, 0x7E, 0x00002)

	testCases := []struct {
		cpu                    *CPU
		expectedHi, expectedLo uint32
	}{
		{
			cpu:        &CPU{K: 0x7E, DBR: 0x12, memory: memory},
			expectedHi: 0x130000,
			expectedLo: 0x12FFFF,
		},
	}

	for i, tc := range testCases {
		addressHi, addressLo := tc.cpu.admAbsoluteP()

		if addressHi != tc.expectedHi {
			t.Errorf("Test %v failed: %x %x\n", i, addressHi, tc.expectedHi)
		} else if addressLo != tc.expectedLo {
			t.Errorf("Test %v failed: %x %x\n", i, addressLo, tc.expectedLo)
		}
	}
}

func TestAbsoluteX(t *testing.T) {
	memory := memory.New()
	memory.SetByteBank(0xFE, 0x7E, 0x00001)
	memory.SetByteBank(0xFF, 0x7E, 0x00002)

	testCases := []struct {
		cpu                    *CPU
		expectedHi, expectedLo uint32
	}{
		{
			cpu:        &CPU{K: 0x7E, DBR: 0x12, X: 0x000A, memory: memory},
			expectedHi: 0x130009,
			expectedLo: 0x130008,
		},
	}

	for i, tc := range testCases {
		addressHi, addressLo := tc.cpu.admAbsoluteXP()

		if addressHi != tc.expectedHi {
			t.Errorf("Test %v failed: %x %x\n", i, addressHi, tc.expectedHi)
		} else if addressLo != tc.expectedLo {
			t.Errorf("Test %v failed: %x %x\n", i, addressLo, tc.expectedLo)
		}
	}
}

func TestPAbsolute(t *testing.T) {
	memory := memory.New()
	memory.SetByteBank(0xFE, 0x7E, 0x00001)
	memory.SetByteBank(0xFF, 0x7E, 0x00002)

	testCases := []struct {
		cpu      *CPU
		expected uint16
	}{
		{
			cpu:      &CPU{K: 0x7E, DBR: 0x12, X: 0x000A, memory: memory},
			expected: 0x0000,
		},
	}

	for i, tc := range testCases {
		address := tc.cpu.admPAbsoluteJ()

		if address != tc.expected {
			t.Errorf("Test %v failed: %x %x\n", i, address, tc.expected)
		}
	}
}

func TestPAbsoluteX(t *testing.T) {
	memory := memory.New()
	memory.SetByteBank(0xFE, 0x7E, 0x00001)
	memory.SetByteBank(0xFF, 0x7E, 0x00002)
	memory.SetByteBank(0x56, 0x7E, 0x0008)
	memory.SetByteBank(0x34, 0x7E, 0x0009)

	testCases := []struct {
		cpu      *CPU
		expected uint16
	}{
		{
			cpu:      &CPU{K: 0x7E, DBR: 0x12, X: 0x000A, memory: memory},
			expected: 0x3456,
		},
	}

	for i, tc := range testCases {
		address := tc.cpu.admPAbsoluteXJ()

		if address != tc.expected {
			t.Errorf("Test %v failed: %x %x\n", i, address, tc.expected)
		}
	}
}
