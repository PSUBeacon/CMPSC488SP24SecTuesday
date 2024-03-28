from luma.led_matrix.device import max7219
from luma.core.interface.serial import spi, noop
from luma.core.render import canvas

# Create a serial interface with the correct GPIO pin numbers
serial = spi(port=0, device=0, gpio=noop())

# Create a device representing the LED matrix
device = max7219(serial, cascaded=1, block_orientation=90, rotate=0)

# Turn on all the LEDs in the matrix
with canvas(device) as draw:
    draw.rectangle(device.bounding_box, outline="white", fill="white")

# The LEDs will stay on until the program is terminated
