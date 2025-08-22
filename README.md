<!--
SPDX-FileCopyrightText: 2025 James Pond <james@cipher.host>

SPDX-License-Identifier: CC0-1.0
-->

# `taskemon`

![A photo of a board with several tasks printed and
displayed.](.github/assets/taskemon-task-example.jpg "Task board")

`taskemon` is a very simple CLI tool that prints daily task tickets on a
thermal printer, complete with a QR code that links to a random Pokémon
as a "reward" for motivation. Basically, it's part of a process that
helps you apply feedback loops to your tasks.

I was inspired by [this article by Laurie
Hérault](https://www.laurieherault.com/articles/a-thermal-receipt-printer-cured-my-procrastination)
which explains how understanding the science of video games helped them
"cure" their procrastination—you're encouraged to read the article for a
better understanding. The Pokémon part came from [this comment by
PaulHoule on Hacker
News](https://news.ycombinator.com/item?id=44257382).

> [!IMPORTANT]  
> `taskemon` only supports the Epson TM-T20II and TM-T20III thermal printers via USB
> right now, and was only tested on macOS Sequoia. Patches or PRs to
> support other models or networked printers are welcome!

## Features

- Prints styled task tickets with owner, category, and task description.
- Generates a QR code linking to a random Pokémon Pokédex entry, or,
  rarely, a fun surprise.
- Supports the Epson TM-T20II and TM-T20III thermal printers via USB out of the box.

## Installation

### From source

First install the dependencies:

- Go 1.24 or above.
- make.
- libusb

Then, clone the repository, switch to the latest stable tag, compile,
and install:

```bash
git clone 'git.sr.ht/~jamesponddotco/taskemon'
cd 'taskemon'
git checkout 'v0.1.0'
make
sudo make install
```

## Usage

```bash
$ taskemon --help
Usage of taskemon:
  -model string
    	the thermal printer model (TM-T20X-II or TM-T20III) (default: TM-T20X-II)
  -nopokedex
    	if set to true, no Pokedex QR code will be included (default: false)
  -owner string
    	the person responsible for the task
  -task string
    	the task description
```

## Contributing

Anyone can help make **taskemon** better.  Check out [the contribution
guidelines](CONTRIBUTING.md) for more information.

## Resources

The following resources are available:

- **Support and general discussions**:
  [Sourcehut](https://lists.sr.ht/~jamesponddotco/taskemon-discuss)/[GitHub](https://github.com/jamesponddotco/taskemon/discussions).
- **Patches and PRs**:
  [Sourcehut](https://lists.sr.ht/~jamesponddotco/taskemon-devel)/[GitHub](https://github.com/jamesponddotco/taskemon/pulls).
- **Feature requests and bug reports**:
  [Sourcehut](https://todo.sr.ht/~jamesponddotco/taskemon)/[GitHub](https://github.com/jamesponddotco/taskemon/issues).

---

The work in this repository complies with [the REUSE
specification](https://reuse.software/spec-3.3/). While [the default
license is EUPL-1.2](LICENSE.md), individual files may be licensed
differently.

Please see the individual files for details and [the LICENSES
directory](LICENSES/) for a full list of licenses used in this
repository.
