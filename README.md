# wstty

<p>
    <a href="https://github.com/aymanbagabas/wstty/actions"><img src="https://github.com/aymanbagabas/wstty/workflows/build/badge.svg" alt="Build Status"></a>
    <a href="https://github.com/aymanbagabas/wstty/releases"><img src="https://img.shields.io/github/release/aymanbagabas/wstty.svg" alt="Latest Release"></a>
</p>

_wstty_ is a small tool that sets or reports various console related settings on
Windows systems. It is meant to be a replacement for the `stty` coreutils
command on Unix systems.

## Installation

### Go Install

```sh
go install github.com/aymanbagabas/wstty@latest
```

### Scoop (Windows)

```sh
scoop bucket add aymanbagabas https://github.com/aymanbagabas/scoop-bucket.git
scoop install aymanbagabas/wstty
```

## Usage

```shell
wstty [-a][-][mode...]
```

To set a specific mode to _off_, prefix the mode with a `-` sign. To set it to
_on_, just specify the mode.

Use `-a` to displays all the current console modes, `-h` to display the help.

### Input Modes

Input modes control how keyboard input is processed by the console.

| Mode       | Description                                                                                           |
| ---------- | ----------------------------------------------------------------------------------------------------- |
| `echo`     | Enables echo input. Characters typed by the user are echoed to the console.                           |
| `line`     | Enables line input mode. Input is not sent to the application until the Enter key is pressed.         |
| `iproc`    | Enables processed input. Special key combinations like Ctrl+C are processed by the system.            |
| `minput`   | Enables mouse input. Mouse events are reported to the application.                                    |
| `winput`   | Enables window input. Window resize events are reported to the application.                           |
| `iinsert`  | Enables insert mode. Text is inserted at the cursor position rather than overwriting existing text.   |
| `iqedit`   | Enables quick edit mode. Allows text selection with the mouse and clipboard operations.               |
| `ivterm`   | Enables virtual terminal input processing. Supports ANSI/VT100 escape sequences for input.            |
| `iautopos` | Enables automatic cursor positioning. The cursor is automatically positioned after a write operation. |
| `iext`     | Enables extended flags. Required for some extended console features.                                  |

### Output Modes

Output modes control how text and formatting are displayed on the console.

| Mode       | Description                                                                                                     |
| ---------- | --------------------------------------------------------------------------------------------------------------- |
| `oproc`    | Enables processed output. Backspace, tab, bell, carriage return, and line feed characters are processed.        |
| `owrap`    | Enables line wrapping. Text that extends beyond the right edge of the console window wraps to the next line.    |
| `ovterm`   | Enables virtual terminal processing. Supports ANSI/VT100 escape sequences for text formatting and colors.       |
| `onewline` | Disables automatic carriage return on line feed. Line feeds are not automatically preceded by carriage returns. |
| `olvb`     | Enables LVB grid worldwide. Supports line-drawing and box-drawing characters.                                   |

### Special Commands

| Command            | Description                                                                                                             |
| ------------------ | ----------------------------------------------------------------------------------------------------------------------- |
| `raw`              | Sets the console to raw mode. Disables echo, line input, and processed input, while enabling virtual terminal input.    |
| `cooked`           | Sets the console to cooked mode. Enables echo, line input, and processed input, while disabling virtual terminal input. |
| `sane`             | Resets the console to a reasonable default state.                                                                       |
| `cp <code>`        | Sets both input and output code pages to the specified code.                                                            |
| `icp <code>`       | Sets the input code page to the specified code.                                                                         |
| `ocp <code>`       | Sets the output code page to the specified code.                                                                        |
| `-a`               | Displays all current console modes and settings.                                                                        |
| `-v` or `-version` | Displays the version of wstty.                                                                                          |
| `-h` or `-help`    | Displays help information.                                                                                              |

### Examples

```shell
# Display all current console modes
wstty -a

# Enable virtual terminal processing for ANSI color support
wstty ovterm

# Disable line input mode (useful for character-by-character input)
wstty -line

# Set raw mode (useful for applications like vim)
wstty raw

# Set cooked mode (standard terminal behavior)
wstty cooked

# Set UTF-8 code page (65001) for both input and output
wstty cp 65001

# Enable virtual terminal processing and disable automatic line wrapping
wstty ovterm -owrap
```

## Technical Details

### Input Mode Flags

These flags control how input is processed by the Windows console:

- `ENABLE_ECHO_INPUT` (`echo`): When enabled, characters typed by the user are echoed to the console. When disabled, input characters are not displayed.
- `ENABLE_LINE_INPUT` (`line`): When enabled, input is not sent to the application until the Enter key is pressed. When disabled, input is sent immediately after each keystroke.
- `ENABLE_PROCESSED_INPUT` (`iproc`): When enabled, Ctrl+C is processed by the system and generates a SIGINT signal. When disabled, Ctrl+C is read as a regular character.
- `ENABLE_MOUSE_INPUT` (`minput`): When enabled, mouse events are reported in the input buffer. When disabled, mouse events are ignored.
- `ENABLE_WINDOW_INPUT` (`winput`): When enabled, window resize events are reported in the input buffer. When disabled, window resize events are ignored.
- `ENABLE_INSERT_MODE` (`iinsert`): When enabled, text is inserted at the cursor position. When disabled, text overwrites existing characters.
- `ENABLE_QUICK_EDIT_MODE` (`iqedit`): When enabled, the mouse can be used to select text and perform clipboard operations. When disabled, mouse input is sent to the application.
- `ENABLE_VIRTUAL_TERMINAL_INPUT` (`ivterm`): When enabled, ANSI/VT100 escape sequences for input are supported. When disabled, these sequences are treated as regular characters.
- `ENABLE_AUTO_POSITION` (`iautopos`): When enabled, the cursor is automatically positioned after a write operation. When disabled, the cursor position must be set explicitly.
- `ENABLE_EXTENDED_FLAGS` (`iext`): When enabled, the console can use extended features. This flag is required for some extended console features.

### Output Mode Flags

These flags control how output is processed by the Windows console:

- `ENABLE_PROCESSED_OUTPUT` (`oproc`): When enabled, backspace, tab, bell, carriage return, and line feed characters are processed. When disabled, these characters are displayed as is.
- `ENABLE_WRAP_AT_EOL_OUTPUT` (`owrap`): When enabled, text that extends beyond the right edge of the console window wraps to the next line. When disabled, text is truncated at the edge of the window.
- `ENABLE_VIRTUAL_TERMINAL_PROCESSING` (`ovterm`): When enabled, ANSI/VT100 escape sequences for text formatting and colors are supported. When disabled, these sequences are displayed as is.
- `DISABLE_NEWLINE_AUTO_RETURN` (`onewline`): When enabled, line feeds are not automatically preceded by carriage returns. When disabled, a carriage return is automatically inserted before each line feed.
- `ENABLE_LVB_GRID_WORLDWIDE` (`olvb`): When enabled, line-drawing and box-drawing characters are supported. When disabled, these characters may not display correctly.

See [Console Modes](https://docs.microsoft.com/en-us/windows/console/console-modes) for more information.

## Notes

- Some modes may not be available on all Windows versions.
- Changing console modes affects the behavior of all applications using the console.
- The `-a` option is useful for debugging console issues.
- The `sane` command is useful for resetting the console to a known good state.
- Code page 65001 corresponds to UTF-8 encoding.
- Common code pages include:
  - 437: US English
  - 850: Western European
  - 932: Japanese
  - 936: Simplified Chinese
  - 949: Korean
  - 950: Traditional Chinese
  - 1200: UTF-16LE
  - 1201: UTF-16BE
  - 65001: UTF-8

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file

