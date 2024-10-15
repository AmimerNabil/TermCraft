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

# Create the .termcraft directory structure
setup_termcraft() {
	echo "Setting up .termcraft directory structure..."
	mkdir -p "$HOME/.termcraft/src" "$HOME/.termcraft/configs"
	echo ".termcraft directory structure created."
}

# Fetch the latest release tarball
fetch_latest_release() {
	local repo_url="AmimerNabil/TermCraft" # Change to your repository URL
	echo "Fetching the latest release from GitHub..."
	local release_info
	release_info=$(curl -s "https://api.github.com/repos/$repo_url/releases/latest")
	local tarball_url=$(echo "$release_info" | grep "tarball_url" | cut -d '"' -f 4)

	if [ -z "$tarball_url" ]; then
		echo "Failed to fetch the latest release."
		exit 1
	fi

	echo "Downloading the latest release..."
	curl -L -o latest_release.tar.gz "$tarball_url"
	echo "Latest release downloaded."

	echo "Extracting the release..."
	tar -xzf latest_release.tar.gz -C "$HOME/.termcraft/src" --strip-components=1
	echo "Latest release extracted."
}

# Build the Go project from the extracted release
install_go_project() {
	echo "Building Go project..."
	cd "$HOME/.termcraft/src" || exit 1 # Change to the extracted repo directory
	go build -o "$HOME/.termcraft/termcraft" . &
	spinner
	echo "Go project built."
}

# Move executable to /usr/local/bin
move_executable() {
	echo "Moving executable to /usr/local/bin..."
	sudo mv "$HOME/.termcraft/termcraft" /usr/local/bin/ &
	spinner
	echo "Executable moved to /usr/local/bin."
}

# Main installation process
echo "Starting installation process..."

install_sdkman
install_pyenv
setup_termcraft
fetch_latest_release
install_go_project
move_executable

echo "Installation complete!"
