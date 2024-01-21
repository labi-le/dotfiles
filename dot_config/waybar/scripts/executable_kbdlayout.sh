#!/bin/sh
hyprctl devices -j | jq '.["keyboards"][-2]["active_keymap"]' | sed 's/"//g'