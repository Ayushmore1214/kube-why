# Contributing

The whole point of this project is that adding an entry should be easy. One
file, one PR.

## Adding a new error

1. Copy the template below into `errors/<slug>.md`. The slug is the filename,
   lowercase, words separated by hyphens (`invalid-image-name.md`, not
   `InvalidImageName.md`).
2. Fill it in based on something you've actually debugged, not a summary of
   someone else's blog post. If you had to figure it out the hard way, that's
   exactly the knowledge this project wants.
3. Run it locally to make sure it parses and reads well in a terminal:
   ```
   go run . <your-slug-or-alias>
   ```
4. Open a PR with just that one file (plus README.md's covered-errors list if
   you want to keep it in sync, not required).

### Template

```markdown
---
title: YourErrorName
aliases: [lowercase, alternate, spellings]
category: pod
---

# YourErrorName

## What it means
One or two sentences. What is Kubernetes actually telling you here, in plain
language, not a restatement of the error string.

## Common causes
Ranked roughly by how often you'll actually see them, not an exhaustive list
of every theoretical cause. Three to five is usually right.

## Check it
The exact kubectl commands to confirm which cause you're dealing with. Real
commands, not "check the logs." Explain what to look for in the output, not
just what command to run.

## Fix it
The actual fix for each common cause, matched to the causes listed above. If
there's a quick mitigation and a real fix, say which is which.

## Related
Two or three other entries someone debugging this might also need.
```

`category` should be one of: `pod`, `deployment`, `scheduling`, `node`,
`networking`, `ingress`, `job`, `hpa`, `namespace`, `webhook`, `helm`,
`storage`, `statefulset`, `quota`, `podsecurity`, `apiserver`,
`container-runtime`, `rbac`, or a new one if none of those fit (say why in
the PR).

## What makes a good entry

- Causes are ranked by likelihood, not alphabetized or randomly ordered.
- The "check it" commands are things you'd actually run, and say what the
  output tells you, not just the command itself.
- No filler. If a sentence doesn't help someone fix the problem faster,
  cut it.
- Written like you're telling a coworker, not writing documentation.

## Fixing an existing entry

Same process, no new file needed. If something's outdated, wrong, or you have
a better fix, edit the file directly and explain what changed and why in the
PR description.

## Reporting an error you don't have time to write up

Open an issue with the error string and, if you know it, what actually caused
it for you. Someone else can turn it into a full entry.
