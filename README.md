# GitHub User Activity

## Description
- Complete and simple solution for [roadmap.sh](https://roadmap.sh) backend [gitHub user activity project](https://roadmap.sh/projects/github-user-activity)
- This is a simple command line interface to fetch gitHub user activity events and print them as readable and formatted and styled text
- I used [cobra](https://github.com/spf13/cobra) for cli and [lip gloss](https://github.com/charmbracelet/lipgloss) for styling the text

## How to run
- clone the project
```shell
git clone https://github.com/HoseinAsadolahi/github-user-activity.git
cd github-user-activity
```
- To run this use commands below
```shell
go build -o activity-event
or
go build -o activity-event.exe
```
then 
```shell
./<file-name> <username> --page <value>
or
./<file-name> <username> -p <value>
or
./<file-name> <username>
```
