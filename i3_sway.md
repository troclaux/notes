
# i3/sway window manager

to reset sway's configuration file:

```bash
rm -rdf ~/.config/sway/config && cp /etc/sway/config ~/.config/sway/config
```

create directories to store configuration files and symlink config files from your dotfiles repo:

```bash
mkdir -p ~/.config/sway ~/.config/waybar ~/.config/wofi
ln -s ~/dotfiles/.config/sway/config ~/.config/sway/config
ln -s ~/dotfiles/.config/waybar/config/*  ~/.config/waybar/config
ln -s ~/dotfiles/.config/wofi/config  ~/.config/wofi/config
```

## prerequisites

fedora linux:

```bash
sudo dnf install sway light swaylock waybar wofi fontawesome-fonts wdisplays
```

debian-based linux:

```bash
sudo dnf install sway light swaylock waybar wofi fonts-font-awesome wdisplays
```

- `light`: modify display's brightness
- `swaylock`: lock screen
- `waybar`: highly customizable navigation bar
- `wofi`: wayland-native application launcher
- `fonts-font-awesome`: add icons for when you change volume
- `wdisplays`: manages display configuration

## keyboard shortcuts

- `mod + <direction>`: focuses window in the direction
- `mod + shift + <direction>`: moves currently focused window to desired direction

- `mod + enter`: opens terminal window
- `mod + f`: toggle fullscreen
- `mod + v`: split window vertically
- `mod + b`: split window horizontally
- `mod + r`: resize mode
- `mod + s`: stack windows
- `mod + d`: open application launcher

- `mod + shift + c`: reload sway's configuration
- `mod + shift + q`: kill focused window

- `mod + <0-9>`: switch to workspace
- `mod + shift + <0-9>`: move window to workspace


to find the name of input devices:

```bash
sudo libinput-listdevices
```

to find the name of monitors:

```bash
swaymsg -t get_outputs
```

to find the class of an application (search for `class` or `app_id`):

```bash
swaymsg -t get_tree | fzf | wl-copy
```
