# SetTopPi
A "dumb" tv style interface for omxplayer pulling network streams running on a raspberry pi.

## Usage
On start this program will attempt to turn on the connected HDMI display via CEC then start playing the specified video source.
On shutdown the video source will be closed and the connected display automatically turned off.

## TODO
This software is incomplete and unusable as of right now but there are a few basic tasks that I have planned out.

1. Use the remote control's channel change buttons to select the currently playing channel.
2. Add a systemd service to auto launch on boot and turn off on shut down.
3. Provide prebuilt binaries and/or debs.
4. Better way of specifying the user to run as.
