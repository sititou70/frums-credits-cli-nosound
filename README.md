# frums-credits-cli-nosound

![image](https://user-images.githubusercontent.com/18525488/115956294-fd372a00-a536-11eb-8722-897f84f4c543.png)

[Frums / Credits BGA (Back Ground Animation)](https://www.youtube.com/watch?v=EOTAWLaDa58) on Linux tty

## Quick Demo

Download [latest binary](https://github.com/sititou70/frums-credits-cli-nosound/releases) and execute.

note: You probably need to set the `RUNEWIDTH_EASTASIAN` environment variable to `0`.  
Linux example: `RUNEWIDTH_EASTASIAN=0 ./frums-credits-cli`

### Options

```
$ ./credits -h
Usage of ./credits:
  -s int
        time to skip play (sec)
  -v    print extra information
```

## Requirements

- [Go (1.16.3~)](https://golang.org/doc/install)
- [hajimehoshi/oto's requirements](https://github.com/hajimehoshi/oto#prerequisite)
  - On Ubuntu or Debian: `apt install libasound2-dev`
- (optional): [FullCyrSlav-TerminusBoldVGA16](https://www.zap.org.au/projects/console-fonts-distributed/psftx-ubuntu-20.04/FullCyrSlav-TerminusBoldVGA16.psf)
  - I recommend using FullCyrSlav-TerminusBoldVGA16 with tty.
  - If you don't have `FullCyrAsia-TerminusBoldVGA16.psf` or` FullCyrAsia-TerminusBoldVGA16.psf.gz` in `/usr/ share/consolefonts/`, download and place the font.

## Usage

- (optional): Replace sound file
  - `frums-credits-cli/credits-csf/credits.mp3` is a silent dummy file in consideration of copyright.
  - You can play music on the CLI by replacing this file with legally obtained music data and adjusting the offset in `frums-credits-cli/credits-csf/meta.yaml`.
- (optional): Switch to tty and setfont
  - On linux, press <key>Ctrl</key> + <key>Alt</key> + <key>F[2 ~ 6]</key> and login
  - setfont
    - `cd /usr/share/consolefonts/`
    - `sudo setfont FullCyrAsia-TerminusBoldVGA16.psf.gz` or `sudo setfont FullCyrAsia-TerminusBoldVGA16.psf`
  - `cd` (to $HOME)
- `git clone [this repo's url]`
- `cd frume-credits-cli`
- `make`
- `RUNEWIDTH_EASTASIAN=0 ./dist/credits`

## CSF: Credits Score Format

CSF is a format for text-based music videos like [Frums / Credits BGA](https://www.youtube.com/watch?v=EOTAWLaDa58). For details, please refer to the [CSF specifications (Japanese)](./docs/csf_spec.ja.md)

## Cross build on docker

`make cross`

## Licence

MIT
