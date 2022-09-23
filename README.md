# kasi-bot

A twitter bot to quote lyrics with.
Expects a couple of environment variables to be set.

To tweet from a twitter account:
- `ACCESS_TOKEN` - Twitter access token
- `ACCESS_TOKEN_SECRET` - Twitter access token secret
- `GOTWI_API_KEY` - Twitter API key
- `GOTWI_API_KEY_SECRET` - Twitter API key secret

To clone the git repository the lyrics are found in:
- `GIT_USER_NAME` - Git user account for the bot
- `GIT_ACCESS_TOKEN` - Git access token for the bot
- `LYRICS_REPOSITORY` - Git repository URL where the lyrics are stored

The lyrics repository is expected to contain a directory named `lyrics` containing all lyrics as individual files, with or without subdirectories.
Each lyric file is expected to begin with a line following the format `[Artist]／[Song title]`
