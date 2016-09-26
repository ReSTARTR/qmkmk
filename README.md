qmkmk
====

[![wercker status](https://app.wercker.com/status/e0a6149454fdb9e8599650b1425645ed/s/master "wercker status")](https://app.wercker.com/project/byKey/e0a6149454fdb9e8599650b1425645ed)

`qmkmk` Helps you to manage keymaps of qmk\_firmware.

Plan:

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

```
$ go install github.com/ReSTARTR/qmkmk
```

Usage
----

create your config file into `~/.config/qmkmk.json`.

```json
{
  "keyboard": "ergodox",
  "subproject": "ez",
  "keymap": "restartr",
  "owner": "ReSTARTR"
}
```

and run with subcommand

```bash
# clone qmk_firmware repository
$ qmkmk init

# clone your keymap, and link from qmk_firmware's keymaps
$ qmkmk get

# build keymap
$ qmkmk build

# install keymap into your keyboard
$ qmkmk install

# push keymap to the repository.
# Then, build and release your firmware by CI service.
$ qmkmk push
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
