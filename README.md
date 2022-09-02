# Made

Automate all your scripts

## Why?

Even if make is an excellent building tool, I only use it today to run some tasks like start development, deploy, compile, etc., without remembering the peculiarities of each project.

Whenever I wanted `make` to do more complex scripting, I spent a lot of time because `Makefiles` are not shell scripts, so I needed to relearn how to do it the `make` way.

Why can't `make` concatenate the scripts in a file and run it?

## Welcome to *madecommand*

*made* is a command line tool that executes commands defined in `Madefiles`.

The madefiles are text files that contain shell scripts grouped in tasks similar to a Makefile.

When you run *made* with some targets, *madecommand* will concatenate your targets and their dependencies into a single script and run it.

## The `Madefile`

A madefile looks like this:

```
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


## The command

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


## Where to put the madefiles

*madecommand* searches for Madefiles in:
All directories from current up to HOME.

It looks for files called `Madefile` or with `.made` extension, in the directory or inside a `.made` directory.

## Install

Visit the (releases page)[http://github.com/madecommand/made/releases/latest] download and unpack the binary in `/usr/local/bin`

## Update

Just run `made --update`




