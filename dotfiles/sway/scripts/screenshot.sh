#!/bin/bash
 
entries="Area Window"
 
selected=$(printf '%s\n' $entries | wofi --style=$HOME/.config/wofi/style.widgets.css --conf=$HOME/.config/wofi/config.screenshot | awk '{print tolower($1)}')
 
case $selected in
  area)
    /usr/share/sway/scripts/grimshot copy area;;
  window)
    /usr/share/sway/scripts/grimshot copy window;;
esac
