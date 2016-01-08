# label-all-the-things

Batch create labels on GitHub repositories.


```
$ ./label-all-the-things 
NAME:
   label-all-the-things - Manipulate GitHub labels

USAGE:
   label-all-the-things [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   create       create labels from a definition file
   help, h      Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --token              GitHub API token
   --token-file         GitHub API token file
   --help, -h           show help
   --version, -v        print the version
```

## Example usage

```
$ ./label-all-the-things --token-file ~/.github-token create --labels common.json icecrime/label-all-the-things
```
