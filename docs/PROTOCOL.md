# BPIP Protocol Specification

## 1. Scope of This Document

This document defines the **BPIP Core Protocol**.

It specifies:
- what the protocol is
- what entities exist
- how events are generated
- how events are validated
- how state changes occur
- what the protocol explicitly does not do

This document is **normative**.  
Any BPIP Core implementation **must comply** with these rules.

---

## 2. What the BPIP Protocol Is

The BPIP Protocol is an **economic participation protocol**.

It defines deterministic rules that:
- observe real digital activity
- validate productive contribution
- transform activity into economic participation
- enforce redistribution rules

The protocol is **not autonomous ideology**.  
It is an **explicitly designed system** with clear constraints.

---

## 3. What the BPIP Protocol Is NOT

The protocol does NOT:
- define a blockchain
- perform decentralized consensus
- guarantee neutrality
- mint speculative currency
- enforce financial settlement
- replace sovereign money
- promise economic returns

The protocol assumes:
- an operator exists
- infrastructure exists
- economic coordination exists

This is a **deliberate design choice**.

---

## 4. Core Protocol Entities

### 4.1 Account

An Account represents an identifiable economic participant.

Properties:
- unique identifier
- cryptographic key pair
- participation history
- accumulated participation units

Accounts may represent:
- individual developers
- teams
- services
- infrastructure operators

---

### 4.2 Event

An Event is the **atomic unit of economic observation**.

An event represents **something that actually happened**.

All economic state changes originate from events.

---

### 4.3 Ledger

The Ledger is the **authoritative record** of:
- events
- participation accumulation
- historical distributions

The ledger is append-only.

---

### 4.4 Protocol Engine

The Protocol Engine:
- validates events
- applies economic rules
- updates the ledger
- enforces limits and caps

---

## 5. Event Model

### 5.1 Event Properties

Every event MUST contain:

- unique event ID
- event type
- timestamp
- originating account
- signed payload
- economic metadata

Events are immutable once accepted.

---

### 5.2 Event Types

The protocol recognizes the following event categories:

#### 5.2.1 Usage Events
Represent consumption of system resources.

Examples:
- AI inference
- build execution
- deployment
- API call

---

#### 5.2.2 Contribution Events
Represent productive contribution.

Examples:
- accepted code changes
- maintained services
- validated improvements

---

#### 5.2.3 Infrastructure Events
Represent operational availability.

Examples:
- service uptime
- node operation
- infrastructure maintenance

---

## 6. Event Lifecycle

1. Event is created by a trusted producer
2. Event is signed by the originating account
3. Event is submitted to BPIP Core
4. Event is validated
5. Event is either:
   - accepted and recorded
   - rejected permanently

Rejected events **never** affect state.

---

## 7. Validation Rules

Every event must pass all validation layers:

### 7.1 Structural Validation
- schema correctness
- required fields present
- valid timestamps

---

### 7.2 Signature Validation
- signature authenticity
- key ownership
- replay protection

---

### 7.3 Economic Validation
- allowed event type
- valid economic metadata
- contribution eligibility
- rate limits

---

### 7.4 Anti-Abuse Validation
- spam detection
- artificial inflation detection
- circular contribution detection

Failure at any stage results in rejection.

---

## 8. State Transitions

State changes occur only through:

- accepted events
- deterministic protocol rules

There are:
- no manual overrides
- no administrative balance changes
- no hidden state transitions

---

## 9. Emission and Distribution Flow

1. Value-generating event is accepted
2. Economic value is computed
3. Value is allocated to pools
4. Contribution pool triggers participation calculation
5. Ledger records updated participation units

If no value enters the system, no distribution occurs.

---

## 10. Determinism and Auditability

For a given set of events:
- protocol outcome must be deterministic
- ledger state must be reproducible
- economic distribution must be auditable

Non-deterministic behavior is forbidden.

---

## 11. Versioning and Evolution

- Protocol rules are versioned
- Changes require explicit version bumps
- Historical events remain governed by their original version

Backward compatibility is preferred but not guaranteed.

---

## 12. Governance Boundary

The protocol enforces rules.
Governance defines rules.

BPIP Core:
- enforces
- validates
- records

It does not:
- negotiate
- interpret intent
- arbitrate disputes manually

---

## 13. Failure Modes

If:
- infrastructure degrades
- usage declines
- abuse increases

The protocol:
- reduces emission
- tightens validation
- preserves solvency

The protocol never emits value to compensate failure.

---

## 14. Security Assumptions

The protocol assumes:
- cryptographic primitives are secure
- event producers are accountable
- infrastructure operators are identifiable

Threat models are addressed at the validation layer.

---

## 15. Final Protocol Rule

> **No event, no value.  
> No value, no distribution.**

This rule overrides all interpretations.