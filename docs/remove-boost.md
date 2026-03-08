# Remove boost container to get latest

Sometimes you will want to remove a previously created `boost` container and restart by rebuilding it (in the section above). Here's how.

> [!WARNING]
> You will lose all your work done in the `boost` container when you remove it unless you have saved it to GitHub or some other way. Use extreme caution.

```sh
podman rm boost
```

You are now able to create a fresh new `boost` container using the latest from this repo.

