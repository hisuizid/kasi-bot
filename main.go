package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/michimani/gotwi/tweets"
	"github.com/michimani/gotwi/tweets/types"

	"matina-bot/git"
	"matina-bot/tweet"
	"matina-bot/twitter"
)

const (
	path      = "collection"
	directory = "lyrics"
)

func main() {
	// Get the twitter client
	client, err := twitter.GetClient(twitter.GetCredentials())
	if err != nil {
		log.Fatal(err)
	}

	// Get the lyrics from remote repository
	err = git.GetRepo(path, git.GetCredentials())
	if err != nil {
		log.Fatal(err)
	}
	err = validateRepo()
	if err != nil {
		log.Fatal(err)
	}

	// Get the tweet
	tweet, err := tweet.GetTweet(path + "/" + directory)
	if err != nil {
		log.Fatal(err)
	}
	p := &types.ManageTweetsPostParams{
		Text: &tweet,
	}

	// Send the Tweet
	_, err = tweets.ManageTweetsPost(context.Background(), client, p)
	if err != nil {
		log.Fatal(err)
	}
}

// Test quote 100 times and see if each quote follows the following format
//
// LYRIC QUOTE
// -
// ARTIST ／ SONG TITLE
//
func validateRepo() error {
	for i := 1; i <= 100; i++ {
		tweet, err := tweet.GetTweet(path + "/" + directory)
		if err != nil {
			return err
		}
		pattern, _ := regexp.Compile("(.+)\n-\n(.+)／(.+)")
		if !pattern.MatchString(tweet) {
			return fmt.Errorf("repo does not follow valid format. Found invalid quote: %s", tweet)
		}
	}
	return nil
}
