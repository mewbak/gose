package core

import (
	"testing"

	"github.com/snes-emu/gose/io"
	"github.com/snes-emu/gose/render"
	"github.com/stretchr/testify/assert"
)

func TestOamLastWrittenAddr(t *testing.T) {
	// Set oam addr register to 0x104
	ppu := newPPU(&render.NoOpRenderer{}, io.NewRegisterFactory())
	ppu.oamaddl(0x04)
	ppu.oamaddh(0x1)

	// Write 4 bytes
	ppu.oamdata(1)
	ppu.oamdata(2)
	ppu.oamdata(7)
	ppu.oamdata(12)

	// Write 1 to 0x2103
	ppu.oamaddh(0x1)

	assert.EqualValues(t, 252+8, ppu.oam.lastWrittenAddr)
}

func TestOamWritesAndReads(t *testing.T) {
	// From: https://wiki.superfamicom.org/sprites
	// Pictorially: Start OAM filled with all zeros.
	// Write 1, read, read, Write 2, read, write 3
	// => OAM is 00 00 01 02 01 03, rather than 01 00 00 02 00 03 as you might expect.

	ppu := newPPU(&render.NoOpRenderer{}, io.NewRegisterFactory())
	ppu.oamdata(0x1)
	assert.EqualValues(t, 0, ppu.rdoam())
	assert.EqualValues(t, 0, ppu.rdoam())
	ppu.oamdata(0x2)
	assert.EqualValues(t, 0, ppu.rdoam())
	ppu.oamdata(0x3)

	assert.EqualValues(t, []byte{0x00, 0x00, 0x01, 0x02, 0x01, 0x03}, ppu.oam.bytes[:6])
}
