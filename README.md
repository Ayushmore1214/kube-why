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

## Why this exists

Every Kubernetes error message is cryptic on purpose, it's a status code, not
an explanation. The explanations exist, but they're scattered across a
thousand near-identical blog posts optimized for search traffic, not for
being useful at 2am when a pod is down. This puts them in one place, in your
terminal, with no ads and no scrolling past a life story to get to the fix.

## Install

Requires Go 1.21+.

```
go install github.com/Ayushmore1214/kube-why@latest
```

Or build from source:

```
git clone https://github.com/Ayushmore1214/kube-why.git
cd kube-why
go build -o kube-why .
```

## Usage

```
kube-why <error>      # what it means, why, how to fix it
kube-why list         # everything currently covered, by category
kube-why random       # a random one, for when you're bored or teaching
```

Lookups are forgiving. `kube-why crashloopbackoff`, `kube-why CrashLoopBackOff`,
and `kube-why "crash loop backoff"` all resolve to the same entry.

## What's covered right now

Fourteen of the errors you'll actually hit running real workloads:
`CrashLoopBackOff`, `ImagePullBackOff`, `OOMKilled`, `Pending`, `Evicted`,
`CreateContainerConfigError`, `InvalidImageName`, stuck `ContainerCreating`,
stuck `Terminating`, `ProgressDeadlineExceeded`, `FailedScheduling`,
`ReadinessProbeFailed`, `NodeNotReady`, `FailedMount`.

Run `kube-why list` for the current, authoritative list.

## Contributing

The bar to add an entry is one markdown file. If you've debugged a
Kubernetes error and know the real causes and the real fix, that's worth more
than another generic blog post, write it down here instead. See
[CONTRIBUTING.md](CONTRIBUTING.md) for the format.

If you don't want to write an entry but know one that's wrong or outdated,
open an issue, that's just as useful.

## Roadmap

Rough order, not promises:

- More entries: Helm, Terraform, and Docker errors, not just raw Kubernetes.
- A `kubectl` plugin (`kubectl why <error>`) via krew, so it lives where
  you're already working.
- Each entry linking to a one-command way to reproduce it in a local `kind`
  cluster, so you can practice recognizing and fixing it before you hit it
  for real.
- A static site version of the same content for people who'd rather search
  than install a CLI.

## License

MIT, see [LICENSE](LICENSE).
