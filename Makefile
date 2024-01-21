backup-packages:
	sudo pacman -Qqen > pkglist-repo.txt
	sudo pacman -Qqem > pkglist-aur.txt

restore-packages:
	sudo pacman -S --needed - < pkglist-repo.txt

restore-aur-packages:
	for x in $(< pkglist-aur.txt); do yay -S $x; done

add-chaotic-aur:
	sudo pacman-key --recv-key 3056513887B78AEB --keyserver keyserver.ubuntu.com &&
	sudo pacman-key --lsign-key 3056513887B78AEB &&
	sudo pacman -U 'https://cdn-mirror.chaotic.cx/chaotic-aur/chaotic-keyring.pkg.tar.zst' \
				   'https://cdn-mirror.chaotic.cx/chaotic-aur/chaotic-mirrorlist.pkg.tar.zst' &&
	sudo printf "[chaotic-aur]\nInclude = /etc/pacman.d/chaotic-mirrorlist" >> /etc/pacman.conf

backup-all:
	make backup-configs
	make backup-packages

reproduce:
	make remove-configs
	make restore-configs
	make add-chaotic-aur
	make restore-packages
	make restore-aur-packages
