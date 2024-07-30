-- query: INSERT_ACCOUNT
INSERT INTO "Account" (
    "AccountID", "Name", "Email", "Balance"
)VALUES(?, ?, ?, ?);

-- query: SELECT_ACCOUNTS
SELECT "AccountID", "Name", "Email", "Balance" FROM "Account"

-- query: SELECT_ACCOUNT_BY_ID
SELECT "AccountID", "Name", "Email", "Balance" FROM "Account"
WHERE "AccountID" = ?

-- query: INSERT_INTO_TRANSACTION
INSERT INTO "Transaction" (
    "TransactionID", "Type", "Amount", "Timestamp", "AccountID"
) VALUES (?, ?, ?, CURRENT_TIMESTAMP, ?);

-- query: UPDATE_ACCOUNT_BALANCE
UPDATE "Account" 
SET "Balance" = ?
WHERE "AccountID" = ?