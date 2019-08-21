# Running
No binaries are provided so in order to run this application you will have to build it first. To build this application simply run `go build` in the `./src` directory.

## Configuration
This application requires a config file to run, by default a config.json is provided in the source. If you would like to you can create your own configuration file or update the one in src accordingly. Here is a general overview of the expected config

```json
{
    "firstnames": [
        { "name": "possible first name", "gender": "gender associated with that name (ex: f/m)"}
    ],
    "lastnames": [
        "possible last name"
    ],
    "occupation": [
        "possible occupation"
    ],
    "adjectives": [
        "list of adjectives to associate with occupation"
    ]
}
```

## Starting the application
`./src config.json`