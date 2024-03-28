from gpiozero import DigitalOutputDevice
from time import sleep

# Define the GPIO pins
CS_PIN = 4
DIN_PIN = 9
CLK_PIN = 10

# Create DigitalOutputDevice objects for each pin
cs = DigitalOutputDevice(CS_PIN)
din = DigitalOutputDevice(DIN_PIN)
clk = DigitalOutputDevice(CLK_PIN)

def send_byte(data):
    for i in range(8):
        # Set DIN to the value of the most significant bit
        din.value = (data & 0x80) != 0
        # Pulse the clock
        clk.on()
        sleep(0.0001)
        clk.off()
        # Shift the data left by one bit
        data <<= 1

def send_command(register, data):
    # Pull CS low to select the device
    cs.off()
    # Send the register address and data
    send_byte(register)
    send_byte(data)
    # Pull CS high to release the device
    cs.on()

def initialize_matrix():
    # Set the number of digits to 8
    send_command(0x0B, 0x07)
    # Set the intensity (brightness) to a medium value
    send_command(0x0A, 0x04)
    # Exit shutdown mode
    send_command(0x0C, 0x01)
    # Disable display test
    send_command(0x0F, 0x00)

def clear_matrix():
    for i in range(1, 9):
        send_command(i, 0x00)

def fill_matrix():
    for i in range(1, 9):
        send_command(i, 0xFF)
def draw_H():
    # Clear the matrix first to start fresh
    clear_matrix()

    # Pattern for an "H"
    # Rows 1-3 and 5-8: 0b10000001
    # Row 4 (middle connector): 0b11111111
    h_pattern = [0b10000001, 0b10000001, 0b10000001, 0b11111111,
                 0b10000001, 0b10000001, 0b10000001, 0b10000001]

    # Send each row's pattern to the matrix
    for i, pattern in enumerate(h_pattern, start=1):
        send_command(i, pattern)

# Initialize and clear the matrix
initialize_matrix()
clear_matrix()

# Fill the matrix
draw_H()

# Wait for 5 seconds
sleep(5)

# Clear the matrix
clear_matrix()
