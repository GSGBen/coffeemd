# CoffeeMD

`coffeemd` is a tool to make changes to files in an an Obsidian vault after you've
created the vault with [t2md](https://github.com/GSGBen/t2md) and have already started
using it. Changes that would have been good to make in t2md, but are now too late to
make there because you're already using the vault.

For example, the first task I wrote this for was to convert the manual and too-heavy
converted card headers to frontmatter. From:

```text
# (emoji) Full Card Name

Original URL: https://trello.com/example

---

(content...)
```

to

```text
---
title: (emoji) Full Card Name
original_url: https://trello.com/example
---

(content...)
```

## usage

Run `coffeemd` to get help. The basic layout is

```
coffeemd [global options] command [command options] [arguments...] 
```

Note that you can set vault-path via environment variable instead:
`COFFEEMD_VAULT_PATH`.

the `--apply` option is required to make actual changes. It's a global option so it goes
before the command you request. Without it, coffeemd just prints what it will affect.

