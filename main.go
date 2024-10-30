package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "unknown"

const usage = `Usage: wstty [-a][-][mode...]

  -a            Show all modes and Windows build numbers
  raw           Put the terminal into raw mode
  -raw          Same as cooked
  cooked        Put the terminal into cooked mode
  -cooked       Same as raw
  sane          Restore the terminal to its default state

Code page:
  cp <code>     Set the code page for both input and output
  icp <code>	Set the code page for input
  ocp <code>	Set the code page for output

Input modes:
  [-]echo       Whether input is echoed (ENABLE_ECHO_INPUT)
  [-]line       Whether input is line buffered (ENABLE_LINE_INPUT)
  [-]minput     Whether mouse events are reported (ENABLE_MOUSE_INPUT)
  [-]winput     Whether window resize events are reported (ENABLE_WINDOW_INPUT)
  [-]iproc      Whether input is processed (ENABLE_PROCESSED_INPUT)
  [-]iqedit     Whether mouse is used to select and edit (ENABLE_QUICK_EDIT_MODE)
  [-]iinsert    Whether insert mode is enabled (ENABLE_INSERT_MODE)
  [-]ivterm     Whether virtual terminal input is enabled (ENABLE_VIRTUAL_TERMINAL_INPUT)
  [-]iautopos   (ENABLE_AUTO_POSITION)
  [-]iext       (ENABLE_EXTENDED_FLAGS)

Output modes:
  [-]oproc      Whether output is processed (ENABLE_PROCESSED_OUTPUT)
  [-]owrap      Whether output is wrapped at the end of the line (ENABLE_WRAP_AT_EOL_OUTPUT)
  [-]ovterm     Whether virtual terminal output is enabled (ENABLE_VIRTUAL_TERMINAL_PROCESSING)
  [-]onewline   Whether newline auto return is disabled (DISABLE_NEWLINE_AUTO_RETURN)
  [-]olvb       (ENABLE_LVB_GRID_WORLDWIDE)
`

func init() {
	rootCmd.Version = Version
	rootCmd.DisableFlagParsing = true
	rootCmd.DisableFlagsInUseLine = true
	rootCmd.SetUsageTemplate(usage)
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
