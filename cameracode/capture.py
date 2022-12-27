#!/usr/bin/python3

from picamera import PiCamera
from time import sleep
import datetime
import sys

# Initiate the camera module with pre-defined settings.
camera = PiCamera()
camera.resolution = (640, 480)
camera.framerate = 15

def capture_photo(file_capture, text):
    # Add date as timestamp on the generated files.
    camera.annotate_text = text
    # Capture an image as the thumbnail.
    sleep(2)
    camera.capture(file_capture)
    print("Image captured " + file_capture)
    
# Get the current date as the timestamp to generate unique file names.
date = datetime.datetime.now().strftime('%d.%m.%Y %H:%M:%S')
capture_img = sys.argv[1]
capture_photo(capture_img, date)

