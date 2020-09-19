# Author + Runner CLI

Part of a larger Donate Project

## Features

### For Runner

* [x] `ps` lists all jobs currently being run on client
  * [ ] `-a` include completed jobs on client
* [ ] `hub` and `hub ps` list all jobs waiting for runner
  * [ ] `-a` include jobs being run or completed
  * [ ] `-u` filter to only jobs belonging to or run by user
* [ ] `start <id>` checks out and runs job with id
* [ ] `stop` stops all jobs and returns them to hub
  * [ ] `stop <id>` stops only job with given id
* [ ] `pause` pauses all jobs
  * [ ] `pause <id>` pauses only job with given id
* [ ] `unpause` unpauses all paused jobs
  * [ ] `unpause <id>` unpauses only job with given id
* [ ] `prune` autoremove all 'completed' jobs

### For Author

* [ ] `add` opens interface to guide thru making new job
  * [ ] `add <title> <description> <image>` will create job directly
    * [ ] `--allow-multiple` on command above will allow multiple runners
* [ ] `rm <id>` removes job from hub

### For Both

* [x] `login <username> <password>` creates login settings file
  * [x] `--register` flag will automatically register user on public hub
* [x] `help` `--help` or `-h` will display help / usage message
  * [x] `` empty command will display help / usage message
