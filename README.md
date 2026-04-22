### lctr

**Installation**

1. Clone the repo
```
git clone git@github.com:infinitej3ss/lctr.git
```

2. Enter the lctr folder and build
```
cd lctr
go build
```
This should give you an executable called **lctr**

**Usage**

Run the command
```
lctr <directory>
```

You should see the output `Total Lines: XX` where `XX` is some integer. Note that lctr only counts lines with content, making it useful for counting lines of code. It also excluded dotfiles and binaries (files without an extension).

I have an alias on my machine so I can run lctr easily from wherever I'm at, and you should too!

To set up an alias, add a line like this to your shell config file (`~/.zshrc` for zsh, `~/.bashrc` for bash)
```
alias lctr='~/path/to/lctr'
```

**Feature Checklist**

- [ ] Flag to ignore comments
