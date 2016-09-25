package main

const (
	werckerYaml = `
#
# $ docker run -it --rm golang:latest uname -r
# 4.4.20-moby
box: golang

build:
  steps:
    - add-to-known_hosts:
      hostname: github.com
      fingerprint: 16:27:ac:a5:76:28:2d:36:63:1b:56:4d:eb:df:a6:48
    - script:
      name: ssh config
      code: >
        ssh -t git@github.com 2>&1 > /dev/null || true
    - script:
      name: go get glide
      code: >
        go get github.com/Masterminds/glide &&
        go install github.com/Masterminds/glide
    - script:
      name: go get qmkmk
      code: >
        go get github.com/ReSTARTR/qmkmk &&
        cd $(go env GOPATH)/src/github.com/ReSTARTR/qmkmk &&
        glide update &&
        go install github.com/ReSTARTR/qmkmk
    - script:
      name: init tool
      code: >
        qmkmk init &&
        apt-get update &&
        sudo apt-get install -y gcc unzip wget zip gcc-avr binutils-avr avr-libc dfu-programmer dfu-util gcc-arm-none-eabi binutils-arm-none-eabi libnewlib-arm-none-eabi
    - script:
      name: build original keymap
      code: >
        qmkmk -keyboard={{.Keyboard}} -subproject={{.Subproject}} -keymap={{.Keymap.Name}} -owner={{.Owner}} get &&
        qmkmk -keyboard={{.Keyboard}} -subproject={{.Subproject}} -keymap={{.Keymap.Name}} -owner={{.Owner}} build

    # deploy steps
    - script:
      name: init tool
      code: >
        mkdir $HOME/bin &&
        sudo wget https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64 -O $HOME/bin/jq &&
        chmod +x $HOME/bin/jq
    - script:
        name: put built hex file to bintray.
        code: >
          HEX_FILE="$(qmkmk -keyboard={{.Keyboard}} -subproject={{.Subproject}} -keymap={{.Keymap.Name}} config | $HOME/bin/jq -r '.|.hex_dir+"/"+.keymap.hex_file')" &&
          HEX_NAME="$(qmkmk -keyboard={{.Keyboard}} -subproject={{.Subproject}} -keymap={{.Keymap.Name}} config | $HOME/bin/jq -r '.|.keymap.hex_file')" &&
          curl -T $HEX_FILE -u$BINTRAY_USERNAME:$BINTRAY_APIKEY \
            https://api.bintray.com/content/$BINTRAY_USERNAME/generic/ergodox.hex/$WERCKER_GIT_BRANCH-$WERCKER_GIT_COMMIT/$HEX_NAME
`
)
