BEGIN TRANSACTION;

-- goto
INSERT INTO projects (title, description, github_link)
VALUES ('goto', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus.', 'https://github.com/gabrielg2020');

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (last_insert_rowid(), 1, 1); -- go

-- chess-api
INSERT INTO projects (title, description, github_link)
VALUES ('ChessAPI','Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus.', 'https://github.com/gabrielg2020');

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (last_insert_rowid(), 1, 1); -- go

-- maze-visualiser
INSERT INTO projects (title, description, github_link)
VALUES ('Maze Visualiser','Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus.', 'https://github.com/gabrielg2020');

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (last_insert_rowid(), 5, 1); -- vbnet

-- portfolio
INSERT INTO projects (title, description, github_link)
VALUES ('Portfolio','Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus.', 'https://github.com/gabrielg2020');

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (last_insert_rowid(), 1, 1), -- go
    (last_insert_rowid(), 2, 2), -- typescript
    (last_insert_rowid(), 4, 3); -- sql

INSERT INTO project_technologies (project_id, technology_id, position)
VALUES 
    (last_insert_rowid(), 1, 1), -- svelte
    (last_insert_rowid(), 2, 2); -- docker

COMMIT;