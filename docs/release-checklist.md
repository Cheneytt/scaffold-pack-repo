# HOWTO: Release a new version of scaffold-pack-repo

1. Check out the source
1. Tag a release from master
1. Push the tag
1. Check out the tag
1. Run `goreleaser release` to compile and build the release assets
1. Upload the assets to the Github Releases page for scaffold-pack-repo
1. Check out master again
1. Bump the version number in plugin.yaml
1. Commit and push to master
