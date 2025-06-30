# Gator - RSS Feed CLI Tool 🐊

A command-line RSS feed aggregator that helps you stay up to date with your favorite blogs and news sources.

## Prerequisites

* Go 1.21+ 
* PostgreSQL database

## Installation

Install the `gator` CLI using Go:

```bash
go install github.com/rdcassin/gator-go@latest
```

Note: Make sure your $GOPATH/bin is in your system's PATH so you can run gator from anywhere.

Alternatively, you can clone and build locally:

```bash
git clone https://github.com/rdcassin/gator-go.git
cd gator-go
go build -o gator
```

## Configuration

Create a `.gatorconfig.json` file in your **home directory** (e.g., `~/.gatorconfig.json` on Linux/macOS, `C:\Users\YourUser\.gatorconfig.json` on Windows) with the following structure:

```json
{
  "db_url": "postgres://username:password@localhost/gator?sslmode=disable",
  "current_user_name": "your_username"
}
```

* db_url: Your PostgreSQL database connection string. Ensure the database gator exists and the user has permissions.
* current_user_name: The username you wish to register or use as the default.

## Usage

Commands are executed using gator <command>.

### Basic Commands

**`gator register <username>`**
Register a new user. You need to do this once for each user.
Example: `gator register john`

**`gator login <username>`**
Switch to a different user. This changes the active user context for subsequent commands.
Example: `gator login jane`

**`gator addfeed <name> <url>`**
Add a new RSS feed to the system. This feed becomes available for any user to follow.
Example: `gator addfeed "Tech News" https://example.com/rss`

**`gator feeds`**
List all available RSS feeds that have been added to the system.

**`gator follow <feed_name>`**
Follow an existing feed for the currently logged-in user. You will start seeing posts from this feed when you browse.
Example: `gator follow "Tech News"`

**`gator following [-v]`**
List all feeds followed for the currently logged-in user. By default, the command lists just the Feed Names. The '-v' tag (verbose)
makes the command list all details for each feed followed.
Example: `gator following -v`

**`gator browse [limit]`**
Browse recent posts from all feeds currently followed by the active user.
[limit] (optional): The maximum number of posts to display. Defaults to a sensible number if not specified (e.g., 20).
Example: `gator browse 10`
Example (default limit): `gator browse`

## Example Workflow

```bash
# Register a new user
gator register john

# Log in as that user (optional, but good practice if you have multiple users)
gator login john

# Add a new feed to the system
gator addfeed "Tech News" https://example.com/rss

# Follow the newly added feed
gator follow "Tech News"

# Browse the latest 10 posts from all followed feeds
gator browse 10
```

## Troubleshooting

- **Database connection issues**: Make sure PostgreSQL is running and the database URL is correct
- **Permission errors**: Ensure your PostgreSQL user has the necessary permissions
- **Command not found**: If using gator, make sure your $GOPATH/bin is in your system's PATH so you can run `gator` from anywhere