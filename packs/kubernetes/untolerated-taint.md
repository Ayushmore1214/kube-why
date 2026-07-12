---
title: Untolerated taint
aliases: [untolerated-taint, node(s)-had-untolerated-taint, had-untolerated-taint]
category: scheduling
---

# Untolerated taint

## What it means
The scheduler found nodes that could run the pod, but every one of them has a
taint the pod doesn't tolerate. This is different from insufficient CPU or
memory: the cluster may have plenty of capacity, but Kubernetes intentionally
keeps the pod off those nodes until a matching toleration is present.

## Common causes
1. The pod is trying to run on control-plane nodes, which are tainted by
   default, but it doesn't define the required toleration.
2. Dedicated nodes (for databases, infrastructure, or other reserved
   workloads) have custom taints, and the pod doesn't tolerate them.
3. GPU nodes have taints to reserve them for GPU workloads, but the pod isn't
   allowed to run there.
4. A node has been tainted temporarily for maintenance or another operational
   reason, so new pods are intentionally kept off it.

## Check it
```
kubectl describe pod <pod>
kubectl get nodes -o custom-columns=NAME:.metadata.name,TAINTS:.spec.taints
kubectl get pod <pod> -o yaml
```

The Events section usually shows the exact taint that blocked scheduling.
Compare it with the pod's `tolerations`. If no toleration matches that taint,
those nodes aren't eligible for scheduling even if they have enough free CPU
and memory.

## Fix it
- If the workload is supposed to run on tainted nodes, add the matching
  toleration to the pod spec.
- If the taint was added temporarily for maintenance, remove it once the work
  is complete instead of adding tolerations to every workload.
- If the workload shouldn't run on those nodes, review the node selector,
  affinity rules, or scheduling constraints instead of adding a toleration.
- Don't add broad tolerations just to make scheduling succeed. They can allow
  workloads onto nodes that were intentionally reserved.

## Related
- FailedScheduling
- Pending