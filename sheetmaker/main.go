package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	cf := flag.String("config", "", "config file name")
	flag.Parse()

	if *cf == "" {
		usage()
		return
	}

	payload, err := ioutil.ReadFile(*cf)
	if err != nil {
		fmt.Printf("Failed to read config file: %v\n", err)
		os.Exit(1)
	}

	var conf Config
	err = yaml.Unmarshal(payload, &conf)
	if err != nil {
		fmt.Printf("Failed to unmarshal config file: %v\n", err)
		os.Exit(1)
	}

	m := Maker{
		Config: conf,
	}

	err = m.Run()
	if err != nil {
		fmt.Printf("Sheet maker error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Done")
}

func usage() {
	fmt.Printf(
		`Sprite sheet maker tool. Combines several png files into sprite sheet, with configuration yaml file.

Usage: %v --config config.yaml
config.yaml example:
=== config.yaml begin ===
name: sprites.yaml    # default sprites.yaml
sheets:
- name: player.png    # required
  width: 1024         # default 1024
  uniformDelayMS: 100 # default 100
  files:
  - name: player-idle.png # file name of the sprite frames stripe
    xOrigin: 16   # default 0
    yOrigin: 32   # default 0
    frameCount: 2 # defaults to number of square sprites based on ratio of image width to height; override for non-square sprites
    frameMS: 100  # defaults to sheets.uniformDelayMS
    animLoop:     # default false
    animNext:     # default ""
  - name: player-walk.png
    xOrigin: 16
    yOrigin: 32
    frameCount: 6
    animLoop: true
- name: ...
  files:
  - ...
embeds: # these sprite sheets are added as is to the resulting config file
- name: terrain.png
  sprites:
    dirt:
      width: 32
      height: 32
      xOffset: 0
      yOffset: 0
      frameCount: 10
    grass:
      width: 32
      height: 32
      xOffset: 0
      yOffset: 32
      frameCount: 10
=== config.yaml end ===
`,
		os.Args[0])
}
