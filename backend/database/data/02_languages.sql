BEGIN TRANSACTION;

INSERT INTO languages (name) VALUES 
  ('go'), -- 1
  ('typescript'), -- 2
  ('c'), -- 3
  ('sql'), -- 4
  ('vbnet'), -- 5
  ('php'); -- 6

COMMIT;