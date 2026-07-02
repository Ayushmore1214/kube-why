# kube-why

A command-line reference for Kubernetes errors. You hit `CrashLoopBackOff`,
you run `kube-why crashloopbackoff`, you get what it means, why it usually
happens, and how to check and fix it. No searching, no fifteen open tabs of
blog posts saying the same three things.

```
$ kube-why oomkilled

OOMKilled

What it means
The container tried to use more memory than its limit allowed, so the kernel's
OOM killer terminated it...

Common causes
1. resources.limits.memory is set lower than what the app actually needs...
...
```

It also reads real `kubectl` output directly, no need to already know the
error name:

```
$ kubectl describe pod myapp-7f9c8d6b5 | kube-why

CrashLoopBackOff
...
```

## Why this exists

Every Kubernetes error message is cryptic on purpose, it's a status code, not
an explanation. The explanations exist, but they're scattered across a
thousand near-identical blog posts optimized for search traffic, not for
being useful at 2am when a pod is down. This puts them in one place, in your
terminal, with no ads and no scrolling past a life story to get to the fix.

## Install

macOS/Linux, via Homebrew:

```
brew install ayushmore1214/tap/kube-why
```

Or, with Go 1.21+:

```
go install github.com/Ayushmore1214/kube-why@latest
```

Or build from source:

```
git clone https://github.com/Ayushmore1214/kube-why.git
cd kube-why
go build -o kube-why .
```

Prebuilt binaries for macOS, Linux, and Windows (amd64/arm64) are attached to
every [release](https://github.com/Ayushmore1214/kube-why/releases) if you'd
rather skip all of the above.

## Usage

```
kube-why <error>      # what it means, why, how to fix it
kube-why list          # everything currently covered, by category
kube-why random        # a random one, for when you're bored or teaching
<kubectl cmd> | kube-why   # auto-detect the error from real kubectl output
```

Lookups are forgiving. `kube-why crashloopbackoff`, `kube-why CrashLoopBackOff`,
and `kube-why "crash loop backoff"` all resolve to the same entry.

Piping works with anything that has the error's reason string in it:
`kubectl describe pod`, `kubectl get events`, `kubectl describe deployment`,
even a pasted Slack message from your incident channel. If more than one
known error shows up in the input, it prints all of them.

## What's covered right now

50 errors spanning pods, scheduling, deployments, jobs, HPAs, storage,
networking, ingress, RBAC, webhooks, Helm, and node-level conditions, not
just the handful of pod-level basics most references stop at.

Run `kube-why list` for the current, authoritative, categorized list.

## Contributing

The bar to add an entry is one markdown file. If you've debugged a
Kubernetes error and know the real causes and the real fix, that's worth more
than another generic blog post, write it down here instead. See
[CONTRIBUTING.md](CONTRIBUTING.md) for the format.

If you don't want to write an entry but know one that's wrong or outdated,
open an issue, that's just as useful.

## Roadmap

Rough order, not promises:

- A `kubectl` plugin (`kubectl why <error>`) via krew, so it lives where
  you're already working.
- `kube-why scan` — inspect your current cluster context directly and
  surface which of your pods are unhealthy and why, no manual lookup needed.
- More entries: Terraform and Docker/containerd-level errors, not just
  Kubernetes-native ones.
- Each entry linking to a one-command way to reproduce it in a local `kind`
  cluster, so you can practice recognizing and fixing it before you hit it
  for real.
- A static site version of the same content for people who'd rather search
  than install a CLI.

## License

MIT, see [LICENSE](LICENSE).
