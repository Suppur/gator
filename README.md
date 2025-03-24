# Overview

Gator is a CLI feed aggregator built using Go, Postgresql, SQLC and Goose.

## Installing
To run gator, you'll need to install Go version 1.24.1+ and Postgres.

### Gator CLI
To install the gator CLI, use the `go install 'github.com/Suppur/gator'` command in your terminal.

### Setting up the config file
Gator uses a JSON file to keep track of the databse url and the current user (the logged in user).
You'll need to create a .gatorconfig.json file in your home directory (Linux).  
Here's the structure:
```json
 {
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```
_Note_: Make sure to add `?sslmode=disable` at the end of your db_url string.  
Example: `postgres://<db_username>:<db_password>@localhost:5432/gator?sslmode=disable` 
## Features
You can use the following commands in the CLI;
- *login*   
    Login as a user  
    usage: `gator login <user>`
- *register*  
    Register a user  
    usage: `gator register <user>`
- *reset*  
    Wipe the database (be careful!)  
    usage: `gator reset`
- *users*  
    List all the users from the database  
    usage: `gator users`
- *agg*  
    Aggregate all the posts of the feeds followed by the user following the provided interval and insert them in the database  
    usage: `gator agg <delay>`
- *addfeed*  
    Add a feed to the database  
    usage: `gator addfeed <name> <url>`
- *feeds*  
    Lists all the feeds added to the database  
    usage: `gator feeds`
- *follow*  
    Follow a given feed for the current user  
    usage: `gator follow <url>`
- *following*  
    List all the feeds followed by the current user  
    usage: `gator following`
- *unfollow*  
    Unfollow a given feed for the current user  
    usage: `gator unfollow <url>`
- *browse*  
    Browse all the posts of all the feeds. Provide a number to limit the posts displayed, default is 2.  
    usage: `gator browse <opt:limit>`