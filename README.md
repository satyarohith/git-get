# git-get

Clone repositories like `go get`.

## Installation

**Note:** `git-get` is currently supported on MacOS & Linux.

If you've `golang` installed, run the following commands:

```sh
$ go get github.com/satyarohith/git-get
$ go install github.com/satyarohith/git-get
```

If not, you can download the binaries from [here](https://github.com/satyarohith/git-get/releases).

## Usage

### Clone repositories hosted on GitHub

```sh
$ git get <username>/<repository_name>
```

Example:

```sh
$ git get satyarohith/shark # Clones to `~/c/github.com/satyarohith/shark`
```

### Clone repositories hosted under your GitHub account

```sh
$ git get <repository_name> # Clones to `~/c/github.com/<username>/<repository_name>`
```

`git-get` reads your GitHub username from `.gitconfig` file.

Append the below to your `.gitconfig`. Commonly found at `~/.gitconfig`.

```
[github]
    user = <your_username>
```

### Clone any git repositories

```sh
$ git get https://gitlab.com/satyarohith/example # Clones to ~/c/github.com/satyarohith/example
```

## License

MIT © [Satya Rohith](https://satyarohith.com)
