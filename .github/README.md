# lt-cli

For a lighting talk event.

# Installation

```
go install github.com/ppluuums-jp/lt@latest
```

# Usage

Note:

- default `number` is 1
- any names match `*.txt` will apply
- [the example `*.txt`](https://github.com/ppluuums-jp/lt-cli/blob/main/applicants.txt)

```
$ lt pick [flags]

$ lt pick -n 2 Bob Alice Natasha
  - Natasha
  - Bob

$ lt pick -n 2 -t applicants.txt
  - Natasha
  - Bob
```

```
$ lt shuffle [flags]

$ lt shuffle Bob Alice Natasha
  No 1. Alice
  No 2. Bob
  No 3. Natasha

$ lt shuffle -t applicants.txt
  No 1. Alice
  No 2. Bob
  No 3. Natasha
```

# Licence

See [here](https://github.com/ppluuums-jp/lt/blob/main/LICENSE).
