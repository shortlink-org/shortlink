Release CLI release `vX.Y.Z`

- [ ] Set the milestone on this issue
- [ ] Review the list of changes since the last release and fill below:
    - [ ] **In the changelog**
    - [ ] **Not in the changelog**
- Decide on the version number by reference to
    the [Versioning](https://gitlab.com/gitlab-org/release-cli/blob/master/PROCESS.md#versioning)
    * Typically if you want to release code from current `master` branch you will update `MINOR` version, e.g. `1.12.0` -> `1.13.0`. In that case you **don't** need to create stable branch
    * If you want to backport some bug fix or security fix you will need to update stable branch `X-Y-stable`
- [ ] Create an MR for [release-cli project](https://gitlab.com/gitlab-org/release-cli).
    You can use [this MR](https://gitlab.com/gitlab-org/release-cli/-/merge_requests/20) as an example.
    - [ ] Update `VERSION`
    - [ ] Update `CHANGELOG`. You can use `make generate_changelog`
    - [ ] Assign to reviewer
- [ ] Once `release-cli` is merged create a signed+annotated tag pointing to the **merge commit** on the **stable branch**
    In case of `master` branch:
    ```shell
    git fetch origin master
    git fetch dev master
    git tag -a -s -m "Release v1.0.0" v1.0.0 origin/master
    ```
    In case of `stable` branch:
    ```shell
    git fetch origin 1-0-stable
    git fetch dev 1-0-stable
    git tag -a -s -m "Release v1.0.0" v1.0.0 origin/1-0-stable
    ```
- [ ] Verify that you created tag properly:
    ```shell
    git show v1.0.0
    ```
    it should include something like:
    * ```(tag: v1.0.0, origin/master, dev/master, master)``` for `master`
    * ```(tag: v1.0.1, origin/1-0-stable, dev/1-0-stable, 1-0-stable)``` for `stable` branch
- [ ] Push this tag to origin(**Skip this for security release!**)
    ```shell
    git push origin v1.0.0
    ```

### In the changelog

- ...
- ...

### Not in the changelog

- ...
- ...
