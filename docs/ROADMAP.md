# BPIP Roadmap

## Purpose of This Roadmap

This roadmap defines the **technical and economic evolution** of the BPIP ecosystem.

It exists to:
- guide contributors and AI systems
- communicate priorities clearly
- prevent scope drift
- align releases with protocol maturity

This roadmap is **descriptive**, not promotional.
Timelines may adapt, principles do not.

---

## Roadmap Principles

All roadmap phases follow these rules:

- Protocol correctness before features
- Economics before UX
- Auditability before speed
- No speculative shortcuts
- No dependency inversion (core before apps)

If a feature violates these principles, it is out of scope.

---

## Phase 0 — Foundation (Completed)

**Status:** Completed  
**Repositories:** bpip-core  

### Goals
- Define the economic protocol
- Lock core principles
- Establish governance and security boundaries

### Delivered
- Economic model (`ECONOMY.md`)
- Protocol rules (`PROTOCOL.md`)
- Contribution model (`CONTRIBUTION.md`)
- Governance structure (`GOVERNANCE.md`)
- Security model (`SECURITY.md`)
- Core documentation and architecture index
- CHANGELOG and release process

---

## Phase 1 — Core Implementation (Current)

**Status:** In progress  
**Repositories:** bpip-core  

### Goals
- Implement protocol logic in Go
- Ensure deterministic behavior
- Establish test coverage

### Scope
- Event system
- Validation layer
- Participation ledger
- Economic engine
- Contribution scoring
- Abuse prevention

### Non-Goals
- No UI
- No payments
- No mobile apps
- No AI orchestration

---

## Phase 2 — Application Backend

**Status:** Planned  
**Repositories:** bpip-backend  

### Goals
- Provide API access to BPIP
- Emit protocol events
- Manage authentication and projects

### Scope
- Auth & identity
- Project lifecycle
- Event emission to bpip-core
- Usage metering

### Non-Goals
- No economic calculation
- No ledger mutation

---

## Phase 3 — Developer Platform (Web)

**Status:** Planned  
**Repositories:** bpip-dashboard  

### Goals
- Enable developers to interact with BPIP
- Provide transparency into participation
- Manage projects visually

### Scope
- Dashboard UI
- Participation views (read-only)
- Usage analytics
- Project management

---

## Phase 4 — AI & Infrastructure Layer

**Status:** Planned  
**Repositories:** bpip-ai, bpip-infra  

### Goals
- Integrate AI usage into BPIP economy
- Provide scalable infrastructure
- Measure real cost and impact

### Scope
- AI request routing
- Usage events
- GPU and compute tracking

---

## Phase 5 — Mobile Applications

**Status:** Planned  
**Repositories:** bpip-mobile-android, bpip-mobile-ios  

### Goals
- Lightweight access to BPIP
- Participation monitoring
- Notifications

### Scope
- Read-only participation views
- Alerts
- Project status

---

## Phase 6 — Ecosystem Expansion

**Status:** Future  

### Goals
- Third-party integrations
- External services
- Partner infrastructure

### Scope
- Integration guidelines
- Event contracts
- Ecosystem tooling

---

## Versioning Strategy

- Pre-1.0 releases may introduce breaking changes
- Protocol rules are versioned explicitly
- Breaking changes are documented in CHANGELOG
- Roadmap phases do not imply release dates

---

## How to Use This Roadmap

For contributors:
- Choose tasks aligned with the current phase
- Do not implement features from future phases

For AI systems:
- Respect phase boundaries
- Do not generate out-of-scope components
- Read `ARCHITECTURE_INDEX.md` first

---

## Final Note

BPIP is built incrementally, with discipline.

The goal is not speed.
The goal is correctness, fairness, and longevity.