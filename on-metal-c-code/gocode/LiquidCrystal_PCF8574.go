package gocode

import (
	"log"
	"time"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

// LCD commands
const (
	LCD_CLEARDISPLAY   = 0x01
	LCD_RETURNHOME     = 0x02
	LCD_ENTRYMODESET   = 0x04
	LCD_DISPLAYCONTROL = 0x08
	LCD_CURSORSHIFT    = 0x10
	LCD_FUNCTIONSET    = 0x20
	LCD_SETCGRAMADDR   = 0x40
	LCD_SETDDRAMADDR   = 0x80
)

// Flags for display entry mode
const (
	LCD_ENTRYRIGHT          = 0x00
	LCD_ENTRYLEFT           = 0x02
	LCD_ENTRYSHIFTINCREMENT = 0x01
	LCD_ENTRYSHIFTDECREMENT = 0x00
)

// Flags for display on/off control
const (
	LCD_DISPLAYON  = 0x04
	LCD_DISPLAYOFF = 0x00
	LCD_CURSORON   = 0x02
	LCD_CURSOROFF  = 0x00
	LCD_BLINKON    = 0x01
	LCD_BLINKOFF   = 0x00
)

// Flags for display/cursor shift
const (
	LCD_DISPLAYMOVE = 0x08
	LCD_CURSORMOVE  = 0x00
	LCD_MOVERIGHT   = 0x04
	LCD_MOVELEFT    = 0x00
)

// Flags for function set
const (
	LCD_8BITMODE = 0x10
	LCD_4BITMODE = 0x00
	LCD_2LINE    = 0x08
	LCD_1LINE    = 0x00
	LCD_5x10DOTS = 0x04
	LCD_5x8DOTS  = 0x00
)

// PCF8574 pin mapping
const (
	PCF_RS        = 0x01
	PCF_RW        = 0x02
	PCF_EN        = 0x04
	PCF_BACKLIGHT = 0x08
)

// RSMODE - Register Select Mode
const (
	RSMODE_CMD  = 0
	RSMODE_DATA = 1
)

// LiquidCrystal represents an LCD connected via PCF8574 I2C expander.
type LiquidCrystal struct {
	i2cAddr    uint16
	backlight  uint8
	bus        i2c.Bus
	displayCtl uint8
}

// NewLiquidCrystal creates a new LiquidCrystal instance.
func NewLiquidCrystal(addr uint16) *LiquidCrystal {
	return &LiquidCrystal{
		i2cAddr:   addr,
		backlight: PCF_BACKLIGHT,
	}
}

// Begin initializes the LCD.
func (lcd *LiquidCrystal) Begin(cols, rows uint8, charsize uint8) error {
	// Initialize I2C bus
	if _, err := host.Init(); err != nil {
		return err
	}
	bus, err := i2creg.Open("")
	if err != nil {
		return err
	}
	lcd.bus = bus

	// Set display function flags
	displayFunction := LCD_4BITMODE | LCD_1LINE | LCD_5x8DOTS
	if rows > 1 {
		displayFunction |= LCD_2LINE
	}
	if charsize != LCD_5x8DOTS && rows == 1 {
		displayFunction |= LCD_5x10DOTS
	}

	// Wait for more than 40 ms after Vcc rises to 2.7V
	time.Sleep(50 * time.Millisecond)

	// Set to 4-bit mode
	lcd.write4bits(0x03)
	time.Sleep(5 * time.Millisecond)
	lcd.write4bits(0x03)
	time.Sleep(5 * time.Millisecond)
	lcd.write4bits(0x03)
	time.Sleep(150 * time.Microsecond)
	lcd.write4bits(0x02)

	// Configure function set, display control, and entry mode
	lcd.command(byte(LCD_FUNCTIONSET | displayFunction))
	lcd.displayCtl = LCD_DISPLAYON | LCD_CURSOROFF | LCD_BLINKOFF
	lcd.Display()
	lcd.Clear()
	lcd.command(LCD_ENTRYMODESET | LCD_ENTRYLEFT | LCD_ENTRYSHIFTDECREMENT)

	return nil
}

// Clear clears the display and returns the cursor to the home position.
func (lcd *LiquidCrystal) Clear() {
	lcd.command(LCD_CLEARDISPLAY)
	time.Sleep(2 * time.Millisecond)
}

// Home returns the cursor to the home position.
func (lcd *LiquidCrystal) Home() {
	lcd.command(LCD_RETURNHOME)
	time.Sleep(2 * time.Millisecond)
}

// Display turns on the display.
func (lcd *LiquidCrystal) Display() {
	lcd.displayCtl |= LCD_DISPLAYON
	lcd.command(LCD_DISPLAYCONTROL | lcd.displayCtl)
}

// NoDisplay turns off the display.
func (lcd *LiquidCrystal) NoDisplay() {
	lcd.displayCtl &^= LCD_DISPLAYON
	lcd.command(LCD_DISPLAYCONTROL | lcd.displayCtl)
}

// Cursor turns on the cursor.
func (lcd *LiquidCrystal) Cursor() {
	lcd.displayCtl |= LCD_CURSORON
	lcd.command(LCD_DISPLAYCONTROL | lcd.displayCtl)
}

// NoCursor turns off the cursor.
func (lcd *LiquidCrystal) NoCursor() {
	lcd.displayCtl &^= LCD_CURSORON
	lcd.command(LCD_DISPLAYCONTROL | lcd.displayCtl)
}

// Blink turns on the blinking cursor.
func (lcd *LiquidCrystal) Blink() {
	lcd.displayCtl |= LCD_BLINKON
	lcd.command(LCD_DISPLAYCONTROL | lcd.displayCtl)
}

// NoBlink turns off the blinking cursor.
func (lcd *LiquidCrystal) NoBlink() {
	lcd.displayCtl &^= LCD_BLINKON
	lcd.command(LCD_DISPLAYCONTROL | lcd.displayCtl)
}

// SetCursor moves the cursor to the specified position.
func (lcd *LiquidCrystal) SetCursor(col, row uint8) {
	rowOffsets := []uint8{0x00, 0x40, 0x14, 0x54}
	if row > 1 {
		row = 1 // Limit row to 0 or 1
	}
	lcd.command(LCD_SETDDRAMADDR | (col + rowOffsets[row]))
}

// Write sends a character to the display.
func (lcd *LiquidCrystal) Write(char byte) {
	lcd.send(char, RSMODE_DATA)
}

// command sends an LCD command.
func (lcd *LiquidCrystal) command(value byte) {
	lcd.send(value, RSMODE_CMD)
}

// send sends data to the LCD.
func (lcd *LiquidCrystal) send(value byte, mode byte) {
	lcd.write4bits(mode | (value >> 4))
	lcd.write4bits(mode | (value & 0x0F))
}

// write4bits writes a nibble (4 bits) to the LCD.
func (lcd *LiquidCrystal) write4bits(value byte) {
	err := lcd.bus.Tx(lcd.i2cAddr, []byte{value | lcd.backlight}, nil)
	if err != nil {
		log.Fatalf("Failed to write to LCD: %v", err)
	}
	lcd.pulseEnable(value)
}

// pulseEnable pulses the enable pin.
func (lcd *LiquidCrystal) pulseEnable(value byte) {
	err := lcd.bus.Tx(lcd.i2cAddr, []byte{value | PCF_EN | lcd.backlight}, nil)
	if err != nil {
		log.Fatalf("Failed to pulse enable: %v", err)
	}
	time.Sleep(1 * time.Microsecond)
	err = lcd.bus.Tx(lcd.i2cAddr, []byte{value&^PCF_EN | lcd.backlight}, nil)
	if err != nil {
		log.Fatalf("Failed to pulse enable: %v", err)
	}
	time.Sleep(50 * time.Microsecond)
}

//func main() {
//	lcd := NewLiquidCrystal(0x27) // Set the I2C address of the PCF8574
//	err := lcd.Begin(16, 2, LCD_5x8DOTS)
//	if err != nil {
//		log.Fatalf("Failed to initialize LCD: %v", err)
//	}
//
//	lcd.Clear()
//	lcd.SetCursor(0, 0)
//	lcd.Write('H')
//	lcd.Write('e')
//	lcd.Write('l')
//	lcd.Write('l')
//	lcd.Write('o')
//	lcd.Write(',')
//	lcd.SetCursor(0, 1)
//	lcd.Write('W')
//	lcd.Write('o')
//	lcd.Write('r')
//	lcd.Write('l')
//	lcd.Write('d')
//	lcd.Write('!')
//}
