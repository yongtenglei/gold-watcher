# Gold Watcher

**_Tracing the gold price and make money!_**

Supports _Linux_ and _Windows_ now!

## For user

### Linux Users

Go and download the release version for Linux, extract it and have a look at **Makefile**.

Normally, you just need:

```bash
sudo make install
```

To uninstall, you may:

```bash
sudo make uninstall
```

### Windows Users

Just download the released .exe file, run and play!

## For developer to test

```bash
git clone https://github.com/yongtenglei/gold-watcher.git
go mod tidy
env DB_PATH=./sql.db go run .
```
