# Author + Runner CLI

![Go Tests](https://github.com/DonateComputing/DonateCLI/workflows/Go%20Tests/badge.svg?branch=master)

Part of a larger Donate Project

## Features

### For Runner

* [x] `ps` lists all jobs currently being run on client
  * [x] `-a` include completed jobs on client
* [x] `hub` and `hub ps` list all jobs waiting for runner
  * [x] `-a` include jobs being run or completed
  * [x] `-u` filter to only jobs belonging to or run by user
* [x] `start <username> <title>` checks out and runs job with given id
* [x] `stop` stops all jobs and returns them to hub
  * [x] `stop <username> <title>` stops only job with given id
* [x] `pause` pauses all jobs
  * [x] `pause <username> <title>` pauses only job with given id
* [x] `unpause` unpauses all paused jobs
  * [x] `unpause <username> <title>` unpauses only job with given id
* [ ] `prune` autoremove all 'completed' jobs

### For Author

* [ ] `add` opens interface to guide thru making new job
  * [x] `add <title> <description> <image>` will create job directly
* [ ] `rm <id>` removes job from hub

### For Both

* [x] `login <username> <password>` creates login settings file
  * [x] `--register` flag will automatically register user on public hub
* [x] `help` `--help` or `-h` will display help / usage message
  * [x] `` empty command will display help / usage message
