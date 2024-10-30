# wstty

*wstty* is a small tool that sets or reports various console related settings on
*Windows systems. It is meant to be a replacement for the `stty` coreutils
*command on Unix systems.


## Usage

```shell
wstty [-a][-][mode...]
```

To set a specific mode to _off_, prefix the mode with a `-` sign. To set it to
_on_, just specify the mode.

## Options

### **-a**

Displays all the current console modes. If the mode is prefixed with a `-`, it
means that the mode is _off_.

## Console Modes

### **build**

Displays the Windows NT build number.

### **raw**

Sets the console to raw mode. This is equivalent to setting the
`ENABLE_LINE_INPUT`, `ENABLE_ECHO_INPUT`, and `ENABLE_PROCESSED_INPUT` flags
to _off_. And setting `ENABLE_VIRTUAL_TERMINAL_INPUT` flag to _on_.

### **columns**

Sets the number of columns in the console window. Optionally, you can specify
a second argument separated by a comma to set the number of buffer lines.

### **rows**

Sets the number of rows in the console window. Optionally, you can specify
a second argument separated by a comma to set the number of buffer lines.

### **cp**

Sets both input and output code pages for the console window. This is equivalent
to both `icp` and `ocp` modes.

### **icp**

Sets the input code page for the console window.

### **ocp**

Sets the output code page for the console window.

### **echo**

This corresponds to the `ENABLE_ECHO_INPUT` flag.

### **iproc**

This corresponds to the `ENABLE_PROCESSED_INPUT` flag.

### **line**

This corresponds to the `ENABLE_LINE_INPUT` flag.

### **minput**

This corresponds to the `ENABLE_MOUSE_INPUT` flag.

### **winput**

This corresponds to the `ENABLE_WINDOW_INPUT` flag.

### **iinsert**

This corresponds to the `ENABLE_INSERT_MODE` flag.

### **iqedit**

This corresponds to the `ENABLE_QUICK_EDIT_MODE` flag.

### **iautopos**

This corresponds to the `ENABLE_AUTO_POSITION` flag.

### **iext**

This corresponds to the `ENABLE_EXTENDED_FLAGS` flag.

### **oproc**

This corresponds to the `ENABLE_PROCESSED_OUTPUT` flag.

### **owrap**

This corresponds to the `ENABLE_WRAP_AT_EOL_OUTPUT` flag.

### **ivterm**

This corresponds to the `ENABLE_VIRTUAL_TERMINAL_INPUT` flag.

### **ovterm**

This corresponds to the `ENABLE_VIRTUAL_TERMINAL_OUTPUT` flag.

### **onewline**

This corresponds to the `DISABLE_NEWLINE_AUTO_RETURN` flag.

### **olvb**

This corresponds to the `ENABLE_LVB_GRID_WORLDWIDE` flag.