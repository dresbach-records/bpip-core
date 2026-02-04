# BPIP Economy Specification

## 1. Purpose of This Document

This document defines the **economic model** enforced by the BPIP Core protocol.

It specifies:
- how value enters the system
- how value is measured
- how participation is calculated
- how redistribution occurs
- how inflation is controlled
- how abuse is prevented

This document is **normative**.  
All implementations of BPIP Core **must follow these rules**.

---

## 2. Economic Principles (Non-Negotiable)

BPIP economy is built on the following principles:

1. **No value is distributed unless value was generated**
2. **Participation is earned through verified contribution**
3. **Holding alone does not generate rewards**
4. **Emission is proportional to real usage**
5. **All redistribution rules are explicit and auditable**
6. **The system must remain solvent at all times**

The BPIP economy does not rely on belief, speculation, or future promises.

---

## 3. What Generates Economic Value

Economic value enters the BPIP system only through **real, measurable activity**.

Valid value-generating activities include:

- AI computation (text, image, code, data)
- Build and deployment operations
- Hosting and infrastructure usage
- API consumption
- Enterprise services
- Paid integrations

Each activity:
- has a measurable cost
- produces a signed usage event
- is recorded by the protocol

If no usage occurs, no value is generated.

---

## 4. Economic Unit (Participation Unit)

BPIP uses a **Participation Unit (PU)**.

### Properties
- PU is **not fiat**
- PU is **not a security**
- PU is **not a promise of profit**
- PU represents **accumulated economic participation**

PU exists only as a ledger record within BPIP Core.

---

## 5. Value Inflow

When a service is used:

1. A **usage event** is generated
2. The event has an associated **economic value**
3. That value enters the BPIP economy
4. The protocol allocates it to predefined pools

Value may originate from:
- external payment
- internal credit
- enterprise agreement

BPIP Core does not manage payment rails.  
It only processes the **economic result**.

---

## 6. Economic Pools

All incoming value is split into fixed pools.

### 6.1 Infrastructure Pool
Covers:
- servers
- GPU
- storage
- bandwidth
- operational costs

This pool ensures system sustainability.

---

### 6.2 Operator Pool
Covers:
- platform operation
- protocol maintenance
- development costs
- security

This is how the BPIP organization sustains itself.

---

### 6.3 Contribution Pool
Used exclusively to:
- reward contributors
- reward active developers
- reward maintainers
- reward operators of useful services

Only this pool generates participation units for users.

---

### 6.4 Reserve Pool
Used to:
- stabilize the economy
- absorb demand spikes
- protect against downturns
- fund long-term resilience

The Reserve Pool does not distribute value directly.

---

## 7. Pool Allocation Rules

- Pool percentages are **explicitly configured**
- Percentages must sum to 100%
- Changes require protocol versioning
- No dynamic or hidden reallocation is allowed

Example (illustrative only):

| Pool | Allocation |
|---|---|
| Infrastructure | 40% |
| Operator | 20% |
| Contribution | 30% |
| Reserve | 10% |

---

## 8. Contribution and Participation

### 8.1 What Counts as Contribution

Contribution is **productive impact**, not raw activity.

Valid contributions include:
- accepted code changes
- maintained services
- meaningful usage of infrastructure
- validated uptime
- system improvements

Contribution must be:
- verifiable
- attributable
- resistant to automation abuse

---

### 8.2 Contribution Scoring

Each contribution receives a **score** based on:
- impact
- reuse
- longevity
- dependency weight

High-volume, low-impact actions score poorly.

---

### 8.3 Participation Distribution

Participation Units are distributed only from the Contribution Pool.

Distribution is:
- proportional to contribution score
- bounded by protocol limits
- time-weighted to discourage farming

---

## 9. Emission Rules

Emission is **not fixed** and **not speculative**.

Rules:
- Emission occurs only when value enters the system
- Emission is capped per epoch
- Emission decays if usage declines
- No emission without usage

There is no concept of “mining rewards” detached from activity.

---

## 10. Inflation Control

BPIP economy uses **elastic emission**.

- Higher usage → higher emission
- Lower usage → lower emission
- No artificial scarcity
- No uncontrolled inflation

Participation Units represent **claims on historical value**, not future promises.

---

## 11. Abuse Prevention

The protocol enforces:
- diminishing returns
- contribution decay
- anti-spam validation
- fraud detection

Any detected abuse:
- invalidates events
- revokes participation
- permanently blocks reward generation

---

## 12. Economic Solvency

At all times:
- infrastructure costs must be covered first
- redistribution never exceeds inflow
- reserve maintains minimum thresholds

If costs exceed inflow:
- redistribution is reduced
- emission slows automatically

The system **cannot go bankrupt by design**.

---

## 13. What the Economy Does NOT Do

BPIP economy does not:
- guarantee appreciation
- promise income
- reward passive holding
- subsidize unproductive behavior
- mint value without usage

---

## 14. Final Principle

> **BPIP distributes only what has already been earned by the system.**

There are no shortcuts.
There is no speculation layer.
There is no free extraction.

Only contribution, usage, and impact.
