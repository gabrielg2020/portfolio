BEGIN TRANSACTION;

INSERT INTO technologies (name) VALUES 
  ('svelte'), -- 1
  ('docker'), -- 2
  ('sqlite'), -- 3
  ('laminas'), -- 4
  ('aws'), -- 5
  ('mysql'); -- 6
  
COMMIT;