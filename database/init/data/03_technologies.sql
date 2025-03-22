BEGIN TRANSACTION;

INSERT INTO technologies (name) VALUES 
  ('svelte'), -- 1
  ('docker'), -- 2
  ('react'); -- 3

COMMIT;