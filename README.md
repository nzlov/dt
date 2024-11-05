# DT

Get the theme color based on the picture and generate the program color according to the template.

## config.json

path: `$HOME/.config/dt/config.yaml`

```
img: /home/nzlov/.config/hypr/wallpaper.jpg 
tmpls:
  - t: /home/nzlov/.nzlovdotfile/config/waybar/waybar.tmpl
    o: /home/nzlov/.config/waybar/style.css
    e: killall -SIGUSR2 waybar
```

### img

local file or online url

### tmpls

template list

* t template path
* o output path
* e after execute command
