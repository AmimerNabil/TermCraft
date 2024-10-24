#!/bin/bash

# Function to show a progress bar
show_progress() {
	local duration=$1
	local bar_length=50
	local filled_length=0
	local increment=$((duration / bar_length))

	for ((i = 0; i <= bar_length; i++)); do
		filled_length=$((i * bar_length / bar_length))
		printf "\r[%-${bar_length}s] %d%%" $(head -c $filled_length </dev/zero | tr '\0' '#') $((i * 100 / bar_length))
		sleep "$increment"
	done
	echo ""
}

# Function to install SDKMAN if not already installed
install_sdkman() {
	if [ -z "$(command -v sdk)" ]; then
		echo -n "Installing SDKMAN... "
		{
			curl -s "https://get.sdkman.io" | bash
		} &
		show_progress 5 # Show progress bar for 5 seconds
		source "$HOME/.sdkman/bin/sdkman-init.sh"
		echo "SDKMAN installed."
		echo "Please add the following to your profile file (e.g., ~/.bashrc or ~/.zshrc):"
		echo 'export SDKMAN_DIR="$HOME/.sdkman"'
		echo '[[ -s "$SDKMAN_DIR/bin/sdkman-init.sh" ]] && source "$SDKMAN_DIR/bin/sdkman-init.sh"'
	else
		echo "SDKMAN already installed."
	fi
}

# Function to install pyenv using brew if not already installed
install_pyenv() {
	if [ -z "$(command -v pyenv)" ]; then
		echo -n "Installing pyenv... "
		{
			brew install pyenv
		} &
		show_progress 5 # Show progress bar for 5 seconds
		echo "pyenv installed."
		echo "Please add the following to your profile file (e.g., ~/.bashrc or ~/.zshrc):"
		echo 'export PATH="$HOME/.pyenv/bin:$PATH"'
		echo 'eval "$(pyenv init --path)"'
	else
		echo "pyenv already installed."
	fi
}

# Function to install fnm using brew if not already installed
install_fnm() {
	if [ -z "$(command -v fnm)" ]; then
		echo -n "Installing fnm... "
		{
			brew install fnm
		} &
		show_progress 5 # Show progress bar for 5 seconds
		echo "fnm installed."
		echo "Please add the following to your profile file (e.g., ~/.bashrc or ~/.zshrc):"
		echo 'eval "$(fnm env)"'
	else
		echo "fnm already installed."
	fi
}

# Function to verify that dependencies are working
verify_dependencies() {
	echo "Verifying dependencies..."

	# Check SDKMAN
	if [ -n "$(command -v sdk)" ]; then
		echo "SDKMAN is installed."
	else
		echo "SDKMAN is not installed. Please install it."
	fi

	# Check pyenv
	if [ -n "$(command -v pyenv)" ]; then
		echo "pyenv is installed."
	else
		echo "pyenv is not installed. Please install it."
	fi

	# Check fnm
	if [ -n "$(command -v fnm)" ]; then
		echo "fnm is installed."
	else
		echo "fnm is not installed. Please install it."
	fi

	echo "Dependency verification complete."
}

# Function to create the .termcraft directory structure
setup_termcraft() {
	echo "Setting up .termcraft directory structure..."
	mkdir -p "$HOME/.termcraft/src" "$HOME/.termcraft/configs"
	echo ".termcraft directory structure created."
}

# Function to fetch the latest release tarball from GitHub
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

# Function to build the Go project from the extracted release
build_go_project() {
	echo "Building Go project..."
	cd "$HOME/.termcraft/src" || exit 1 # Change to the extracted repo directory
	go build -o "$HOME/.termcraft/termcraft" . &
	show_progress 5 # Show progress bar for 5 seconds
	echo "Go project built."
}

# Function to remove old executable
remove_old_executable() {
	echo "Removing old termcraft files..."
	sudo rm -rf ~/.termcraft/src/* &
	show_progress 5 # Show progress bar for 5 seconds

	echo "Removing old termcraft executable..."
	sudo rm "$HOME/.termcraft/termcraft" &
	show_progress 5 # Show progress bar for 5 seconds
	echo "Old executable removed."
}

# Main process
case "$1" in
-U | --update)
	echo "Starting update process..."
	verify_dependencies # Verify dependencies before updating
	remove_old_executable
	fetch_latest_release
	build_go_project
	echo "Update complete! Please ensure that the following path is added to your PATH in your profile file (e.g., ~/.bashrc or ~/.zshrc):"
	echo 'export PATH="$HOME/.termcraft:$PATH"'
	;;
*)
	echo "Starting installation process..."
	install_sdkman
	install_pyenv
	install_fnm
	verify_dependencies # Verify dependencies during installation
	setup_termcraft
	fetch_latest_release
	build_go_project
	echo "Installation complete! Please add the following path to your PATH in your profile file (e.g., ~/.bashrc or ~/.zshrc):"
	echo 'export PATH="$HOME/.termcraft:$PATH"'
	;;
esac
