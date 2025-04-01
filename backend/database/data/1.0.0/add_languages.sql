BEGIN TRANSACTION;

INSERT INTO languages (name) VALUES 
  ('Go'), -- 1
  ('TypeScript'), -- 2
  ('VB.Net'), -- 3
  ('PHP'), -- 4
  ('Python'); --5

COMMIT;