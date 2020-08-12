# Contributing

## Making Issues
This project can ALWAYS be improved. If you come across any bugs, have any questions or dream up a new feature for Odin then you should [open an issue](https://github.com/theycallmemac/odin/issues/new/choose) to talk about it. This is the easiest way to get involved with the project!

## Building Features
If you want to help build out Odin's capabilities, feel free to join the discussion on any issues. I'd be more than happy to assign new people to issues on the project.

If you do want to contribute by enhancing Odin's feature set, please consult the section below on:
	- Commit Message Format
	- Code Linting
	- Code Formatting
	- Running Github Actions Locally

#### Commit Message Format

The format of all commit messages should correspond to the following format:

```
<type>: <message>
```

Where `<message>` clearly explains why this change is being made and `<type>` corresponds to one of the following:

- ADD  (used for a new feature)
- FIX  (used for a bug fix)
- REMOVE  (used to delete feature)
- REFACTOR  (used when refactoring code)
- STYLE  (used for lint fixes, formatting, missing semi colons, etc; no code change)
- DOCS  (used for changes to documentation)
- TEST  (used when adding or refactoring tests)
- OTHER  (other, explain in the body)

#### Linting

The following tools are used for code linting:
	- Go - [golint](https://github.com/golang/lint)
	- Node.js - [eslint](https://github.com/eslint/eslint)
	- Python - [pylint3](https://github.com/golang/lint)
	- Bash - [shellheck](https://github.com/koalaman/shellcheck)

#### Formatting

The following tools are used for code formatting:
	- Go - [gofmt](https://golang.org/cmd/gofmt/)
	- Node.js - [prettier](https://github.com/prettier/prettier)
	- Python - [black](https://github.com/psf/black)
	- Bash - [beautysh](https://github.com/papampi/beautysh)

#### Running Github Actions Locally

We strongly recommend running your code through our Github Actions before opening a Pull Request. Running our Github Actions locally can be accomplished by using [act](https://github.com/nektos/act).


------------
