# Gator

gator is an CLI app, designed to aggregate RSS feeds, from the websites of your liking.

## Installation

To install gator, you must need `go` and `postgres` preinstalled. Afterwards, you can do:

```bash
go install
```

Next, you'll have to setup your config in `~/.gatorconfig.json`. Here's an example of a config:

```json
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable","current_user_name":"kahya"}
```

## Usage

Here are a few commands that can be used:

```bash
$ go run .

Usage:

  login          logs a user in
  register       registers a user
  reset          resets database
  users          lists all users
  agg            aggregates feed
  addfeed        adds a feed to be aggregated
  feeds          list all feeds
  follow         follows a feed
  following      lists all feeds that has been followed
  unfollow       unfollows a feed
  browse         browses feed that has been aggregated

```
