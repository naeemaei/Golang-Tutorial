# Golang-Tutorial

## Session 2

**Tools**:

### Install Go Language 1.18

#### Windows:

[Download here](https://go.dev/dl/go1.18.1.windows-amd64.msi) and then install

#### Linux Method 1:

Run this commands in terminal:

    sudo apt update
    sudo apt search golang-go
    sudo apt search gccgo-go
    sudo apt install golang-go

#### Linux Method 2:

1. [Download file](https://go.dev/dl/go1.18.1.linux-amd64.tar.gz)
2. Extract file `tar -xf "go1.18.1.linux-amd64.tar.gz"`
3. Setup permissions `sudo chown -R root:root ./go`
4. Move go binaries to local folder `sudo mv -v go /usr/local`
5. for Setup go path, open profile: `code ~/.bash_profile` or `code ~/.profile`
6. Append following to lines to file:
   ```
   export GOPATH=$HOME/go
   export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
   ```
7. source this file: `source ~/.bash_profile` or `source ~/.profile`
8. Run `go version` to check if go is installed

### Install Go Debugger

Run this command: ` go install github.com/go-delve/delve/cmd/dlv`

### Install Visual Studio Code

[Download here](https://code.visualstudio.com/download)

### Install Git

#### Windows:

[Download frome here](https://gitforwindows.org/) and then install

#### Linux:

Run this commands in terminal or see this [link](https://git-scm.com/download/linux):

    sudo apt update
    sudo apt-get install git

### VS Code Extensions

Run this commands in terminal

```
code --install-extension aaron-bond.better-comments
code --install-extension christian-kohler.path-intellisense
code --install-extension ckolkman.vscode-postgres
code --install-extension cweijan.vscode-redis-client
code --install-extension donjayamanne.githistory
code --install-extension eamodio.gitlens
code --install-extension esbenp.prettier-vscode
code --install-extension formulahendry.code-runner
code --install-extension golang.go
code --install-extension HookyQR.beautify
code --install-extension IBM.output-colorizer
code --install-extension mhutchie.git-graph
code --install-extension ms-azuretools.vscode-docker
code --install-extension PKief.material-icon-theme
code --install-extension premparihar.gotestexplorer
code --install-extension rangav.vscode-thunder-client
code --install-extension redhat.vscode-yaml
code --install-extension streetsidesoftware.code-spell-checker
code --install-extension TabNine.tabnine-vscode
code --install-extension VisualStudioExptTeam.vscodeintellicode
code --install-extension vscode-icons-team.vscode-icons
code --install-extension wmaurer.change-case
```
