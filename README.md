# SOREC - Sonos Refurbished Checker

![License](https://img.shields.io/badge/license-GPL3-green?style=plastic)
![Go](https://img.shields.io/github/go-mod/go-version/thomasboom89/sorec/main)

A cli scraper to check price and availabilty of sonos refurbished products

## Installation

Installation can be different, depending on what OS you are on. Guide is written for linux x86 OS.
\
Requirements: installed working go environment on build machine >= 1.20

Clone the repo and build binary

```zsh
git clone https://github.com/ThomasBoom89/sorec
cd sorec/sorec
go build -o ../build/sorec -ldflags "-w -s"
```

Now you can use it directly from the build folder

```zsh
cd ../build
./sorec
```

## Usage

SOREC will always try to help you if you don't know what to do or what to type just append the -h or --help flag.
\
Just start with the plain ```sorec``` command and it will show the possible commands and flags. A good start is to list
all possible locale, to find your specific shop for your country
```sorec list locale```. Let's say you are from Germany so your locale will be "de-de". You can now use this to check
all possible products by typing ```sorec check -l de-de```. You will see, depending on the available products, which
products are checked and when finished you will get a table with all products including prices and current stock.

## Contribution

Contributions are always welcome.
\
If you find a bug, report it. If you have question, ask for help (in Disskussions). If you want to improve something,
make a PR. Also, there are several points that could be worked off to help other users to get SOREC running:

- ~~Check installation on linux x86 (report also your version) and add a section to the installation guide~~
- Check installation on Windows (report also your version) and add a section to the installation guide
- Check installation on macOS (report also your version) and add a section to the installation guide
- Check installation on linux arm (report also your version) and add a section to the installation guide
- Generate GitHub action to build binaries for all working OS at release

## License

SOREC - Sonos Refurbished Checker, a cli scraper to check price and availabilty of sonos refurbished products\
Copyright (C) 2023 ThomasBoom89. GNU GPL3 license (GNU General Public License Version 3).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
License as published by the Free Software Foundation, either version 3 of the License, or (at your option)
any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see:

1. The LICENSE file in this repository.
2. https://www.gnu.org/licenses/.
3. https://www.gnu.org/licenses/gpl-3.0.txt.

However, SOREC - Sonos Refurbished Checker includes several third-party Open-Source libraries, which are licensed under
their own respective Open-Source licenses. Libraries or projects directly included:

- getoutreach/goql: [APACHE-2.0](https://github.com/getoutreach/goql/blob/main/LICENSE)
- olekukonko/tablewriter: [MIT](https://github.com/olekukonko/tablewriter/blob/master/LICENSE.md)
- spf13/cobra: [APACHE-2.0](https://github.com/spf13/cobra/blob/master/LICENSE.txt)
