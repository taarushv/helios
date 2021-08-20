CREATE TABLE transactions IF NOT EXISTS
(
    hash PRIMARY KEY varchar(66),
    nonce bigint,
    transaction_index bigint,
    from_address varchar(42),
    to_address varchar(42),
    value numeric(38),
    gas bigint,
    gas_price bigint,
    input text,
    receipt_cumulative_gas_used bigint,
    receipt_gas_used bigint,
    receipt_contract_address varchar(42),
    receipt_root varchar(66),
    receipt_status bigint,
    block_timestamp timestamp,
    block_number bigint,
    block_hash varchar(66),
    max_fee_per_gas bigint,
    max_priority_fee_per_gas bigint,
    transaction_type bigint,
    receipt_effective_gas_price bigint
);

CREATE TABLE liquidation_attempts IF NOT EXISTS 
(
    hash PRIMARY KEY varchar(66)
    liquidator varchar(42)
    loan_owner varchar(42)
    gas_price bigint,
    repayAmount bigint,
    time timestamp,
    collateralAddress varchar(42)
    debtAddress varchar(42)
    block_number bigint
    isMined boolean
)

CREATE TABLE trade_attempts IF NOT EXISTS
(
    hash PRIMARY KEY varchar(66)
    from varchar(42)
    to varchar(42)
    fromAmount bigInt
    toAmount bigInt
    isMined boolean
)
