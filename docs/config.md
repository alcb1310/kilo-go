# Configuring `Kilo` text editor

The `Kilo` text editor will look for its configuration file inside the `XDG_CONFIG` directory, which is usually `$HOME/.config`.

- On Unix systems, it returns `$XDG_CONFIG_HOME` as specified by https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if non-empty, else `$HOME/.config`.
- On Darwin (macOS), it returns `$HOME/Library/Application Support`.
- On Windows, it returns `%AppData%`.

Inside that directory it will look for a file named `kilo/config.toml`.

The `config.toml` file is a [TOML](https://toml.io/) file that contains the configuration for `Kilo` text editor, and will have the following sections:

- Settings
- Theme

## The `Settings` section

It will be a table with the following keys:

```toml
[settings]
tab_size = 4 # Default: 8
quit_times = 3 # Default: 3
```

### `tab_size`

This will be a number that will represent how many spaces will be used for a tab.

### `quit_times`

This will be a number that will represent how many times the user has to press `Ctrl+Q` to quit `Kilo` when a file has been modified.


## The `Theme` section

It will be a table with the following keys:

```toml
[theme]
comment=[0,25,255] # Default [0,255,255]
default=[240,240, 240] # Default [255,255,255]
keyword=[255,239,0] # Default [255,239,0]
number=[255,0,0] # Default [255,0,0]
search=[51,255,0] # Default [51,255,0]
string=[255,39,155] # Default [255,39,155]
type=[12,239,55] # Default [12,239,55]
```

Each color is a list of 3 numbers that represent the red, green and blue values of the color, with values that ranges from 0 to 255.
