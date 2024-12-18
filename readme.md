# General Notes

# https://www.udemy.com/course/github-actions-the-complete-guide

GH Actions is a service for automating repo related processes:
1. CI/CD -> lint/build/test/deploy automatically
2. Repo management functions such as automatic code reviews and issue management


Building blocks:
1. Workflows
- Attached to a GH repo
- Contains one or more jobs
- Triggered by events such as human interaction or a push to a branch

2. Jobs
- Defines a runner (the execution environment). Can use default ones or define your own
- Contains one or more steps
- Run in parallel by default (can change to sequential)
- Can be conditional

3. Steps
- Executes a shell script or Action (custom or third party)
- Executed in order
- Can be conditional

Pricing is free for public repos, and you get a quota for private repos
Using large runners always costs money and using certain types of runners such as Mac and Windows uses the quota faster
https://docs.github.com/en/billing/managing-billing-for-your-products/managing-billing-for-github-actions/about-billing-for-github-actions

Path in repo: .github/workflow/file.yml


Runner types:
https://docs.github.com/en/actions/using-github-hosted-runners/using-github-hosted-runners/about-github-hosted-runners
Runners are VMs
Linux/Windows/Mac VMs can be used for free
Large runners are ones with more cores and cost money
Arm and GPU based runners are classed as large runners
Default runners: ubuntu-latest, windows-latest, macos-latest
You can run the jobs directly in the VM or in a Docker container within the VM
All the STEPS are run in the same VM so you can use the filesystem to share information between them
Jobs can run on different runners/VMs
Within a runner the GitHub actions runner (C# based) is installed
VMs are hosted in Azure
VMs for Linux/macOS use password-less sudo for when you need to elevate privileges such as installing packages
A list of installed packages for the VM can be viewed as a link in the "Set up job" built in step


Environment variables. File paths on VMs are not static so always use env vars to construct paths:
- HOME - user related data such as credentials from a login attempt
- GITHUB_WORKSPACE - actions and shells execute in this directory. Shared space between other actions
- GITHUB_EVENT_PATH - the POST payload of a webhook event that triggered this workflow
- GITHUB_SHA - commit ID which triggered the workflow
- GITHUB_REF - reference (branch, tag etc) which triggered the workflow e.g. "refs/heads/main"

If running a Docker container in the VM a directory prefix of /github is used. But it's recommended to still use env vars to construct paths


Workflow triggers:
https://docs.github.com/en/actions/writing-workflows/choosing-when-your-workflow-runs/events-that-trigger-workflows
- push - pushing a commit
- pull_request - pull request event (open, close etc)
- create  - a branch or tag was created etc
- workflow_dispatch - manual trigger
- repository_dispatch - manual trigger via an API call
- schedule - runs on a cron
- workflow_call - a workflow is called by another workflow


Actions:
Re-usable modules which can be official or community created, or private ones.
Marketplace which lists all the public ones: https://github.com/marketplace?type=actions
You can typically link through to the marketplace if on the GitHub source code repo for the Action (there is a displayed button at the top)

When a workflow runs it runs against the target branch for whatever triggered it

Workflows do not automatically run for PRs which are opened from forked repos if it's a first time contributor, even if they should.
They require manual approval to run from a maintainer. After being approved once they are able to run subsequent workloads without approval.
Contributors which have been assigned via the contributors list for the repo are   

Workflow triggers:
There are two ways to filter workflow triggers: "activity types" and "filters"
"Activity types": applied against types like "pull_request" to scope it to certain event types such as "opened", "closed"
"Filters": allow you to filter branch, tag and path names when using events like "push". Can also use negation

When using wildcards "*" matches multiple characters, but not slashes. "**" also include slashes

Workflows - by default - fail if at least one job fails
Jobs - by default - fail if at least one-step fails
You can override this behaviour
You can manually cancel workflows

You can force a workflow to be skipped by adding a particular commit message such as:
"my commit msg [skip ci]"
https://docs.github.com/en/actions/managing-workflow-runs-and-deployments/managing-workflow-runs/skipping-workflow-runs

To pass details between jobs you can:
1. Use simple values (text/numbers etc.). These are defined as outputs on the job level and then steps push key/value pairs to a $GITHUB_OUTPUT envar
2. Artifacts can be created/retrieved using upload-artifact and download-artifact actions. These are files/directories

Caches can be used to reduce build times for config such as build dependencies which change infrequently. Caches can be shared
between multiple jobs and workflows. There are actions for `actions/cache` as well as more granular actions `actions/cache/restore`
and `actions/cache/save` (similar to cache but omit the save and restore built-in steps respectively). An alternative is to use the higher
level actions such as `actions/setup-node` and `actions/setup-go`, which can also handle this caching action.