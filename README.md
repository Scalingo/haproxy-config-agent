# HAProxy Configuration Agent v1.0

The haproxy-config-agent executable is embedded in our databases base image and launched at the start of the HAProxy instance. It aims at handling the hot reload of the HAProxy configuration. It is called when clicking the button "Reload HAProxy config" on the DB admin.

### Release a New Version

Bump new version number in:

- `CHANGELOG.md`
- `README.md`

Commit, tag and create a new release:

```sh
git add CHANGELOG.md README.md
git commit -m "Bump v1.0"
git tag v1.0
git push origin master
git push --tags
hub release create v1.0
```

The title of the release should be the version number and the text of the release is the same as the changelog.

Then you need to update the version specified in the [Dockerfile](https://github.com/Scalingo/appsdeck-database/blob/master/docker/haproxy/1.7/Dockerfile) of the HAProxy image on [appsdeck-database](https://github.com/Scalingo/appsdeck-database).

The process to precisely define the HAProxy version started for each database version has not yet been defined.
