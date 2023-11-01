-- Add column
ALTER TABLE transactions
ADD COLUMN product_id INT,
ADD COLUMN quantity INT;

-- Drop transaction_details table
DROP TABLE transaction_details;