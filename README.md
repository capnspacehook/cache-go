# cache-go

Github Action to restore and save cached Go module and build files as efficiently as possible.

Module and build files are cached seperately to avoid not restoring cached build files because
dependencies were changed. Likewise a Go version upgrade should not invalidate cached module
files as they aren't dependent on a specific Go version. Cache restore keys are used to ensure
a cache will be loaded whenever possible.

## Inputs

```yaml
mod-key:
  description: "An explicit key for restoring and saving the module cache"
  required: false
  default: ${{ github.job }}-${{ hashFiles('**/go.sum') }}
mod-restore-key:
  description: "A prefix of `mod-key` to use to restore a stale cache if no cache hit occurred for `mod-key`"
  required: false
  default: ${{ github.job }}-
build-key:
  description: "An explicit key for restoring and saving the build cache"
  required: false
  default: ${{ github.job }}-${{ runner.os }}-${GO_VERSION}-${{ github.run_id }}-${{ github.run_attempt }}
build-restore-key:
  description: "A prefix of `build-key` to use to restore a stale cache if no cache hit occurred for `build-key`"
  required: false
  default: ${{ github.job }}-${{ runner.os }}-${GO_VERSION}-
```
