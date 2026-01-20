# suckless-dwmblocks

A lightweight [dwm](https://dwm.suckless.org/) status bar updater written in **Go**.
It updates the status bar at fixed intervals.

## Features

* Interval updates per block
* Modular block functions
* Writes to dwm status via X11 (fallback to xsetroot)

## 📦 Installation

```bash
git clone https://github.com/zetatez/suckless-dwmblocks.git
cd suckless-dwmblocks
make install
```

## 🚀 Usage

1. Define your blocks in `./config.go`:

   * `Interval`: update frequency
   * `Func`: Go function callback (optional)

2. Run it in the background:

   ```bash
   ./dwmblocks &
   ```

## ⚙️ How it works

* Each block is updated by interval
* Results are concatenated with a delimiter
* The final string is written to the X root window via `xsetroot -name`
* `dwm` displays it in the status bar
