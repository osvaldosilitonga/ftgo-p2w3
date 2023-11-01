ALTER TABLE transactions
  RENAME COLUMN users_id TO user_id;

ALTER TABLE transaction_details
  RENAME COLUMN transactions_id TO transaction_id;

ALTER TABLE transaction_details
  RENAME COLUMN products_id TO product_id;