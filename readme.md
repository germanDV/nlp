# NLP

Natural Language Password generator based on **diceware**.


## Usage

By default, it generates a 4-word password.
This can be overridden by simply passing the desired number of words as an argument.

```sh
# generate a 4-word password
$ go run main.go

# generate a 6-word password
$ go run main.go 6
```

## Tests

Run all tests:

```sh
$ go test ./...
```

## Wordlists

The idea behind having 2 wordlists is that, eventually, the _4-dice.csv_ list will only contain adjectives and the _5-dice.csv_ nouns; so that two words form an adjective-noun pair, making it easier to remember passwords.

For instance, a 4-word password could look something like:
 + **fast car blue moon**
 
And a 6-word password could be:
 + **ugly day important pizza great idea**

