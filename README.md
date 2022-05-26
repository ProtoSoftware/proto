<div align="center">

<img src="./.assets/Banners/banner.png" alt="Proto Logo">
  
Install and manage custom Proton installations with ease on supported systems.

**[View Issues](https://github.com/BitsOfAByte/proto/issues) · [Download](#installation-methods) · [Contributing](https://github.com/BitsOfAByte/proto/blob/main/CONTRIBUTING.md)**
  
<a href="#"> 
  <img src="https://img.shields.io/github/downloads/BitsOfAByte/proto/total?style=flat" alt="Download Count Badge">
  <img src="https://img.shields.io/github/v/tag/BitsOfAByte/proto?color=blue&label=Version&sort=semver&style=flat" alt="Release Badge">
</a>
  
</div>

## About
Proto was designed as a beginner friendly and convinent tool for downloading and managing external Proton installations that follow the [Proto Standards](./STANDARDS.md) for their releases. Power users can enjoy the extra configuration options & even build ontop of Proto for automations.

## How to Use
Install Proto to your system using one of the methods listed below, then simply run `proto -h` for an up to date list of commands. 

To see all configuration options for Proto run `proto config -h`

## Installation Methods
### APT Package Manager
If you are using an Ubuntu-derivative system then use this installation method.

<details>
<summary>Show Steps</summary>

<br>
  
1. Add the repository hosting Proto to your apt sources directory
```
echo "deb [trusted=yes] https://packages.bitsofabyte.dev/apt/ /" | sudo tee -a /etc/apt/sources.list.d/bitsofabyte.list
``` 

2. Fetch all sources again to detect the new list
```
sudo apt update
 ```

3. Install Proto to your system
```
sudo apt install proto
```
  
</details>  

---

### Yum/DNF Package Manager
If you are using Fedora, OpenSUSE, or any other system that supports the yum/dnf package manager then use this installation method.

<details>
<summary>Show Steps</summary>
<br>
  
1. Add the repository hosting Proto to your yum repos directory
```
echo "[BitsOfAByte]            
name=BitsOfAByte Packages         
baseurl=https://packages.bitsofabyte.dev/yum/
enabled=1
gpgcheck=0" | sudo tee -a /etc/yum.repos.d/bitsofabyte.repo
``` 
  
2. Fetch all sources again to detect the new list
```
sudo yum update
```

3. Install Proto to your system
```
sudo yum install proto
```

</details>  

---

### Homebrew Package Manager
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
brew tap bitsofabyte/proto https://github.com/BitsOfAByte/proto.git
```
3. Install proto to your system
```
brew install proto
```
  
</details>

---

### Manual Installation
Manually install a Binary from the release Archives.
<details>  
<summary>Show Steps</summary>
  
1. Download the [newest release](https://github.com/BitsOfAByte/proto/releases/latest) for your system/architecture
2. Extract the binary into your system path or add the binary to your path.

If you aren't sure on what architecture you need to download, you should try `amd64` first as it is the most common for everyday hardware.

</details>

---

### From Source
Build Proto directly from the GitHub source for any supported platform.
<details>  
<summary>Show Steps</summary>
  
Building Proto from source is not recommended for beginners, but if you know what you're doing then follow these steps: 
1. Install [Go](https://go.dev/) on your system
2. Download the [GoReleaser](https://goreleaser.com/) package
3. Clone the repository to your system with `git clone https://github.com/BitsOfAByte/proto`
4. Inside the repository directory, run `goreleaser build --single-target --rm-dist --snapshot` to build.

You will find the compiled binary for your OS & Arch inside of the `/dist` folder.

</details>  

## FaQ 

### Will there be a GUI for Proto?
Yes! A GUI is planned to be added to the project in due time, but development on this has not yet started as other things take higher priority (improved utility, decoupled methods)

### What makes this different from other tools?
A lot of other tools that offer this utility are no longer maintained actively, leaving them to stop working over time. Proto is planned to be recieve bug fixes and optimizations long after active feature development is finished.

## Contributing
Contributions are welcomed and encouraged by all levels of skill be it an issue or pull request. If you would like to contribute, please check out the [guide](./CONTRIBUTING.md) for more information on how to get started.

## License
[GPL-3.0-only](https://choosealicense.com/licenses/gpl-3.0/), license can be found [here](./LICENSE).
