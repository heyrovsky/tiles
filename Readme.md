# Tiles - Git-based Infrastructure Orchestration

**Tiles** is a GitOps-inspired orchestrator that automates infrastructure provisioning, deployment, and task execution using Git as the source of truth. It should (in future, thats the plan 😄) integrates seamlessly with Ansible and provides powerful branching logic, reconciliation workflows, and distributed node coordination.

## ⚙️ Key Concepts

| Concept      | Description                                                                     |
| ------------ | ------------------------------------------------------------------------------- |
| **Tile**     | A Git branch representing a unit/domain of automation (e.g., `tiles-hyperlane`) |
| **Node**     | A participant that subscribes to specific tiles to execute automation logic     |
| **Metadata** | A centralized registry of versioning, checks, archive policy, and conditions    |

---

## 📦 Workflow

1. **Metadata Check Before Cloning**

   * Nodes first consult the `tiles-orchestrator-metadata` repo to:

     * Validate version compatibility
     * Check if the tile is archived or active
     * Fetch configuration hints

2. **Tile Cloning**

   * If the node passes validation, it proceeds to clone the appropriate `tiles-tile-<NAME>` repo.

3. **Execution and Reconciliation**

   * Nodes pull instructions from their subscribed tile branches.
   * PRs require manual approval (`tiles approve`).
   * On merge, nodes reconcile their state with the latest Git definition.

## 🛡️Benefits
* **Immutable History**: Git as the source of truth.
* **Distributed Control**: Nodes act independently but coordinate via Git.
* **Manual Gates**: Approvals via PR workflow.
* **Declarative Infra**: Use Git branches to define automation per domain.



## Motivation

* **Boredom**: I am bored, so I’m creating the code.
