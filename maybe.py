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

def set_bit(value, bit):
    """Set the bit at `bit` position to 1 in `value`."""
    return value | (1 << bit)

def draw_smile():
    clear_matrix()  # Clear the matrix to start fresh
    smile_bmp = [
        0b00011110,
        0b00100001,
        0b11010010,
        0b11000000,
        0b11010010,
        0b11001100,
        0b00100001,
        0b00011110
    ]

    # Iterating over each row in the smile_bmp array
    for i, row_data in enumerate(smile_bmp, start=1):
            send_command(i, row_data)

    sleep(0.33)  # Display the smiley face for a short period
def prepare_and_draw(byte_sequence):
    # Step 1: Reverse the byte order
    reversed_bytes = byte_sequence[::-1]

    # Step 2: Convert from LSB to MSB for each byte
    def reverse_bits(byte):
        return int('{:08b}'.format(byte)[::-1], 2)
    converted_bytes = bytes(reverse_bits(b) for b in reversed_bytes)

    # Step 3: Draw the prepared bytes on the matrix
    clear_matrix()  # Ensure a clean slate
    for i, byte in enumerate(converted_bytes):
        send_command(i + 1, byte)
    sleep(0.33)  # Display for a short duration

# Example byte sequence for a character (replace with your actual character bytes)
byte_sequence = b'\x7e\x18\x18\x18\x1c\x18\x18\x00'
prepare_and_draw(byte_sequence)
sleep(5)

# Initialize and clear the matrix
initialize_matrix()
clear_matrix()

# Fill the matrix
draw_smile()

# Wait for 10 seconds
sleep(10)

# Clear the matrix
clear_matrix()
