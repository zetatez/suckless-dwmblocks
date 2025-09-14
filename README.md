# suckless-dwmblock

A lightweight [dwm](https://dwm.suckless.org/) status bar updater written in **Go**.
It updates the status bar at fixed intervals or on demand via Unix signals.

## ✨ Features

* 🕒 **Interval updates**: update each block at a configurable time interval
* 📡 **Signal updates**: refresh specific blocks instantly via signals
* 🧩 **Modular design**: each block can have its own interval, signal, icon, and command/function
* ⚡ **Lightweight & fast**: pure Go implementation with minimal dependencies

## 📦 Installation

```bash
git clone https://github.com/zetatez/dwm-statusbar-go.git
cd suckless-dwmblock
make install
```

## 🚀 Usage

1. Define your blocks in `./config.go`:

   * `Interval`: update frequency in seconds (`0` = never update automatically)
   * `Signal`: signal ID to trigger manual updates
   * `Icon`: prefix icon/string for the block
   * `Command`: shell command to run (optional)
   * `Func`: Go function callback (optional)

2. Run it in the background:

   ```bash
   ./dwmblock &
   ```

3. Trigger a block update via signal:

   ```bash
   kill -10 $(pidof dwmblock)   # refresh block with signal ID = 10
   ```

## ⚙️ How it works

* Each block is updated either by interval or signal
* Results are concatenated with a delimiter
* The final string is written to the X root window via `xsetroot -name`
* `dwm` displays it in the status bar

