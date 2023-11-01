-- Drop column
ALTER TABLE transactions
DROP COLUMN product_id,
DROP COLUMN quantity;

-- Add transaction_details table
CREATE TABLE transaction_details(
  id SERIAL PRIMARY KEY,
  transaction_id INT NOT NULL,
  product_id INT NOT NULL,
  qty INT NOT NULL,
  FOREIGN KEY (transaction_id) REFERENCES transactions(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);