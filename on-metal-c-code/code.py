# set GPIO pin numbering method to BCM
import RPi.GPIO as GPIO
GPIO.setmode(GPIO.BCM)

#* define pins
#define DHT_PIN_DATA	4  #! Same pin diff machine please comment accordingly
#define BUZZER_PIN_SIG	4
#define KEYPADMEM3X4_PIN_ROW1	22
#define KEYPADMEM3X4_PIN_ROW2	23
#define KEYPADMEM3X4_PIN_ROW3	24
#define KEYPADMEM3X4_PIN_ROW4	25
#define KEYPADMEM3X4_PIN_COL1	17
#define KEYPADMEM3X4_PIN_COL2	18
#define KEYPADMEM3X4_PIN_COL3	27
#define LEDMATRIX_PIN_DIN	9
#define LEDMATRIX_PIN_CLK	10
#define LEDMATRIX_PIN_CS	4
#define PIR_PIN_SIG	11
#define PCFAN_PIN_COIL1	18
