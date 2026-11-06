# VirtualEconomy

A virtual economy simulation API built with **Go (Golang)**, featuring user profiles, item management, wallet management, and financial transactions powered by a **double-entry ledger system**.

## Overview

**VirtualEconomy** is a backend API designed to simulate a simple game-like economy. It provides the foundation for managing virtual currencies, creating items, and processing financial transactions between users and the application.

The main goal of this project is to build a transparent and consistent economic system where every financial movement is traceable.

At its core, VirtualEconomy uses a **double-entry bookkeeping system**, where every transaction is recorded through balanced ledger entries. Wallet balances are derived from these entries, ensuring that the ledger remains the source of truth for the entire economy.

The system also implements atomic transactions and policy-driven ledger snapshots to maintain financial consistency while optimizing balance retrieval as the economy grows.

## Features

- **User Profiles** – Create and manage profiles within the virtual economy.
- **Item Management** – Create and manage virtual items.
- **Wallet System** – Retrieve wallet balances and transaction history.
- **Double-Entry Ledger** – Track every financial movement through balanced debit and credit entries.
- **Dedicated Fee Accounts** – Application fees are recorded through their own ledger entries, making platform revenue fully traceable.
- **Atomic Transactions** – Ensure that all ledger entries belonging to a transaction are committed together.
- **Ledger-Based Balances** – Wallet balances are derived from ledger movements rather than being treated as the source of truth.
- **Policy-Driven Snapshots** – Generate wallet snapshots according to configurable age and ledger-count thresholds.
- **Transaction Tracking** – Maintain a complete history of incoming and outgoing movements.
- **Financial Transparency** – Keep track of where money comes from, where it goes, and how fees are distributed.

## Architecture and Core Concepts

### 1. Profiles

Each user can create a profile representing their identity within the virtual economy.

Profiles are associated with wallets, allowing users to manage and track their virtual currency.

### 2. Items

The application allows users to create items that belong to the virtual economy.

### 3. Wallets

Every profile has a wallet that provides information about its financial state.

Through the wallet endpoint, users can retrieve:

- Current wallet balance.
- Ledger entries associated with their wallet.
- Incoming and outgoing movements.
- Transaction history.
- Associated transaction fees, where applicable.

Wallet balances are not independently maintained as the source of truth. Instead, they are calculated using ledger entries and optimized through snapshots.

### 4. Double-Entry Ledger

The double-entry ledger is the foundation of VirtualEconomy's financial architecture.

Every financial transaction generates corresponding debit and credit entries. The total debits and credits must always balance.

For example, imagine a user purchases an item for 100 virtual coins, with a 5% application fee.

The transaction generates the following ledger movements:

| Account | Debit | Credit |
|---|---:|---:|
| Buyer Wallet | 100 | - |
| Seller Wallet | - | 95 |
| Platform Fee Account | - | 5 |
| **Total** | **100** | **100** |

Platform fees are not simply deducted and forgotten. They have their own ledger entries, making it possible to independently track the application's revenue.

This architecture provides:

- Complete financial traceability.
- Transparent fee distribution.
- Consistent accounting.
- Auditable transaction history.
- A clear record of every unit of virtual currency.

### 5. Atomic Transactions

Financial consistency is a fundamental requirement of the system.

All ledger entries associated with a financial transaction are processed atomically.

This means that either every entry belonging to a transaction is successfully committed, or none of them are.

For example, when processing a purchase involving a buyer, seller, and platform fee account, all corresponding ledger movements must succeed together.

This prevents partially processed transactions and ensures that the double-entry accounting rules are preserved.

### 6. Policy-Driven Ledger Snapshots

To avoid recalculating an account's entire financial history whenever its balance is requested, VirtualEconomy implements a policy-driven snapshot mechanism.

Instead of generating snapshots after every transaction or following a fixed schedule, the system uses a configurable `WalletPolicy` to determine when a profile should receive a new snapshot.

The policy evaluates two conditions:

- **Maximum Age (`MaxAge`):** The maximum number of days allowed since the last snapshot.
- **Maximum Ledgers (`MaxLedgers`):** The maximum number of new ledger entries allowed since the last snapshot.

A new snapshot is generated when either condition is met, provided that new ledger entries exist since the previous snapshot.

The snapshot generation rule can be represented as:

```text
Should Generate Snapshot =
    Has New Ledger Entries AND
    (
        Snapshot Age >= MaxAge OR
        Ledgers Since Snapshot >= MaxLedgers
    )
```

This allows snapshot generation to adapt to different usage patterns:

- Profiles with frequent transactions can generate snapshots based on their ledger activity.
- Profiles with infrequent transactions can generate snapshots based on elapsed time.
- Profiles without new financial movements don't generate unnecessary snapshots.

Snapshots act as checkpoints in the ledger history.

Instead of iterating through every financial movement since a profile was created, the system can retrieve the latest snapshot and process only the subsequent ledger entries.

Conceptually:

```text
Current Balance =
    Latest Snapshot Balance +
    Ledger Movements Since Snapshot
```

Snapshots are an optimization mechanism and never replace the ledger as the source of truth.

By combining configurable snapshot policies with ledger-based balance calculation, VirtualEconomy reduces the amount of historical data that needs to be processed when retrieving wallet balances.

### 7. Transactions and Fees

Transactions are responsible for moving virtual currency between accounts.

Whenever a transaction occurs, the application records its corresponding ledger entries, including any applicable fees.

This allows the system to:

- Track the origin and destination of funds.
- Calculate and record application fees.
- Maintain a consistent financial history.
- Audit transactions and wallet movements.
- Prevent money from appearing or disappearing without an associated ledger operation.

## Technology Stack

VirtualEconomy uses a lightweight and performance-oriented technology stack.

| Technology | Purpose |
|---|---|
| Go (Golang) | Backend development |
| `net/http` | HTTP server and API routing using Go's standard library |
| PostgreSQL 17 | Relational database |
| sqlc | Type-safe, compile-time generated Go code from SQL queries |
| Docker | Application containerization |
| Docker Compose | Local development environment and service orchestration |

### Go Standard Library — net/http

The API uses Go's built-in `net/http` package instead of relying on third-party HTTP frameworks.

### PostgreSQL 17

PostgreSQL 17 is used as the relational database, providing transactional guarantees and reliable data persistence.

### sqlc

Database access is implemented using [sqlc](https://sqlc.dev/).

Instead of relying on runtime query builders or ORM abstractions, sqlc generates type-safe Go code directly from SQL queries.

This provides:

- Compile-time type safety for database operations.
- Explicit SQL queries and better control over database interactions.
- Reduced runtime overhead.
- Less boilerplate database code.
- Better maintainability and performance.

### Docker and Docker Compose

The project uses Docker and Docker Compose to simplify the development environment and service orchestration.

PostgreSQL 17 runs as a containerized service, making it easier to reproduce the development environment across different machines.

## API

VirtualEconomy exposes endpoints for interacting with the economy.

The main resources include:

| Resource | Description |
|---|---|
| Profiles | Create and manage user profiles. |
| Items | Create and manage virtual items. |
| Wallet | Retrieve wallet balances and ledger movements. |
| Transactions | Process and track economic transactions. |

Detailed endpoint documentation, request examples, and response schemas can be added as the API evolves.

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/)
- [Docker](https://www.docker.com/)
- Docker Compose

### Installation

Clone the repository:

```bash
git clone <repository-url>
```

Navigate to the project directory:

```bash
cd VirtualEconomy
```

### Start the Environment

Start the services using Docker Compose:

```bash
docker compose up -d
```

### Run the Application

Install Go dependencies:

```bash
go mod download
```

Start the API:

```bash
go run .
```

Depending on the project structure, you may need to adjust the command to point to the application's entry point.

## Project Goals

VirtualEconomy is an exploration of financial system concepts applied to game economies.

The project focuses on:

- Double-entry bookkeeping.
- Ledger-based financial systems.
- Atomic financial transactions.
- Wallet and balance management.
- Policy-driven ledger snapshots.
- Fee accounting and distribution.
- Type-safe database access.
- Transaction consistency and financial auditing.
- Building lightweight APIs with Go's standard library.

The ultimate goal is to build an economy where every financial movement is traceable, every transaction is consistent, and every wallet balance can be explained through its underlying ledger entries.

## Future Improvements

Potential features for future development include:

- Player-to-player transactions.
- Transaction reversal mechanisms.
- Wallet reconciliation and integrity checks.
- Economic statistics and reporting.
- Leaderboards and economy analytics.
- Automated financial consistency tests.
