# axlr-repair-lab

A disposable Go repository where [AXLR](https://github.com/underpass-ai/AXLR)'s
`/repair` workflow is tested in public. Each round starts from a commit that
breaks `go test ./...`. The console clones the repository, reproduces the
failure, writes a fix, opens a pull request, waits for CI and merges only when
CI passes.

Expect the history to be rewritten by machines.

## Rounds so far

| PR | Bug planted | Fix AXLR proposed | CI |
|---|---|---|---|
| [#1](https://github.com/underpass-ai/axlr-repair-lab/pull/1) | `WordCount` splits on single spaces, so empty text and repeated spaces miscount | `len(strings.Fields(text))` | pass, merged |
| [#2](https://github.com/underpass-ai/axlr-repair-lab/pull/2) | `Reverse` reverses bytes, breaking `"año"` | reverse runes | pass, merged |
| [#3](https://github.com/underpass-ai/axlr-repair-lab/pull/3) | `Sum` skips the first element | loop from index 0 | pass, merged |
| [#4](https://github.com/underpass-ai/axlr-repair-lab/pull/4) | `Max` ignores negative numbers | start from the first element, plus two negative-number tests | pass, merged |

All four rounds ran on 4 October 2026. They are small, single-file bugs: they
show that the workflow closes the loop from a red test to a merged PR, not that
AXLR repairs hard bugs. Each PR body holds the diagnosis and the plan the
model wrote.

## What CI checks

`gofmt`, `go vet ./...` and `go test ./...` on every push and pull request,
with read-only repository permissions and no secrets.

## Run a round yourself

1. Install AXLR (see [Install](https://axlr.underpassai.com/install/)).
2. Fork this repository and push a commit that makes `go test ./...` fail.
3. In `axlr-tui`, run `/repair` against your fork and approve the cards.

A test that passes only checks what it asserts. Read the diff before you trust
the green check.
