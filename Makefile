CONFIG_DIR = $(HOME)/.config/
DOTFILES_DIR = ./dotfiles/
CONFIGS = {sway,mako,waybar,wofi,micro,ranger,fish}
CONFIGS_IN_HOME = .profile

build:
	rm -rf $(DOTFILES_DIR) && mkdir -p $(DOTFILES_DIR)
	cp -r $(CONFIG_DIR)$(CONFIGS) $(DOTFILES_DIR)
	cp -r $(HOME)/$(CONFIGS_IN_HOME) $(DOTFILES_DIR)

install:
	mkdir -p $(CONFIG_DIR) && cp -r $(DOTFILES_DIR)$(CONFIGS) $(CONFIG_DIR)

remove:
	rm -rf $(CONFIG_DIR)$(CONFIGS)