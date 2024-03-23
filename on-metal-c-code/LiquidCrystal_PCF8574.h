#ifndef LiquidCrystal_PCF8574_h
#define LiquidCrystal_PCF8574_h

#include <inttypes.h>
#include "Print.h"

// commands
#define LCD_CLEARDISPLAY 0x01
#define LCD_RETURNHOME 0x02
#define LCD_ENTRYMODESET 0x04
#define LCD_DISPLAYCONTROL 0x08
#define LCD_CURSORSHIFT 0x10
#define LCD_FUNCTIONSET 0x20
#define LCD_SETCGRAMADDR 0x40
#define LCD_SETDDRAMADDR 0x80

// flags for display entry mode
#define LCD_ENTRYRIGHT 0x00
#define LCD_ENTRYLEFT 0x02
#define LCD_ENTRYSHIFTINCREMENT 0x01
#define LCD_ENTRYSHIFTDECREMENT 0x00

// flags for display on/off control
#define LCD_DISPLAYON 0x04
#define LCD_DISPLAYOFF 0x00
#define LCD_CURSORON 0x02
#define LCD_CURSOROFF 0x00
#define LCD_BLINKON 0x01
#define LCD_BLINKOFF 0x00

// flags for display/cursor shift
#define LCD_DISPLAYMOVE 0x08
#define LCD_CURSORMOVE 0x00
#define LCD_MOVERIGHT 0x04
#define LCD_MOVELEFT 0x00

// flags for function set
#define LCD_8BITMODE 0x10
#define LCD_4BITMODE 0x00
#define LCD_2LINE 0x08
#define LCD_1LINE 0x00
#define LCD_5x10DOTS 0x04
#define LCD_5x8DOTS 0x00

class LiquidCrystal_PCF8574 : public Print {
public:

  LiquidCrystal_PCF8574();
  LiquidCrystal_PCF8574(uint8_t addr);

  void begin(uint8_t cols, uint8_t rows, uint8_t charsize = LCD_5x8DOTS);
  void begin(uint8_t cols, uint8_t lines, uint8_t addr, uint8_t backlight);
  bool checkI2CConnection();

  void clear();
  void home();
  
  void noDisplay();
  void display();
  void noBlink();
  void blink();
  void noCursor();
  void cursor();
  void scrollDisplayLeft(int charsToScroll, int scrollSpeed);
  void scrollDisplayRight(int charsToScroll, int scrollSpeed);
  void leftToRight();
  void rightToLeft();
  void autoscroll();
  void noAutoscroll();

  void setBacklight(uint8_t brightness);
  
  void createChar(uint8_t, uint8_t[]);
  void setCursor(uint8_t col, uint8_t row);
  void selectLine(uint8_t line);

  virtual size_t write(uint8_t);
  using Print::write;

private:
  // low level functions
  void _command(uint8_t);
  void _send(uint8_t value, uint8_t mode);
  void _sendNibble(uint8_t halfByte, uint8_t mode);
  void _write2Wire(uint8_t halfByte, uint8_t mode, uint8_t enable);

// NEW:
  uint8_t _Addr;        ///< Wire Address of the LCD
  uint8_t _backlight;   ///< the backlight intensity 