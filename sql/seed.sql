INSERT INTO accounts (account_id, balance) VALUES ('mauricio', 1000.00) ON CONFLICT DO NOTHING;
INSERT INTO login_details (account_id, password_hash) VALUES ('mauricio', '$2y$10$dUFdsYDq4TSm/Sna/3rBtOImKQzHC9sFt6dqfvrqnk5up0pc5kK.i') ON CONFLICT DO NOTHING;
