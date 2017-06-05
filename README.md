# Github PRs, Slack reminder

[![CircleCI](https://circleci.com/gh/bureaux-a-partager/gh-slack-reminder/tree/master.svg?style=svg&circle-token=86fdb799067b80015475113b1061810f81bfb33e)](https://circleci.com/gh/bureaux-a-partager/gh-slack-reminder/tree/master)

This project is used to notify a team in Slack of pending pull-requests in GitHub.


## Get started

```bash
# Required
export SLACK_TOKEN=xxxx-xxxxxxxxxx-xxxxxxxxxxx-xxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
export GITHUB_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

# Optional
export SLACK_CHANNEL=random      # Default: engineering
export SLACK_ICON_EMOJI=:rocket: # Default: :tada:
export SLACK_USERNAME=Bob        # Default: GitHub Reminder

export GITHUB_QUERY=is:open is:pr user:bureaux-a-partager label:RFR # Default: is:open is:pr

./build/linux/gh-slack-reminder
```


## License

See [LICENSE](https://github.com/bureaux-a-partager/gh-slack-reminder/blob/master/LICENSE).


## GPG Signatures

You can download Julien Breux's public key to verify the signature.

    gpg --keyserver hkp://pgp.mit.edu --recv-keys 951C3F93B6A8C22C
