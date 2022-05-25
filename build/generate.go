package main

import (
	"fmt"
	"io"
	"os"
)

// Tasks to run as a pre-build hook
func main() {
	build_dir := "./.build_data/"

	createBuildDir(build_dir)

	generateDesktop()
	generateMetainfo()
	generateIcon()

}

// Create the build directory if it doesn't exist
func createBuildDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
}

// Create a file in the build directory
func createBuildFile(fileName string, data string) {
	file, err := os.Create("./.build_data/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	file.Sync()
}

// Generate the .desktop file
func generateDesktop() {

	version := os.Args[1]

	if version == "" {
		panic("No version specified")
	}

	fileData := fmt.Sprintf(`[Desktop Entry]
Version=%s
Type=Application
Name=Proto
GenericName=Proto
Comment=Proto compatability tool manager
Icon=/usr/share/icons/proto/icon.png
Exec=proto gui
Terminal=true
Actions=NewShortcut;
Categories=ConsoleOnly;Utility;X-GNOME-Utilities;FileTools;
Keywords=proton;steamplay;
StartupNotify=true

[Desktop Action NewShortcut]
Name=Install Latest Proton
Exec=proto install -y`, version)

	createBuildFile("proto.desktop", fileData)
}

// Generate the .metainfo.xml file
func generateMetainfo() {
	fileData := `<?xml version="1.0" encoding="UTF-8"?>
<!-- Copyright 2022 BitsOfAByte -->
<component type="desktop">
  <name>Proto</name>
  <id>proto.desktop</id>
  <developer_name>BitsOfAByte<developer_name/>
  <launchable type="desktop-id">proto.desktop</launchable>
  <metadata_license>CC0-1.0</metadata_license>
  <project_license>GPL-3.0-only</project_license>
  <provides>
  	<binary>proto</binary>
  </provides>
  <summary>Proto compatability tool manager</summary>
  <description>
    <p>
      Easily manage multiple custom installations of the Proton compatability tool.
    </p>
  </description>
  <keywords>
    <keyword>proto</keyword>
    <keyword>utility</keyword>
  </keywords>
  <url type="bugtracker">https://github.com/BitsOfAByte/proto/issues</url>
  <url type="contact">https://github.com/BitsOfAByte/proto</url>
</component>`

	createBuildFile("proto.metainfo.xml", fileData)
}

// Fetch the icon from the assets and put it in the build directory
func generateIcon() {
	srcFile, err := os.Open("./.assets/Logos/icon.png")
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	destFile, err := os.Create("./.build_data/icon.png")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		panic(err)
	}

	err = destFile.Sync()
	if err != nil {
		panic(err)
	}
}
