# Github / Slack reminder

This project is used to notify a team in Slack of pending pull-requests in GitHub.

## Get started

```bash
# Required
export SLACK-TOKEN=xxxx-xxxxxxxxxx-xxxxxxxxxxx-xxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
export GITHUB-TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

# Optional
export SLACK-CHANNEL=random      # Default: engineering
export SLACK-ICON-EMOJI=:rocket: # Default: :tada:
export SLACK-USERNAME=Bob        # Default: GitHub Reminder

./build/linux/gh-slack-reminder
```
