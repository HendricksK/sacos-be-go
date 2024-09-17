# TGFDO (thank god for digital ocean)
# https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04

# WATCHER 
# https://pkg.go.dev/github.com/radovskyb/watcher
# https://github.com/radovskyb/watcher


## Go 
## go mod tidy 

## submodules
git submodule update --init --recursive


## New local setup ... need to clean this all up since even I cannot remember what is required

## https://github.com/canthefason/go-watcher

## go mod tidy 

## submodules - no real need for submodules at the moment, but might just do it for db backups, but I think we should build out seeders here.
git submodule update --init --recursive

### using an updated version of fresh which is very helpful since watcher is now giving errors and I am a bit too lazy to figure out why https://github.com/zzwx/fresh
### fresh -g to generate new yml config file -> fresh to start 