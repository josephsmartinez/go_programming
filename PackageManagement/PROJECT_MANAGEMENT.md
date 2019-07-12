# Project Management with GO

## Options

Either way the actual source code for your project needs to be placed in the *src subdirectory (e.g., $GOPATH/src/your_project or $GOPATH/src/github.com/your_github_username/your_project).*

Go v1.1 does introduce the concept of *modules* that allows you to have your project code outside of your GOPATHwithout the import restrictions mentioned above, but it’s still an experimental feature at this point in time.

## Set up environment variables

export GOPATH=/Users/Uday.Hiwarale/uday-gh/go_workspaces/main
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOPATH:$GOBIN
export PATH

## Resources

Go Project Layout
https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2