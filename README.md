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

### Custom keybindings

You can create custom keybindings by creating the `$HOME/.keydroid`
directory and creating a `config.json` file.  The format is as
follows:

```json
{
	"keybindings": {
		"\b":   "KEYCODE_BACK",
		"\t":   "KEYCODE_TAB",
		"\r":   "KEYCODE_ENTER",
		" ":    "KEYCODE_SPACE",
		"-":    "KEYCODE_VOLUME_DOWN",
		"=":    "KEYCODE_VOLUME_UP",
		"+":    "KEYCODE_VOLUME_UP",
		"1":    "KEYCODE_TV_INPUT_HDMI_1",
		"2":    "KEYCODE_TV_INPUT_HDMI_2",
		"3":    "KEYCODE_TV_INPUT_HDMI_3",
		"4":    "KEYCODE_TV_INPUT_HDMI_4",
		"a":    "KEYCODE_ASSIST",
		"b":    "KEYCODE_BACK",
		"c":    "KEYCODE_DPAD_CENTER",
		"f":    "KEYCODE_MEDIA_FAST_FORWARD",
		"h":    "KEYCODE_DPAD_LEFT",
		"i":    "KEYCODE_TV_INPUT",
		"j":    "KEYCODE_DPAD_DOWN",
		"k":    "KEYCODE_DPAD_UP",
		"l":    "KEYCODE_DPAD_RIGHT",
		"m":    "KEYCODE_MUTE",
		"o":    "KEYCODE_HOME",
		"r":    "KEYCODE_MEDIA_REWIND",
		"s":    "KEYCODE_MEDIA_STOP",
		"t":    "KEYCODE_TV_TIMER_PROGRAMMING",
		"u":    "KEYCODE_MENU",
		"v":    "KEYCODE_VOICE_ASSIST",
		"w":    "KEYCODE_WAKEUP"
	}
}
```

All `KEYCODE` names listed at the [KeyEvent page](https://developer.android.com/reference/android/view/KeyEvent) are supported.

## Contributors

[stephenwithav](https://github.com/stephenwithav) - Steven Edwards
