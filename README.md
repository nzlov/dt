# DT

Get the theme color based on the picture and generate the program color according to the template.

## config.json

path: `$HOME/.config/dt/config.yaml`

```
size: 3
bg: 2
img: /home/nzlov/.config/hypr/wallpaper.jpg 
tmpls:
  - t: /home/nzlov/.nzlovdotfile/config/waybar/waybar.tmpl
    o: /home/nzlov/.config/waybar/style.css
    e: killall -SIGUSR2 waybar
```

### size

pick color number

### bg

default background color index

### img

local file or online url

### tmpls

template list

* t template path
* o output path
* e after execute command

## templates

* CX out: hex color
* ORX out: or background color hex color
* ANDX out: and background color hex color
* RGBX out: r,g,b color

`X` is color index
