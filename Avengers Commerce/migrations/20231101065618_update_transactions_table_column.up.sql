ALTER TABLE transactions
  RENAME COLUMN user_id TO users_id;

ALTER TABLE transaction_details
  RENAME COLUMN transaction_id TO transactions_id;
  
ALTER TABLE transaction_details
  RENAME COLUMN product_id TO products_id;