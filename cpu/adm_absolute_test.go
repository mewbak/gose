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
