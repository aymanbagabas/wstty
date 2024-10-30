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

Use `-a` to displays all the current console modes, `-h` to display the help.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file