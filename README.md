# Telomere

Workspaces with an expiration date.

Create a scratch directory, give it a TTL, and forget about it — Telomere deletes it when the time is up.

## Install

```sh
git clone https://github.com/valusun/Telomere.git
cd Telomere
./setup.sh
```

Installs `telomere` into `~/.local/bin` and initializes `~/.telomere`.

## Usage

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

```
NAME      CREATED      EXPIRES      TELOMERE
scratch   2026-08-16   2026-08-19   ███████░░░  70%
```

## Auto cleanup

Expired workspaces are removed by `telomere gc`. Run it hourly with systemd:

```sh
mkdir -p ~/.config/systemd/user
cp contrib/systemd/telomere-gc.* ~/.config/systemd/user/
systemctl --user enable --now telomere-gc.timer
```

## License

Apache-2.0
