# TaskForge — Distributed Job Scheduling SaaS

## Overview

TaskForge is a developer-focused SaaS platform for scheduling, executing, and monitoring background jobs across distributed worker machines.

Developers can create jobs, schedule them using cron expressions, execute them on demand, monitor their status, and inspect execution logs through a web dashboard.

The core idea is to separate **job orchestration** from **job execution**: TaskForge Cloud manages scheduling and queues, while lightweight workers execute jobs on the user's machines or servers.

## Problem

Developers frequently need to run recurring background tasks such as:

- Database backups
- Data processing
- Web scraping
- Automated scripts
- Report generation
- CI/CD tasks
- Periodic API calls

Building reliable infrastructure for scheduling, retries, queues, worker management, and monitoring is unnecessarily complicated for small projects.

TaskForge aims to provide a simple interface for managing these workloads.

## Core Features

### Job Management
- Create and delete jobs
- Manual job execution
- Cron-based scheduling
- Job priorities
- Execution timeouts
- Retry policies

### Distributed Workers
- Register workers with TaskForge
- Distribute jobs across available workers
- Worker health monitoring
- Concurrent job execution
- Automatic retry when workers fail

### Monitoring
- Real-time job status
- Execution history
- Execution logs
- Failed-job inspection
- Worker status dashboard

### Developer API
TaskForge will expose an API so developers can trigger jobs programmatically.

Example:

```bash
POST /api/jobs/{job_id}/run
```

This allows TaskForge to integrate with existing applications and automation pipelines.

## Architecture

```text
                    TaskForge
                       │
             ┌─────────┴─────────┐
             │                   │
        Web Dashboard         Go API
                                 │
                    ┌────────────┴────────────┐
                    │                         │
                Scheduler                 Job Queue
                                          (Redis)
                                              │
                              ┌───────────────┼───────────────┐
                              ↓               ↓               ↓
                           Worker 1        Worker 2        Worker 3
                              │               │               │
                              └───────────────┴───────────────┘
                                              │
                                           Results
                                              ↓
                                         PostgreSQL
```

## Technology Stack

**Backend**
- Go
- REST API
- WebSockets

**Frontend**
- Next.js
- Tailwind CSS

**Infrastructure**
- Docker
- Redis
- PostgreSQL

**Development**
- GitHub
- GitHub Actions

The entire MVP can be developed and tested locally without paid APIs or cloud infrastructure.

## 12-Week Development Plan

### Week 1 — Foundation
Build the Go backend, project structure, database models, and basic API.

### Week 2 — Job System
Implement job creation, deletion, execution, and job states.

### Week 3 — Scheduler
Add cron-based scheduling and automatic job triggering.

### Week 4 — Queue
Introduce Redis and build the job queue.

### Week 5 — Workers
Create independent workers that register with the server and execute queued jobs.

### Week 6 — Distributed Execution
Implement worker selection, concurrency, retries, and failure handling.

### Week 7 — Logging
Capture stdout/stderr and store execution logs.

### Week 8 — Dashboard
Build the web interface for managing jobs and viewing execution history.

### Week 9 — Real-Time Monitoring
Add WebSockets for live job status, worker health, and logs.

### Week 10 — Developer API
Add API keys, programmatic job execution, and webhooks.

### Week 11 — Docker & Deployment
Containerize the complete system and create a simple deployment workflow.

### Week 12 — Polish & Launch
Improve UI, documentation, testing, security, demo, and publish the project publicly.

## Future Features

If the core system is completed early, TaskForge could eventually support:

- GitHub integration
- Docker-based isolated execution
- Team accounts
- Role-based access control
- Usage-based billing
- Worker autoscaling
- Job dependencies/DAGs
- Metrics and analytics
- Self-hosted workers
- Kubernetes workers

## Final Goal

By the end of Terra, TaskForge should be a working distributed job scheduling platform where a developer can:

```text
Create Job
    ↓
Schedule Job
    ↓
TaskForge queues it
    ↓
Available worker receives it
    ↓
Worker executes it
    ↓
Logs/results returned
    ↓
Dashboard shows the result
```

The project will demonstrate practical knowledge of **Go, distributed systems, networking, databases, queues, concurrency, Docker, APIs, and SaaS architecture** while remaining completely buildable with free/local infrastructure.