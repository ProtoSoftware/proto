<div align="center">

<img src="./.assets/Banner.png" alt="Proto Logo">
  
Install and manage custom Proton installations with ease on supported systems.

**[View Issues](https://github.com/BitsOfAByte/proto/issues) · [Download](https://github.com/BitsOfAByte/proto#quick-start=) · [Contributing](https://github.com/BitsOfAByte/proto/blob/main/CONTRIBUTING.md)**
  
<a href="#"> 
  <img src="https://img.shields.io/github/downloads/BitsOfAByte/proto/total?style=flat" alt="Download Count Badge">
  <img src="https://img.shields.io/github/v/tag/BitsOfAByte/proto?color=blue&label=Version&sort=semver&style=flat" alt="Release Badge">
</a>
  
</div>

## About
Proto was designed as a beginner friendly and convinent tool for downloading and managing external Proton installations that follow the [Proto Standards](./STANDARDS.md) for their releases. Power users can enjoy the extra configuration options & even build ontop of Proto for automations.

### Key Features
Proto currently boasts the following features: 

- Download Checksums (See STANDARDS)
- Multiple proton installations
- Shell completion generation

## Quick Start
####  Express Installation 🚀 (Recommended)
Coming Soon 🚀

### Other Methods
#### Homebrew
<details>  
<summary>Show Steps</summary>
  
1. Install [homebrew](https://brew.sh).
2. Run `brew tap bitsofabyte/proto https://github.com/BitsOfAByte/proto.git` to add the repository to homebrew.
3. Run `brew install proto` to install Proto.
</details>

#### Manual Installation
<details>  
<summary>Show Steps</summary>
  
1. Download the [newest release](https://github.com/BitsOfAByte/proto/releases/latest) for your system/architecture
2. Extract the binary into your system path or add the binary to your path.

If you aren't sure on what architecture you need to download, you should try `amd64` first as it is the most common for everyday hardware.
</details>

#### Building From Source
<details>  
<summary>Show Steps</summary>
  
Building Proto from source is not recommended for beginners, but if you know what you're doing then follow these steps: 
1. Install [Go](https://go.dev/) on your system
2. Download the [GoReleaser](https://goreleaser.com/) package
3. Clone the repository to your system with `git clone https://github.com/BitsOfAByte/proto`
4. Inside the repository directory, run `goreleaser build --single-target --rm-dist --snapshot` to build.

You will find the compiled binary for your OS & Arch inside of the `/dist` folder,
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
