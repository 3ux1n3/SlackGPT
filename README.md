# SlackGPT

SlackGPT is a simple Slack bot powered by OpenAI API and implemented in Go. It can be used to generate text responses to messages sent in a Slack channel.

## Features

SlackGPT has the following features:

- Text generation: SlackGPT uses the GPT model to generate responses to messages sent in a Slack channel.
- Customizable responses: You can configure the bot to generate responses based on specific keywords or patterns.
- Easy to deploy: SlackGPT is implemented in Go and can be easily deployed to a cloud platform like Heroku or AWS.

## Installation

To install SlackGPT, follow these steps:

1. Clone the repository: `git clone https://github.com/3ux1n3/SlackGPT.git`
2. Create a Slack bot and get an API token. You can follow the instructions [here](https://www.twilio.com/blog/how-to-build-a-slackbot-in-socket-mode-with-go) to create a bot and obtain an API token. 
3. Use the `manifest.yml` file to get all correct settings easily
4. create a app level token for the socket write 
5. Set the `SLACK_AUTH_TOKEN` and `SLACK_APP_TOKEN` environment variables to the app and bot tokens respectively.
6. Set the `OPENAI_API_KEY` from your openai account
7. Run the bot: `go run main.go`


## Usage

Once SlackGPT is up and running, you can send messages to your Slack channel and the bot will generate responses based on the incoming message. the bot will only generate a response if the incoming message starts with `!bot`.

## Contributing

If you'd like to contribute to SlackGPT, please fork the repository and make changes as you'd like. Pull requests are warmly welcome.

## License

SlackGPT is licensed under the [MIT license](https://github.com/3ux1n3/SlackGPT/blob/main/mit.md).
