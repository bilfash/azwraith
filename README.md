# Azwraith

Azwraith is a cli command to manage credential when committing your changes to version control system.

If you are working with multiple credential on git or multiple account on different git domain 
you often forget to change credential. For some people it is not a big deal, but we don't want to have
commit with unknown name. Azwraith prevent that to happen by matching git remote url with azwraith config.

### Requirement

- Git CLI
    
    Follow this step to install git cli [link](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) 
    
- Golang
    
    `Version >  1.11.4`, follow this step to install golang [link](https://golang.org/doc/install)

### How to install

```
go get github.com/bilfash/azwraith
```

### Usage

##### Overall

```
% azwraith help

Azwraith is a cli command to manage credential when pushing your changes to version control system

Usage:
  azwraith [command]

Available Commands:
  commit      Commit your code
  config      Run azwraith config related command
  ensure      Ensure azwraith config is working as expected
  help        Help about any command

Flags:
  -h, --help   help for azwraith

Use "azwraith [command] --help" for more information about a command.
```
##### Config
Azwraith store its configuration on  `~/.azwraith`, azwraith config consist of list :
- name : git config username
- email : git config email
- pattern : git remote url pattern to determine which config to use by matching git remote url to url pattern

```
% azwraith config -h

Run azwraith config related command

Usage:
  azwraith config [command]

Available Commands:
  add         Add config
  delete      Delete config given ID
  get         Get all config

Flags:
  -h, --help   help for config
```
##### Ensure
Azwraith provide feature to make sure your configuration working properly and return config used.
```
% azwraith ensure -h

Ensure will match remote url from command argument to current azwraith config. This will help you to make sure your 
azwraith config is working as expected

Usage:
  azwraith ensure [flags]

Flags:
  -h, --help         help for ensure
  -u, --url string   git remote url
```
`-u` flags is mandatory.

##### Commit
Commit command will get remote url and match it with azwraith config, after getting the right config azwraith will 
commit your code using credential from matched config
```
Usage:
  azwraith commit [flags] <commit message>

Flags:
  -h, --help            help for commit
  -r, --remote string   specify which git remote is used (default "origin")

Example:
  azwraith commit "This is proper commit message"

```

### Project Status
Development

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)