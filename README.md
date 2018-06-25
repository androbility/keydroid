# Introduction

KeyDroid lets you interact with your Android device from the terminal
using adb, the Android Debug Bridge.

It enables control of your Android devices, like the Fire TV Player
and Google's Nexus player, from your cli.  A decent analogy is:

    kodikeys:kodi::keydroid:android

In contrast to kodikeys, which lets you communicate with kodi,
keydroid lets you control the entire Android system; it also attempts
to reconnect when the adb connection is lost.

pre-1.0, the initial keybindings will be in flux and all debug
messages will be shown.  At 0.4, all keybindings should be
configurable.

On a final note, I would like to create an organiation to house this
project and all of the keymaps for different devices.  The keydroid
account already exists on GitHub, so I am open to suggestions for new
names.

I created [TightPants](https://github.com/tightpants) as a Firefly reference, with a tp subrepo, but
`tp` as a command to supervise connections to Android devices?
Something better must exist.  If you can come up with a good
alternative, I will link you (or your company/organization) in this
README file forever.


## Installation

```bash
$ go get github.com/stephenwithav/keydroid
$ glide update
$ go install github.com/stephenwithav/keydroid
```


## Requirements

Install [adb](https://developer.android.com/studio/command-line/adb).


## Keybindings

The keybindinds that should work on all Android devices are:

```
<bcksp> - Send Backspace key.
<tab>   - Send Tab key.
<space> - Select item; also, pause.
<enter> - Select item; also, pause.
      - - Turn volume down.
      = - Turn volume up.
      + - Turn volume up.
      h - Move selection left.
      j - Move selection down.
      k - Move selection up.
      l - Move selection right.
      s - Stop.
```


The full, non-final list of default keybindings is:

```
<bcksp> - Send Backspace key.
<tab>   - Send Tab key.
<space> - Select item; also, pause.
<enter> - Select item; also, pause.
      - - Turn volume down.
      = - Turn volume up.
      + - Turn volume up.
      1 - Switch TV input to HDMI1.
      2 - Switch TV input to HDMI2.
      3 - Switch TV input to HDMI3.
      4 - Switch TV input to HDMI4.
      a - Trigger voice assistant.
      b - Return to prevous screen.
      c - Press center button.
      f - Fast-forward.
      h - Move selection left.
      i - Change TV inputs.
      j - Move selection down.
      k - Move selection up.
      l - Move selection right.
      m - Mute volume.
      o - Return to Home screen.
      q - Quit KeyDroid.
      r - Rewind.
      s - Stop.
      t - TV timer button.
      u - Menu.
      v - Voice assistant.
      w - Wakeup Android device.
```


## Contributors

[stephenwithav](https://github.com/stephenwithav) - Steven Edwards
