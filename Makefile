CONFIG_DIR = $(HOME)/.config/
DOTFILES_DIR = ./dotfiles/

build:
	rm -rf $(DOTFILES_DIR) && mkdir -p $(DOTFILES_DIR)
	cp -r $(CONFIG_DIR){sway,mako,waybar,wofi,micro} $(DOTFILES_DIR)
	cp -r $(HOME)/.bashrc $(DOTFILES_DIR)

install:
	mkdir -p $(CONFIG_DIR) && cp -r $(DOTFILES_DIR){sway,mako,waybar,wofi,micro} $(CONFIG_DIR)

remove:
	rm -rf $(CONFIG_DIR){sway,mako,waybar,wofi,micro}