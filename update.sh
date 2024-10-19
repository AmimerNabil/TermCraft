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

# Remove the old executable
remove_old_executable() {
	echo "Removing old termcraft files..."
	sudo rm -rf ~/.termcraft/src/* &
	spinner

	echo "Removing old termcraft executable..."
	sudo rm /usr/local/bin/termcraft &
	spinner
	echo "Old executable removed."
}

# Fetch the latest release tarball from GitHub
fetch_latest_release() {
	local repo_url="AmimerNabil/TermCraft" # Change this to your repo URL
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
}

# Extract the tarball to the .termcraft directory
extract_release() {
	echo "Extracting the latest release..."
	tar -xzf latest_release.tar.gz -C "$HOME/.termcraft/src" --strip-components=1
	rm latest_release.tar.gz
	echo "Latest release extracted."
}

# Build the Go project from the extracted release
build_go_project() {
	echo "Building Go project..."
	cd "$HOME/.termcraft/src" || exit 1
	go build -o "$HOME/.termcraft/termcraft" . &
	spinner
	echo "Go project built."
}

# Move the new executable to /usr/local/bin
update_executable() {
	echo "Moving new executable to /usr/local/bin..."
	sudo mv "$HOME/.termcraft/termcraft" /usr/local/bin/ &
	spinner
	echo "Executable updated."
}

# Main update process
echo "Starting update process..."

remove_old_executable
fetch_latest_release
extract_release
build_go_project
update_executable

echo "Update complete!"
