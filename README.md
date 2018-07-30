# What's this?

Create a bookmark which works as a gmail template

# Usage

Edit sample.json.

```
{
    "to":["test@example.com"],
    "bcc":["test@example.com"],
    "subject": "Subject comes here",
    "body": [
        "Hi, ",
        "",
        "This is a sample mail to show you how to use a template.",
        "",
        "Enjoy!",
        "",
        ""
    ]
}
```

Execute gmailtemplate.go.

```
$ go run gmailtemplate.go sample.json
mailto:test@example.com?BCC=test@example.com&Subject=Subject%20comes%20here&Body=Hi%2C%20%0A%0AThis%20is%20a%20sample%20mail%20to%20show%20you%20how%20to%20use%20a%20template.%0A%0AEnjoy%21%0A%0A&
```

Copy and paste the output in a bookmark's URL field.

