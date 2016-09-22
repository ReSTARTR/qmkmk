mkmg
====

`mkmg `Helps you to manage keymaps of qmk\_firmware.


- Easy to follow `jackhumbert/qmk_firmware` repository
  - separate repositories of building tool and keymap.
- Easy to version your firmware(`firmware-v1.0.0.hex`)
  - manage versions of keymaps with git on your repository
- Easy to share your firmware (via werker and bintray)
  - enable to download firmware always without any build tools
- Easy to build your firmware continously(check dependencies are breaking)
  - build your keymap with werker (CI service)

Setup
----

Before do this instruction, you must fork the repository of jackhumbert/qmk_firmware.

- github.com/`$YOURNAME/qmk_firmware-ergodox-ez`

```
$ git clone https://github.com/jackhumbert/qmk_firmware
```

Usage
----

```bash
$ cd /path/to/qmk_firmware

# Add your remote repository, and make your branch
$ mkmg init remote add ${KEYMAP}

# Make dir of your keymap, and copy files from default.
$ mkmg create ${KEYMAP}

# Open keymap.c in your favorite editor
$ mkmg edit ${KEYMAP}

# Build
$ mkmg build ${KEYMAP}

# Push to your repository, and run ci at werker.
$mkmg push ${KEYMAP}

# Update repository
$ mkmg update 
```

Mechanism
----

directories

```
|
|- jackhumbert/
|    `- qmk_firmware/
|          `- keyboards/
|              `- ergodox/
|                  |- default/
|                  |-  :
|                  |-  :
|                  `- restartr/ #-> symbolic link to your keymap
|
`- ReSTARTR/
      `- qmk_firmware-ergodox-ez-restartr/
            `- keymap.c
```

commands

```bash
BASEPATH=$HOME/src/github.com #= GOPATH
OWNER=ReSTARTR

KEYBOARD=ergodox
SUBPROJECT=ez
KEYMAP=restartr
BUILD_NAME=${KEYBOARD}-${SUBPROJECT}-${KEYMAP}  #= "ergodox-ez-restartr"
KEYMAP_REPO=${OWNER}/qmk_firmware-${BUILD_NAME} #= "ReSTARTR/qmk_firmware-ergodox-ez-restartr"

# qmk init remote add https://github.com/${KEYMAP_REPO}
git clone https://github.com/jackhumbert/qmk_firmware \
  ${BASEPATH}/jackhumbert/qmk_firmware
git clone https://github.com/${KEYMAP_REPO} \
  ${BASEPATH}/${KEYMAP_REPO}
ln -s ${BASEPATH}/${KEYMAP_REPO} \
  ${BASEPATH}/qmk_firmware/keyboards/${KEYBOARD}/keymaps/${KEYMAP}

# qmk build
cd ${BASEPATH}/jackhumbert/qmk_firmware
make ${BUILD_NAME}

# qmk teensy
make teensy ${BUILD_NAME}

# qmk push
cd ${BASEPATH}/${KEYMAP_REPO}
git checkout -b keymap/${BUILD_NAME}
git commit -m "Update $(date)"
git push origin

# mkmg update 
cd $BASEPATH/jackhumbert/qmk_firmware
git pull origin master

# and build agin
make ${BUILD_NAME}
```