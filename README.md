# Made - Automate all your scripts


The **made** program is a command line tool that automates the execution of commands defined in Madefiles and Makefiles, searching for these files in various directories, concatenating their sections, and producing a shell script for execution. It extends the functionality of make by reading not only the **Makefile**/**Madefile** in the current directory but also Madefiles in different locations within the system. 

> Even if make is an excellent building tool, I only use it to remember long commands.
> 
> _Guillermo Álvarez_ author of the made command

<details open>
  <summary>Table of Contents</summary>

<!-- vscode-markdown-toc -->
* [What is made command](#Whatismadecommand)
* [Why?](#Why)
* [The `Madefile`](#TheMadefile)
* [The command](#Thecommand)
* [Where to put the madefiles](#Wheretoputthemadefiles)
* [Install](#Install)
* [Update](#Update)

<!-- vscode-markdown-toc-config
	numbering=false
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

</details>

## Installation

Download the binary for your system from github and place it within your path.



## Define scripts everywere

Contrary to `make`, `made` not only reads targets from the Makefile in the current directory.

It reads:

* `Madefile`
* `madefiles/*.made`
* `.madefiles/*.made`

And it also loads `Madefiles` from different parts of the system.

* `Current directory`
* Each directory up until it reaches your HOME or /
* `~/.config/madefiles`
* `~/.local/share/madefiles`
* `/etc/madefiles`
* `/var/lib/madefiles`

## <a name='TheMadefile'></a>The `Madefile`

A madefile looks like this:

```Make
# This is a comment
# And the header of the Madefile
# Every thing will be appended to the generated script

: ${BRANCH:?master}  # Assign a default value to the variable BRANCH

# Now we have the tasks
# they have a name semicolon dependencies and an optional comment prepended with #

deploy: check_clean # Deploy 
	git push dokku $BRANCH

staging: # Configuration for staging environment
	BRANCH=staging

check_clean:
	if ! $(git diff --exit-code > /dev/null) ; then
		echo "Can't operate while there are non commited changes." >&2
		exit 2
	fi 
  
notify: # Notify in the chat 
	curl https://somenotificationservice.com/new_deploy
	
run: # Run a command in the server
	ssh dokku $*
```

Now let's run it:

```
$ made staging deploy notify
... and it deploys stagings and notifies
```


## <a name='Thecommand'></a>The command

```shell
$ made # To display all the tasks
deploy  Deploy
staging Configuration for staging environment
notify  Notify in the chat

$ made deploy # To run a single task
...

$ made -s notify # Output the generated script
: ${BRANCH:?master}  # Assign a default value to the variable BRANCH
curl https://somenotificationservice.com/new_deploy

$ BRANCH=production made deploy  # Pass variables to made
...

$ made run -- help # Pass arguments to the script
...


```


## <a name='Wheretoputthemadefiles'></a>Where to put the madefiles

*madecommand* searches for Madefiles in:
All directories from current up to HOME.

It looks for files called `Madefile` or with `.made` extension, in the directory or inside a `.made` directory.

## <a name='Install'></a>Install

Visit the [releases page](http://github.com/madecommand/made/releases/latest) download and unpack the binary in `/usr/local/bin`

## <a name='Update'></a>Update

Just run `made --update`




## <a name='Why'></a>Why?



Whenever I wanted `make` to do more complex scripting, I spent a lot of time because `Makefiles` are not shell scripts, so I needed to relearn how to do it the `make` way.

Why can't `make` concatenate the scripts in a file and run it?

