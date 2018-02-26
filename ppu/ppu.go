package ppu

// PPU represents the Picture Processing Unit of the SNES
type PPU struct {
	vram  [0x10000]byte      // vram represents the VideoRAM (64KB)
	oam   [0x200 + 0x20]byte // oam represents the object attribute memory (512 + 32 Bytes)
	cgram [0x200]byte        // cgram represents the color graphics ram and stores the color palette with 256 color entries
}