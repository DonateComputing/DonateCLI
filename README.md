# Author + Runner CLI

Part of a larger Donate Project

## Features

### For Runner:

* [ ] `ps` lists all jobs currently being run on client
* [ ] `hub ps` lists all jobs waiting for runner
	- [ ] `-a` include jobs being run or completed
	- [ ] `-m` filter to only jobs belonging to or run by user
* [ ] `start <id>` checks out and runs job with id
* [ ] `stop` stops all jobs and returns them to hub
	- [ ] `stop <id>` stops only job with given id
* [ ] `pause` pauses all jobs
	- [ ] `pause <id>` pauses only job with given id
* [ ] `unpause` unpauses all paused jobs
	- [ ] `unpause <id` unpauses only job with given id
* [ ] `prune` autoremove all 'completed' jobs

### For Author:

* [ ] `add` opens interface to guide thru making new job
	- [ ] `add <id> <description>` will create job directly
* [ ] `rm <id>` removes job from hub

### For Both:

* [ ] `login` opens interface to create login settings file
