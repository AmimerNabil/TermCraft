## TermCraft 🧰

![image](https://github.com/user-attachments/assets/621fafda-33bd-48d4-bb75-36f84530409f)

Termcraft is a powerful text-based user interface designed to simplify the management of multiple programming languages and frameworks. With Termcraft, you can easily install, manage, and view different versions of languages such as Java ☕, Python 🐍, Node 🌐, Kotlin 🎯, Go 🚀, and many more. This tool streamlines version control across your development environment, making it easy to switch between languages and frameworks as needed. In the future, Termcraft will also offer functionality to help users manage and share configuration files 🛠️, providing a comprehensive solution for developers.


#### Available Languages for now:
1. java
2. python

#### Languages next:
1. go
2. (node versions)

## Table of Contents

1. [TermCraft 🧰](#termcraft-)  
2. [Installation](#installation)  
   2.1. [Prerequisites](#prerequisites)  
      2.1.1. [1. Install Homebrew](#1-install-homebrew)  
      2.1.2. [2. Install Go (v1.19 or higher)](#2-install-go-v119-or-higher)  
   2.2. [Proceed with the Installation](#proceed-with-the-installation)  
   2.3. [Additional Notes](#additional-notes)  
3. [How to Use 🚀](#how-to-use-)  

## Installation

The goal is for you to have to install a minimal amount of things to use termcraft. It will handle the rest for you so go through these and never look back.

## Prerequisites

Before running the setup script, ensure you have the following installed on your system:

### 1. Install Homebrew

Homebrew is a package manager for macOS and Linux. If you haven't installed it yet, open your terminal and run:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

### 2. Install Go (v1.19 or higher)

Once you have Homebrew installed, you can install Go using the following command:

```bash
brew install go
```

To verify that Go is installed correctly, you can check the version:

```bash
go version
```

Ensure the output indicates that you have Go version 1.19 or higher.

### Proceed with the Installation

After completing the prerequisites, you can run the setup script to install and configure the project:

```bash
bash <(curl -s https://raw.githubusercontent.com/AmimerNabil/TermCraft/main/setup.sh)
```

Add the following to your curr profile like `~/.zshrc` or `~/.bashrc` but make sure to always keep **SDKman at the very bottom of your profile**

```bash
export PYENV_ROOT="$HOME/.pyenv"
[[ -d $PYENV_ROOT/bin ]] && export PATH="$PYENV_ROOT/bin:$PATH"
eval "$(pyenv init -)"
```

### Additional Notes

- **For Linux Users**: If you are using Linux and prefer to install Go using the official method instead of Homebrew, you can follow the [official Go installation guide](https://golang.org/doc/install/source).

- **Troubleshooting**: If you encounter any issues during the installation, please refer to the official documentation for [Homebrew](https://docs.brew.sh/) and [Go](https://golang.org/doc/).

Here’s the updated "How to Use" section with code blocks for running Termcraft:

## How to Use 🚀

Using Termcraft is simple! To get started, just run the Termcraft application:

```bash
termcraft
```
Once you're in, you'll see the supported languages displayed on the left side of the interface.

Keep in mind that Termcraft only manages language versions for the languages it installs. It does not oversee versions of other languages outside of its purview.

Whenever you need assistance, you can type `?` at any point in the application to view the current available commands:

Now, just navigate to your preferred language, select your desired version, and you’re all set! With Termcraft, you’ll never have to look anywhere else for managing your programming language versions. Happy coding!
