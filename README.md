# ucc_twitter_bot
Bot for monitoring [Uncommon Core](https://uncommoncore.co/podcast/) podcast releases and tweeting them out. The program compares the episodes that have already been tweeted with the episodes present in Spotify and tweets out the delta (if any).


## Install
```go get github.com/davidiola/ucc_twitter_bot```

## Usage
Set appropriate API Key environment variables (see [constants](https://github.com/davidiola/ucc_twitter_bot/blob/master/constants/constants.go)).

```go run main.go```

## Twitter Account
Tweets are being submitted by the ```@ucc__bot``` account: https://twitter.com/ucc__bot. Turn on tweet notifications for this account to be notified of new episodes minutes after release.

