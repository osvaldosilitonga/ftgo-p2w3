CREATE TABLE transactions(
  id SERIAL PRIMARY KEY,
  user_id INT,
  product_id INT,
  quantity INT,
  total_amount INT
);