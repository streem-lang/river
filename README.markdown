
# River: Streem's package manager
This is a small, GitHub based package manager for Streem, called River. It is very similar to GoLang's package manager. Since Streem is not finished yet, River will not be fully functional till Streem hits a version where it can execute code. River is written in Go. It's commands include:

  * river **install** _author/name_: Installs Streem package with same name and author from github.
  * river **install**: Installs Streem package that you are in
  * river **remove** _name_: Removes streem package that you name
  * river **setup**: Interactive prompt to setup your project
  * river **run**: Runs your project in the `src/main.strm`
  * river **bin**: Creates an executable of your project in the `src/main.strm`
  * river **version**: Version
  * river **help**: A help similar to this
