<div align="center">

# 🧬 Telomere

**Workspaces with an expiration date.**

Create a scratch directory, give it a TTL, and forget about it —<br>
Telomere deletes it when the time is up.

</div>

---

## 📦 Install

Requires Go 1.26.6 or later and a C compiler (`gcc` or `clang`), since the SQLite driver is cgo-based.

```sh
git clone https://github.com/valusun/Telomere.git
cd Telomere
./setup.sh
```

> Installs `telomere` into `~/.local/bin` and initializes `~/.telomere`.

## 🚀 Usage

```sh
# create a workspace that lives for 3 days
telomere new scratch --ttl 3d

# jump into it
cd "$(telomere path scratch)"

# see what's left
telomere list

# delete it now
telomere kill scratch
```

## ♻️ Auto cleanup

Expired workspaces are removed by `telomere gc`. Run it hourly with systemd:

```sh
mkdir -p ~/.config/systemd/user
cp contrib/systemd/telomere-gc.* ~/.config/systemd/user/
systemctl --user enable --now telomere-gc.timer
```

## 📄 License

Apache-2.0
