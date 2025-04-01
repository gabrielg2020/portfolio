BEGIN TRANSACTION;

INSERT OR IGNORE INTO technologies (name) VALUES 
  ('Svelte'),   -- 1
  ('Docker'),   -- 2
  ('SQLite'),   -- 3
  ('Laminas'),  -- 4
  ('AWS'),      -- 5
  ('MySQL');    -- 6
  
COMMIT;