# financial-obligation

### Install 
```
    . set-env.sh rosered

    set-chain-env.sh -n my -v 1.0 -p my

    chain.sh install -p
```
### Instantiate

```
    set-chain-env.sh        -c   '{"Args":["init","ACFT","1000", "A Cloud Fan Token!!!","john"]}'
    chain.sh  instantiate
```
### Query

```
    set-chain-env.sh         -q   '{"Args":["balanceOf","john"]}'
    chain.sh query
```

### Invoke

```
    set-chain-env.sh         -i   '{"Args":["transfer", "john", "sam", "10"]}'
    chain.sh  invoke
```

# Reference

- https://raw.githubusercontent.com/acloudfan/HLF-GO-2.0/master/token/ERC20/README.md