
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

fedora distribution:

```bash
sudo dnf install -y sway kanshi light swaylock waybar wofi fuzzel wdisplays wlr-randr pavucontrol grimshot copyq playerctl dunst
sudo dnf install -y fontawesome-fonts network-manager-applet
```

debian-based distribution:

```bash
sudo apt install -y sway kanshi light swaylock waybar wofi fuzzel wdisplays wlr-randr pavucontrol grimshot copyq playerctl dunst
sudo apt install -y fonts-font-awesome network-manager-gnome
```

- `sway`: tiling window manager
- `kanshi`: display configuration tool
- `light`: modify display's brightness
- `swaylock`: lock screen
- `waybar`: highly customizable navigation bar
- `wofi`: wayland-native application launcher (compatible with shell piping, used for scripts)
- `fuzzel`: wayland-native application launcher (used to open gui applications)
- `wdisplays`: manages display configuration
- `wlr-randr`: display management
- `pavucontrol`: input/output audio device management
- `playerctl`: playback control tool
- `grimshot`: screenshot utility tool
- `copyq`: clipboard manager
- `fonts-font-awesome/fontawesome-fonts`: add icons for when you change volume
- `network-manager-applet/network-manager-gnome`: wi-fi configuration
- `dunst`: notification manager

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
swaymsg -t get_outputs | grep 'name' | awk '{print $2}' | tr -d '",'
```

to find the class of an application (search for `class` or `app_id`):

```bash
swaymsg -t get_tree | fzf | wl-copy
```

```bash
mkdir -p ~/.config/kanshi
```

## change default application for a type of file

[arch wiki](https://wiki.archlinux.org/title/XDG_MIME_Applications)

find desktop entry id of default application:

```bash
ls /usr/share/applications/
```

edit mimeapps.list file:

```bash
nvim ~/.config/mimeapps.list
```

add entry with `type/extension=desktop_entry_id`, just like the example below (loupe is gnome's image viewer):

```
image/jpeg=org.gnome.Loupe.desktop
image/png=org.gnome.Loupe.desktop
image/gif=org.gnome.Loupe.desktop
```
