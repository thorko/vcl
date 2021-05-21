# vcl
vault cli client

## Install
just run `make`

will install `vcl` in current directory

`make clean` will delete `vcl` file

`make test` will give you the version



## Run it

```bash
usage: vcl --url=URL --secret=SECRET [<flags>]

Flags:
      --help           Show context-sensitive help (also try --help-long and --help-man).
  -u, --url=URL        url to vault
  -s, --secret=SECRET  secret to get from vault
  -t, --token=TOKEN    token to authorize at vault
  -n, --user=USER      user for vault
  -p, --pass=PASS      password for vault
  -k, --key=KEY        the key in secret to get
      --version        Show application version.
```

#### Example:

```bash
# if you have a selfsigned certificate
export VAULT_SKIP_VERIFY=1
./vcl -u https://myvault:8200 -n thorko -s secrets/data/backup -p mypassword -k password
```

