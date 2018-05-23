# La7lo7, the first true egyption cryptocurrency 😂

## IMPORTANT CLARIFICATIONS

1. I am doing this project for fun
2. Don't relay on my implementation for any means 😅

## Roadmap

- [x] Basic structure of blocks and transactions.
- [x] Basic proof of work.
- [x] Storing the transactions in the disk.
- [x] Some sort of cli to ease the mining process.
- [ ] Advanved Transactions and user's adresses.
- [ ] Implement a true network.

### Build the app

```
$ go build .
```

### CLI

The cli here is very simple and has only 2 commands

1. Add a new block
    ``` 
    $ ./la7lo7 addblock -data "Send one la7lo7 to 7oss"
    ```

2. Print the chain
    ```
    $ ./la7lo7 printchain
    ```

### Persestance of the data

I have used a simple key value database implemented with golang [bolt](https://github.com/boltdb/bolt)