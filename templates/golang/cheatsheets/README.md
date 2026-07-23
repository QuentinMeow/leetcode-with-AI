# Go cheatsheets

- Numbered `NN_<topic>_<structures>.go` files in this folder are the editable source topics.
- `0_cheatsheet.go` is generated from `manifest.json`; do not edit it directly.
- Regenerate with `npm run cheatsheets:generate`.
- Verify with `npm run cheatsheets:verify`.

Topic files use `package cheatsheets` and compile together; the generated aggregate uses `package main` and is excluded from that test via `//go:build ignore`.
