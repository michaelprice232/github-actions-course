# github-actions-course

This repo hosts my notes and implementations from the [github-actions-the-complete-guide](https://www.udemy.com/course/github-actions-the-complete-guide) Udemy course.
The sample Node.js app was provided by the course author. I implemented the workflows under `.github/` and added some notes in `my-notes.md`.
Under `extra/` there are some more workflows which didn't fit nicely with the sample code. 


```text
% tree -n .github extra my-notes.md readme.md
.github
├── actions
│ ├── cached-deps
│ │ └── action.yml
│ └── deploy-s3-docker
│     ├── Dockerfile
│     ├── action.yml
│     ├── go.mod
│     ├── go.sum
│     └── main.go
└── workflows
    ├── actions-composite.yml
    ├── actions-docker.yml
    ├── caches-explicit.yml
    ├── caches-implicit.yml
    ├── conditionals.yml
    ├── contextual-info.yml
    ├── issue-events.yml
    ├── matrix.yml
    ├── needs.yml
    ├── outputs.yml
    ├── reusable-workflow-consumer.yml
    ├── reusable-workflow.yml
    ├── triggers-2.yml
    ├── triggers-from-paths.yml
    └── triggers.yml
extra
├── containers
│ ├── code
│ │ ├── go.mod
│ │ ├── go.sum
│ │ ├── main.go
│ │ └── main_test.go
│ └── containers.yml
└── environments
    └── environment-demo.yml
my-notes.md
readme.md

9 directories, 29 files
```

