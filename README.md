# CoffeeMD

`coffeemd` is a tool to make changes to files in an an Obsidian vault after
you've created the vault with [t2md](https://github.com/GSGBen/t2md) and have already
started using it. Changes that would have been good to make in t2md, but are now
too late to make there because you're already using the vault.

For example, the first task I wrote this for was to convert the manual and
too-heavy converted card headers to frontmatter. From:

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

TODO. Something like `coffeemd -help` to start.

## TODO:

- test nested files
- test a file with emoji in the title
- write tests for parsing and other new functions
- run on full vault