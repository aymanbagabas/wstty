package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows"
)

var rootCmd = &cobra.Command{
	Use:  "wstty [-a][-][mode...]",
	RunE: cmd,
	Args: cobra.MinimumNArgs(1),
}

func cmd(cmd *cobra.Command, args []string) (err error) {
	for i, arg := range args {
		on := !strings.HasPrefix(arg, "-")
		arg := strings.TrimPrefix(arg, "-")
		switch arg {
		case "a":
			err = errors.Join(err, all(cmd))
		case "v", "version":
			cmd.Printf("wstty version %s\n", Version)
		case "h", "help":
			return cmd.Usage()
		case "raw":
			err = errors.Join(err, raw(on))
		case "cooked":
			err = errors.Join(err, cooked(on))
		case "sane":
			err = errors.Join(err, sane())
		case "cp", "icp", "ocp":
			if len(args) < i+2 {
				return fmt.Errorf("missing code page")
			}

			code, err := strconv.Atoi(args[i+1])
			if err != nil {
				return err
			}

			if arg == "cp" || strings.HasPrefix(arg, "i") {
				return cp(true, code)
			}
			if arg == "cp" || strings.HasPrefix(arg, "o") {
				return cp(false, code)
			}
		default:
			fn, ok := modes[arg]
			if !ok {
				return fmt.Errorf("unknown mode: %s", arg)
			}

			err = errors.Join(err, fn(on)(cmd, args))
		}
	}

	return
}

var modes = map[string]func(on bool) func(cmd *cobra.Command, args []string) error{
	"echo":     inputMode(windows.ENABLE_ECHO_INPUT),
	"line":     inputMode(windows.ENABLE_LINE_INPUT),
	"iproc":    inputMode(windows.ENABLE_PROCESSED_INPUT),
	"minput":   inputMode(windows.ENABLE_MOUSE_INPUT),
	"winput":   inputMode(windows.ENABLE_WINDOW_INPUT),
	"iinsert":  inputMode(windows.ENABLE_INSERT_MODE),
	"ivterm":   inputMode(windows.ENABLE_VIRTUAL_TERMINAL_INPUT),
	"iqedit":   inputMode(windows.ENABLE_QUICK_EDIT_MODE),
	"iautopos": inputMode(windows.ENABLE_AUTO_POSITION),
	"iext":     inputMode(windows.ENABLE_EXTENDED_FLAGS),
	"oproc":    outputMode(windows.ENABLE_PROCESSED_OUTPUT),
	"owrap":    outputMode(windows.ENABLE_WRAP_AT_EOL_OUTPUT),
	"ovterm":   outputMode(windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING),
	"onewline": outputMode(windows.DISABLE_NEWLINE_AUTO_RETURN),
	"olvb":     outputMode(windows.ENABLE_LVB_GRID_WORLDWIDE),
}

func raw(on bool) error {
	stdin, err := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
	if err != nil {
		return err
	}

	if on {
		var mode uint32
		if err := windows.GetConsoleMode(stdin, &mode); err != nil {
			return err
		}

		mode &^= windows.ENABLE_ECHO_INPUT | windows.ENABLE_LINE_INPUT | windows.ENABLE_PROCESSED_INPUT
		mode |= windows.ENABLE_VIRTUAL_TERMINAL_INPUT

		return windows.SetConsoleMode(stdin, mode)
	} else {
		return cooked(on)
	}
}

func cooked(on bool) error {
	stdin, err := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
	if err != nil {
		return err
	}

	if on {
		return raw(on)
	} else {
		var mode uint32
		if err := windows.GetConsoleMode(stdin, &mode); err != nil {
			return err
		}

		mode |= windows.ENABLE_ECHO_INPUT | windows.ENABLE_LINE_INPUT | windows.ENABLE_PROCESSED_INPUT
		mode &^= windows.ENABLE_VIRTUAL_TERMINAL_INPUT

		return windows.SetConsoleMode(stdin, mode)
	}
}

func inputMode(mode int) func(on bool) func(cmd *cobra.Command, args []string) error {
	return func(on bool) func(cmd *cobra.Command, args []string) error {
		return func(cmd *cobra.Command, args []string) error {
			stdin, err := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
			if err != nil {
				return err
			}

			var st uint32
			if err := windows.GetConsoleMode(stdin, &st); err != nil {
				return err
			}

			if on {
				st |= uint32(mode)
			} else {
				st &^= uint32(mode)
			}

			return windows.SetConsoleMode(stdin, st)
		}
	}
}

func outputMode(mode int) func(on bool) func(cmd *cobra.Command, args []string) error {
	return func(on bool) func(cmd *cobra.Command, args []string) error {
		return func(cmd *cobra.Command, args []string) error {
			stdout, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
			if err != nil {
				return err
			}

			var st uint32
			if err := windows.GetConsoleMode(stdout, &st); err != nil {
				return err
			}

			if on {
				st |= uint32(mode)
			} else {
				st &^= uint32(mode)
			}

			return windows.SetConsoleMode(stdout, st)
		}
	}
}

var imodes = []int{
	windows.ENABLE_ECHO_INPUT,
	windows.ENABLE_LINE_INPUT,
	windows.ENABLE_PROCESSED_INPUT,
	windows.ENABLE_MOUSE_INPUT,
	windows.ENABLE_WINDOW_INPUT,
	windows.ENABLE_INSERT_MODE,
	windows.ENABLE_QUICK_EDIT_MODE,
	windows.ENABLE_VIRTUAL_TERMINAL_INPUT,
	windows.ENABLE_AUTO_POSITION,
	windows.ENABLE_EXTENDED_FLAGS,
}

var imodestr = map[int]string{
	windows.ENABLE_ECHO_INPUT:             "echo",
	windows.ENABLE_LINE_INPUT:             "line",
	windows.ENABLE_PROCESSED_INPUT:        "iproc",
	windows.ENABLE_MOUSE_INPUT:            "minput",
	windows.ENABLE_WINDOW_INPUT:           "winput",
	windows.ENABLE_INSERT_MODE:            "iinsert",
	windows.ENABLE_QUICK_EDIT_MODE:        "iqedit",
	windows.ENABLE_VIRTUAL_TERMINAL_INPUT: "ivterm",
	windows.ENABLE_AUTO_POSITION:          "iautopos",
	windows.ENABLE_EXTENDED_FLAGS:         "iext",
}

var omodes = []int{
	windows.ENABLE_PROCESSED_OUTPUT,
	windows.ENABLE_WRAP_AT_EOL_OUTPUT,
	windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING,
	windows.DISABLE_NEWLINE_AUTO_RETURN,
	windows.ENABLE_LVB_GRID_WORLDWIDE,
}

var omodestr = map[int]string{
	windows.ENABLE_PROCESSED_OUTPUT:            "oproc",
	windows.ENABLE_WRAP_AT_EOL_OUTPUT:          "owrap",
	windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING: "ovterm",
	windows.DISABLE_NEWLINE_AUTO_RETURN:        "onewline",
	windows.ENABLE_LVB_GRID_WORLDWIDE:          "olvb",
}

func build(cmd *cobra.Command) error {
	windows.RtlGetVersion()
	major, minor, build := windows.RtlGetNtVersionNumbers()
	cmd.Printf("WindowsNT %d.%d.%d\n", major, minor, build)
	return nil
}

func all(cmd *cobra.Command) error {
	if err := build(cmd); err != nil {
		return err
	}

	icp, err := getcp(true)
	if err != nil {
		return err
	}

	ocp, err := getcp(false)
	if err != nil {
		return err
	}

	if icp == ocp {
		cmd.Printf("cp %d\n", icp)
	} else {
		cmd.Printf("icp %d; ocp %d\n", icp, ocp)
	}

	stdin, err := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
	if err != nil {
		return err
	}

	stdout, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		return err
	}

	var imode uint32
	if err := windows.GetConsoleMode(stdin, &imode); err != nil {
		return err
	}

	var omode uint32
	if err := windows.GetConsoleMode(stdout, &omode); err != nil {
		return err
	}

	type mode struct {
		mode  uint32
		modes []int
		strs  map[int]string
	}

	for _, m := range []mode{
		{imode, imodes, imodestr},
		{omode, omodes, omodestr},
	} {
		for i, mode := range m.modes {
			if i > 0 {
				cmd.Print(" ")
			}

			on := m.mode&uint32(mode) != 0
			str := m.strs[mode]
			if !on {
				cmd.Print("-")
			}

			cmd.Print(str)
		}

		cmd.Println()
	}

	return nil
}

func sane() error {
	stdin, err := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
	if err != nil {
		return err
	}

	stdout, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		return err
	}

	// Reset input mode
	// Enable all input modes except ENALBE_WINDOW_INPUT and ENABLE_VIRTUAL_TERMINAL_INPUT.
	if err := windows.SetConsoleMode(stdin, windows.ENABLE_PROCESSED_INPUT|windows.ENABLE_LINE_INPUT|windows.ENABLE_ECHO_INPUT|windows.ENABLE_MOUSE_INPUT|windows.ENABLE_INSERT_MODE|windows.ENABLE_QUICK_EDIT_MODE|windows.ENABLE_AUTO_POSITION|windows.ENABLE_EXTENDED_FLAGS); err != nil {
		return err
	}

	// Reset output mode
	// Enable all output modes except ENABLE_WRAP_AT_EOL_OUTPUT and ENABLE_LVB_GRID_WORLDWIDE.
	if err := windows.SetConsoleMode(stdout, windows.ENABLE_PROCESSED_OUTPUT|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING|windows.ENABLE_WRAP_AT_EOL_OUTPUT); err != nil {
		return err
	}

	return nil
}

func getcp(input bool) (cp uint32, err error) {
	var getter func() (uint32, error)
	if input {
		getter = windows.GetConsoleCP
	} else {
		getter = windows.GetConsoleOutputCP
	}

	cp, err = getter()
	return
}

func cp(input bool, code int) error {
	var setter func(code uint32) error
	if input {
		setter = windows.SetConsoleCP
	} else {
		setter = windows.SetConsoleOutputCP
	}

	return setter(uint32(code))
}
