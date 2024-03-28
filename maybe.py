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

def draw_lightbulb():
    # Clear the matrix first to start with a blank slate
    clear_matrix()
    # Define the pattern for a lightbulb
    lightbulb_pattern = [
        0b00111100,  # Row 1: Top of the bulb
        0b01111110,  # Row 2
        0b01111110,  # Row 3
        0b01111110,  # Row 4: Middle of the bulb (widest)
        0b01111110,  # Row 5: Start of the base
        0b00111100,  # Row 6: Narrower part of the base
        0b00011000,  # Row 7: Bottom of the base
        0b00011000   # Row 8: Optional, to give the base a more defined look
    ]

    # Send each row's pattern to the matrix
    for row, pattern in enumerate(lightbulb_pattern, start=1):
        send_command(row, pattern)

def draw_lock():
    # Clear the matrix first to start with a blank slate
    clear_matrix()
    # Define the pattern for a lock
    lock_pattern = [
        0b00111100,  # Row 1: Top empty space
        0b00111100,  # Row 2: Top part of the shackle
        0b01111110,  # Row 3: Bottom part of the shackle
        0b11111111,  # Row 4: Top of the lock body
        0b11000011,  # Row 5: Middle of the lock body, widest part
        0b11000011,  # Row 6: Lower part of the lock body
        0b11000011,  # Row 7: Bottom of the lock body
        0b01111110   # Row 8: Base of the lock, gives it a rounded look
    ]

    # Send each row's pattern to the matrix
    for row, pattern in enumerate(lock_pattern, start=1):
        send_command(row, pattern)

def draw_H():
    # Clear the matrix first to start with a blank slate
    clear_matrix()
    # Define the pattern for an "H"
    # Rows 1-3 and 5-8: 0b10000001 (vertical lines)
    # Row 4: 0b11111111 (horizontal line connecting the vertical lines)
    h_pattern = [0b10000001, 0b10000001, 0b10000001, 0b11111111,
                 0b10000001, 0b10000001, 0b10000001, 0b10000001]

    # Send each row's pattern to the matrix
    for row, pattern in enumerate(h_pattern, start=1):
        send_command(row, pattern)

# Initialize and clear the matrix
initialize_matrix()
clear_matrix()

# Fill the matrix
#draw_H()

# Wait for 10 seconds
#sleep(10)
draw_lock()
sleep(30)
# Clear the matrix
clear_matrix()
