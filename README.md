# Extend PocketBase with Go To Send Text Messages

With a little bit of Go code and the Twilio Go Package we can send text messages with PocketBase.

Don't forget to [install Go](https://go.dev/doc/install) first.

Next, download this project and run this
```bash
go mod init myapp && go mod tidy
go run main.go serve
```

Create a new .env file with these environment variables.
```bash
TWILIO_ACCOUNT_SID="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
TWILIO_AUTH_TOKEN="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
TWILIO_NUMBER="+18888675309"
```

Now, create your PocketBase admin here: http://127.0.0.1:8090/_/ 

See [PocketBase Docs](https://pocketbase.io/docs/go-overview/) for more information on how to setup PocketBase on your machine.

## Texts Collection

Now create a new collection called texts in PocketBase for the text messages. It should have a text field called to, and a text field called message.

Finally, create a new record to send a text message.

## SvelteKit Front End UI

[Also checkout this simple user interface made with SvelteKit.](https://github.com/thedittmer/pbtexts)
