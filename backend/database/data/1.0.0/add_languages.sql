BEGIN TRANSACTION;

INSERT INTO languages (name) VALUES 
  ('go'), -- 1
  ('typescript'), -- 2
  ('vbnet'), -- 3
  ('php'), -- 4
  ('python'); --5

COMMIT;