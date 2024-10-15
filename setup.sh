#!/bin/bash

# Function to show spinner
spinner() {
	local pid=$!
	local delay=0.1
	local spinstr='|/-\'
	while [ "$(ps a | awk '{print $1}' | grep $pid)" ]; do
		local temp=${spinstr#?}
		printf " [%c]  " "$spinstr"
		spinstr=$temp${spinstr%"$temp"}
		sleep $delay
		printf "\b\b\b\b\b\b"
	done
	printf "    \b\b\b\b"
}

# Install SDKMAN if not already installed
install_sdkman() {
	if [ -z "$(command -v sdk)" ]; then
		echo "Installing SDKMAN..."
		curl -s "https://get.sdkman.io" | bash
		source "$HOME/.sdkman/bin/sdkman-init.sh"
		echo "SDKMAN installed."
	else
		echo "SDKMAN already installed."
	fi
}

# Install pyenv using brew if not already installed
install_pyenv() {
	if [ -z "$(command -v pyenv)" ]; then
		echo "Installing pyenv with Homebrew..."
		brew install pyenv &
		spinner
		echo "pyenv installed."
	else
		echo "pyenv already installed."
	fi
}

# Build the Go project
install_go_project() {
	echo "Building Go project..."
	go build -o termcraft . &
	spinner
	echo "Go project built."
}

# Move executable to /usr/local/bin
move_executable() {
	echo "Moving executable to /usr/local/bin..."
	sudo mv termcraft /usr/local/bin/ &
	spinner
	echo "Executable moved to /usr/local/bin."
}

# Main installation process
echo "Starting installation process..."

install_sdkman
install_pyenv
install_go_project
move_executable

echo "Installation complete!"
