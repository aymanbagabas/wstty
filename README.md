# wstty

<p>
    <a href="https://github.com/aymanbagabas/wstty/actions"><img src="https://github.com/aymanbagabas/wstty/workflows/build/badge.svg" alt="Build Status"></a>
    <a href="https://github.com/aymanbagabas/wstty/releases"><img src="https://img.shields.io/github/release/aymanbagabas/wstty.svg" alt="Latest Release"></a>
</p>

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