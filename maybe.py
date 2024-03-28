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

# Initialize and clear the matrix
initialize_matrix()
clear_matrix()

# Fill the matrix
fill_matrix()

# Wait for 5 seconds
sleep(5)

# Clear the matrix
clear_matrix()
