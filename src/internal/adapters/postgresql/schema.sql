CREATE TABLE IF NOT EXISTS profile (
    id TEXT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS exchange (
    id TEXT PRIMARY KEY
);

CREATE TYPE movement_type AS ENUM (
    'send',
    'receive',
    'external_send',
    'external_receive',
    'fee'
);

CREATE TABLE IF NOT EXISTS profile_ledge (
    profile_id TEXT NOT NULL,
    exchange_id TEXT NOT NULL,
    amount INT,
    movement movement_type NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_profile_ledge_profile FOREIGN KEY (profile_id) REFERENCES profile(id)
);

CREATE INDEX IF NOT EXISTS idx_profile_ledge_profile_id ON profile_ledge(profile_id);

CREATE TABLE IF NOT EXISTS wallet_daily_snapshot (
    profile_id TEXT NOT NULL,
    effective_at DATE NOT NULL DEFAULT CURRENT_DATE,
    balance INT,
    CONSTRAINT pk_wallet_daily_snapshot PRIMARY KEY (profile_id, effective_at),
    CONSTRAINT fk_wallet_daily_snapshot_profile FOREIGN KEY (profile_id) REFERENCES profile(id)
);

CREATE TABLE IF NOT EXISTS item (
    id TEXT NOT NULL,
    creator_id TEXT NOT NULL,
    name TEXT NOT NULL,
    value INT,
    created_at DATE DEFAULT CURRENT_DATE,
    CONSTRAINT fk_item_creator FOREIGN KEY (creator_id) REFERENCES profile(id)
);

CREATE TABLE IF NOT EXISTS inventory (
    profile_id TEXT NOT NULL,
    item_id TEXT NOT NULL,
    purchased_at DATE DEFAULT CURRENT_DATE,
    CONSTRAINT fk_inventory_profile FOREIGN KEY (profile_id) REFERENCES profile(id),
    CONSTRAINT fk_inventory_item FOREIGN KEY (item_id) REFERENCES item(id),
    CONSTRAINT uq_inventory_profile_item UNIQUE (profile_id, item_id)
);