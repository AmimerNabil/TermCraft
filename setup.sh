#!/bin/bash

# Simple function to show a progress bar
show_progress() {
	local duration=$1
	local bar_length=50
	local filled=0

	for ((i = 0; i <= bar_length; i++)); do
		filled=$((i * bar_length / bar_length))
		printf "\r[%-${bar_length}s] %d%%" $(head -c $filled </dev/zero | tr '\0' '#') $((i * 100 / bar_length))
		sleep $((duration / bar_length))
	done
	echo ""
}

# Function to install SDKMAN if not already installed
install_sdkman() {
	if [ -z "$(command -v sdk)" ]; then
		echo "Installing SDKMAN..."
		curl -s "https://get.sdkman.io" | bash >/dev/null 2>&1
		show_progress 5
		source "$HOME/.sdkman/bin/sdkman-init.sh"
		echo "SDKMAN installed."
		echo "Add this to your profile (e.g., ~/.bashrc or ~/.zshrc):"
		echo 'export SDKMAN_DIR="$HOME/.sdkman"'
		echo '[[ -s "$SDKMAN_DIR/bin/sdkman-init.sh" ]] && source "$SDKMAN_DIR/bin/sdkman-init.sh"'
	else
		echo "SDKMAN already installed."
	fi
}

# Function to install pyenv
install_pyenv() {
	if [ -z "$(command -v pyenv)" ]; then
		echo "Installing pyenv..."
		brew install pyenv >/dev/null 2>&1
		show_progress 5
		echo "pyenv installed."
		echo "Add this to your profile (e.g., ~/.bashrc or ~/.zshrc):"
		echo 'export PATH="$HOME/.pyenv/bin:$PATH"'
		echo 'eval "$(pyenv init --path)"'
	else
		echo "pyenv already installed."
	fi
}

# Function to install fnm
install_fnm() {
	if [ -z "$(command -v fnm)" ]; then
		echo "Installing fnm..."
		brew install fnm >/dev/null 2>&1
		show_progress 5
		echo "fnm installed."
		echo "Add this to your profile (e.g., ~/.bashrc or ~/.zshrc):"
		echo 'eval "$(fnm env)"'
	else
		echo "fnm already installed."
	fi
}

# Function to fetch the latest release tarball from GitHub
fetch_latest_release() {
	local repo_url="AmimerNabil/TermCraft" # Replace with your repository URL
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

# Function to build the Go project from the extracted release
build_go_project() {
	echo "Building Go project..."
	cd "$HOME/.termcraft/src" || exit 1
	go build -o "$HOME/.termcraft/termcraft" . &
	show_progress 5
	echo "Go project built."
}

# Function to remove old termcraft executable
remove_old_executable() {
	echo "Removing old TermCraft files..."
	rm -rf "$HOME/.termcraft/src/*"
	show_progress 5
	echo "Removing old TermCraft executable..."
	rm -f "$HOME/.termcraft/termcraft"
	show_progress 5
	echo "Old executable removed."
}

# Function to verify dependencies
verify_dependencies() {
	echo "Verifying dependencies..."

	if [ -f "$HOME/.sdkman/bin/sdkman-init.sh" ]; then
		# Source the SDKMAN initialization script to make the 'sdk' command available
		source "$HOME/.sdkman/bin/sdkman-init.sh"
		if command -v sdk >/dev/null 2>&1; then
			echo "SDKMAN is installed."
		else
			echo "SDKMAN is installed, but 'sdk' command is not available."
		fi
	else
		echo "SDKMAN is not installed."
	fi

	if [ -n "$(command -v pyenv)" ]; then
		echo "pyenv is installed."
	else
		echo "pyenv is not installed."
	fi

	if [ -n "$(command -v fnm)" ]; then
		echo "fnm is installed."
	else
		echo "fnm is not installed."
	fi
	echo "Dependency verification complete."
}

# Main process for updating or installing
case "$1" in
-U | --update)
	echo "Starting update process..."
	verify_dependencies
	remove_old_executable
	fetch_latest_release
	build_go_project
	echo "Update complete! Ensure the path is added to your PATH (e.g., ~/.bashrc or ~/.zshrc):"
	echo 'export PATH="$HOME/.termcraft:$PATH"'
	;;
*)
	echo "Starting installation process..."
	install_sdkman
	install_pyenv
	install_fnm
	verify_dependencies
	fetch_latest_release
	build_go_project
	echo "Installation complete! Add the following to your PATH (e.g., ~/.bashrc or ~/.zshrc):"
	echo 'export PATH="$HOME/.termcraft:$PATH"'
	;;
esac
