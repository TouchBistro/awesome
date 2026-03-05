# Contributing to `awesome`

## **Extending & Feature Requests:**

There are two possible areas where functionality can be added.

### 1. **Credentials Provider**
Creating a new `CredsProvider` implementation to address a specific configuration for a
supported config provider or a missing credentials type supported by the Go AWS SDK v2.

Also, functional options for a **provider's** `NewXXXProvider()` function may be added 
in future to cover missing or new configuration options for that provider.


### 2. Clients Function
Maintenance of AWS client builders implmented in the `client.go` files under each client
package. This is done using the `codegen` auto-generator. More details on it [here](./codegen/README.md)

**Codegen must be re-run and the regenerated `clients/` output committed alongside your change in the following cases:**

- **The codegen template changes** — any edit to the code template in `codegen/codegen.go` must be followed by a regeneration so the committed `client.go` files reflect the new template.
- **A new AWS service needs to be added** — if a service is missing from `clients/` or has been added to the AWS SDK since the last generation, run codegen to pick it up.
- **An AWS service package is renamed or removed** — the codegen fetches the current service list from the AWS SDK; re-running it will reconcile any renames or removals.

To regenerate:
```bash
cd codegen && go generate codegen.go
```

CI does not regenerate client code automatically. It trusts what is committed and validates it via `make build`, `make lint`, and `make test`. It is the contributor's responsibility to run codegen locally and include the updated `clients/` files in the same PR as the change that triggered the regeneration.

Features or extensions in any other areas, especially the ones that are not backward compatible
will not be considered. 

## **Bugs:**

You can file a bug report with relevant code snippets and reproduction instructions & send
it to the maintainers at the contact information at the end of this document. 


## **Contribution Guidelines:**
<br>

> *Pull Requests are only accepted by members of the TouchBistro Github organization*

<br>

## **How to**:

When making a PR, make sure:
 
1. You're working from the latest source on the `master` branch.
2. You've checked the existing open, and recently closed, pull requests to be sure
   that someone else hasn't already addressed the problem.
3. You create an issue before working on the contribution with all the details
   pertaining to the recommended changes or updates you're contribution will 
   address.


## Maintainers:

You can contact, send bug reports or features requests to the `awesome` maintainers 
by sending an email at: devops@touchbistro.com 

Members of the *TouchBistro* Github org can alternatively create a ticket in the `DEVOPS` Jira project 
& tag a *DevOps* team member for action in the internal communications channels.