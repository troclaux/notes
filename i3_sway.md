
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
sudo dnf install -y sway light swaylock waybar wofi fuzzel wdisplays network-manager-applet wlr-randr pavucontrol copyq
sudo dnf install -y fontawesome-fonts
```

debian-based distribution:

```bash
sudo apt install -y sway light swaylock waybar wofi fuzzel wdisplays network-manager-applet wlr-randr pavucontrol copyq
sudo apt install -y fonts-font-awesome
```

- `light`: modify display's brightness
- `swaylock`: lock screen
- `waybar`: highly customizable navigation bar
- `fuzzel`: wayland-native application launcher (I prefer this one)
- `wofi`: wayland-native application launcher (I don't like this one)
- `fonts-font-awesome`: add icons for when you change volume
- `wdisplays`: manages display configuration
- `network-manager-applet`: wi-fi configuration
- `wlr-randr`: display management
- `pavucontrol`: input/output audio device management

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
