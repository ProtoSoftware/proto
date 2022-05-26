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
Exec=proto install`, version)

	createBuildFile("dev.bitsofabyte.proto.desktop", fileData)
}

// Generate the .metainfo.xml file
func generateMetainfo() {
	fileData := `<?xml version="1.0" encoding="UTF-8"?>
<!-- Copyright 2020 BitsOfAByte -->
<component type="desktop-application">
  <id>dev.bitsofabyte.proto</id>
  <metadata_license>MIT</metadata_license>
  <project_license>GPL-3.0-only</project_license>
  <name>Proto</name>
  <summary>Manage custom Proton installations</summary>

  <description>
    <p>
      Install and manage custom Proton installationsw easily and quickly.
    </p>
  </description>

  <launchable type="desktop-id">dev.bitsofabyte.proto.desktop</launchable>

  <screenshots>
    <screenshot type="default">
      <caption>The Main CLI Page</caption>
      <image>https://github.com/BitsOfAByte/proto/blob/main/.assets/Screenshots/main_app_screenshot.png</image>
    </screenshot>
  </screenshots>

  <url type="homepage">http://github.com/BitsOfAByte/proto</url>
  <developer_name>BitsOfAByte</developer_name>

  <provides>
    <binary>proto</binary>
  </provides>
</component>`

	createBuildFile("dev.bitsofabyte.proto.metainfo.xml", fileData)
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
