# Holla

A golang tool for checking liveliness of HTTP reachable targets. 

## General architecture

### Scheduler

The scheduler is in charge of triggering the liveliness check process for a TargetGroup based on its cron definition

### Group

A group is a group of targets that share some similarity.
A group has a name, a list of Targets and a cron schedule

### Target

A target represents an HTTP Reachable target. 
It consists of a Name and a URL.

### TaskManager

The taskmanager is responsible for executing the tasks within a group, collecting its results and reporting on them

### Task

A task is a single instance of a liveliness check. Its sole purpose is performing the liveliness check and reporting on its results

### Storer

A storer has the responsibility to persist results, so we can later check trends based on past results

