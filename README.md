<div align="center">

<!-- <img src = ".assets/logo.png" alt="Repo Logo" height="320"/> -->

### Proto
Install and manage Proton versions easily 

#### Built With & Powered By
[![GoLang-Badge](https://img.shields.io/badge/GoLang-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

</div>

---

## About
Todo: create an about a good about explaining the whys and hows of the project

## Installing
### Using Homebrew
Install [homebrew](https://brew.sh/) on your system and then run the following:
```
brew tap bitsofabyte/proto https://github.com/BitsOfAByte/proto.git
brew install proto
```

### GitHub Releases
Download the [newest release](https://github.com/BitsOfAByte/proto/releases/latest) and extract it somewhere inside of your system path, or add the executable to your PATH instead. 

Refer to your processors documentation for the architecture you need to download.

### From Source (All Platforms)
If you would like to build Proto from source, install both GoLang & GoReleaser on to your system and then run the one-liner below:
```
git clone https://github.com/BitsOfAByte/proto && cd proto && goreleaser build --single-target --rm-dist --snapshot
```
You will find the compiled binary for your OS & Arch inside of the `/dist` folder,

