<!-- Repository Header Begin -->
<div align="center">

<img src="./.assets/Banners/banner.png" alt="Proto Logo">
  
Install and manage custom Proton installations with ease on supported systems.

**[View Issues](https://github.com/ProtoSoftware/proto/issues) · [Install](#installation) · [Contributing](https://github.com/ProtoSoftware/proto/blob/main/CONTRIBUTING.md)**
  
<a href="#"> 
  <img src="https://img.shields.io/github/downloads/ProtoSoftware/proto/total?style=flat" alt="Download Count Badge">
  <img src="https://img.shields.io/github/v/tag/ProtoSoftware/proto?color=blue&label=Version&sort=semver&style=flat" alt="Release Badge">
</a>
  
</div>

---

<!-- Repository Header End -->

## About
Proto is a tool designed to make downloading and managing custom WINE/Proton installations as convinent and easy as possible. It provides support for multiple sources, custom installation directories, in-tool release data and more. 

We aim to keep the tool both beginner friendly so that any user regardless of skill level can use it with a little bit of guidance while still maintaining the customisability that a power user might want from the tool.

## How to Use
First off, you'll need to download Proto to your system using one of the supported [installation methods](#installation). Once you have it installed, the entire app is documented in the command line by running the `--help` flag after any command, which will provide details on how to use it. 

Configuration is also provided by running the `proto config` command, which will allow you to tweak a variety of settings straight from the command line.

## FaQ 
### Is There a GUI?
Not quite yet, a GUI is definitely planned for a future release once the command line has been polished and is feature complete and the shared utility & optimised code is all in place to make a GUI implementation as smooth as possible. Once finished, the GUI should have feature parity with the CLI. The GUI is not planned for release until `v2.0.0` at the earliest.

### What makes this different from other tools?
A lot of other tools that download custom Proton installations might not offer the same extendability that Proto is aiming to offer, and often times get left to the side lacking bug fixes and optimizations. Once Proto is complete, it will be maintained to continue offering the same features until a time where it is no longer needed, either due to in-app support from platforms like Steam and Lutris or otherwise.

### What are "Curated Sources"?
A curated source is just a repository that repackages releases from other maintainers in a way that will be ensured to work with Proto, preventing the need to implement multiple methods of handling sources. Releases are automatically fetched and relayed to the curated repositories daily and a maintainer will repackage it manually without modifiying any of the contents.

However, it is very important to know that these do not have to be used! If you do not feel comfortable with the repackaging of releases and would rather get it straight from the Maintainer, you can do this by adding their repository as a source with `proto config sources add <owner/repo>` as long as they release in a format that supports the [standard](./STANDARDS.md) for Proto releases.

### Is Proto Stable?
Not currently, breaking changes are likely to occur at this time in development as the tool is still being built and things need to be adjusted a lot. However, this does not mean it is unusable. When this tool is finished with early rapid-development, a version released as "v1.0.0" will be published.

## Installation

### Dependancies
Proto currently requires the following packages in order to function: [tar](https://www.gnu.org/software/tar/)

If you are using a package manager to install, these should be automatically installed alongside Proto if they are missing from your system, however if you are building from source or installing from an archive, make sure these are also present. It is planned to be dependency free by time a full release is made. 

### Methods

#### APT Package Manager
If you are using an Ubuntu-derivative system then use this installation method.

<details>
<summary>Show Steps</summary>

<br>
  
1. Add the repository hosting Proto to your apt sources directory (Only run this once)
```
echo "deb [trusted=yes] https://packages.bitsofabyte.dev/apt/ /" | sudo tee -a /etc/apt/sources.list.d/bitsofabyte.list && sudo apt update
``` 

2. Install Proto to your system
```
sudo apt install proto
```

</details>  

---

#### Yum/DNF Package Manager
If you are using Fedora, OpenSUSE, or any other system that supports the yum/dnf package manager then use this installation method.

<details>
<summary>Show Steps</summary>
<br>
  
1. Add the repository hosting Proto to your yum/dnf repo directory (Only run this once)
```
echo "[BitsOfAByte]            
name=BitsOfAByte Packages         
baseurl=https://packages.bitsofabyte.dev/yum/
enabled=1
gpgcheck=0" | sudo tee -a /etc/yum.repos.d/bitsofabyte.repo && sudo yum update
``` 

2. Install Proto to your system
```
sudo yum install proto
```

</details>  

---

#### Homebrew Package Manager
If your distributions package manager is not listed here or you wish to use [Homebrew](https://brew.sh).

<details>
<summary>Show Steps</summary>
<br>
  
1. Install homebrew if you haven't already got it
```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

2. Add the tap for Proto to homebrew
```
brew tap ProtoSoftware/proto https://github.com/ProtoSoftware/proto.git
```

3. Install proto to your system
```
brew install proto
```
  
</details>

---

#### Manual Installation
Manually install a Binary from the release Archives.
<details>  
<summary>Show Steps</summary>
  
1. Download the [newest release](https://github.com/ProtoSoftware/proto/releases/latest) for your system/architecture
2. Extract the binary into your system path or add the binary to your path.

If you aren't sure on what architecture you need to download, you should try `amd64` first as it is the most common for everyday hardware.

</details>

---

#### From Source
Build Proto directly from the GitHub source for any supported platform.
<details>  
<summary>Show Steps</summary>
  
Building Proto from source is not recommended for beginners, but if you know what you're doing then follow these steps: 
1. Install [Go](https://go.dev/) on your system
2. Download the [GoReleaser](https://goreleaser.com/) package
3. Clone the repository to your system with `git clone https://github.com/ProtoSoftware/proto`
4. Inside the repository directory, run `goreleaser build --single-target --rm-dist --snapshot` to build.

You will find the compiled binary for your OS & Arch inside of the `/dist` folder.

</details>  
