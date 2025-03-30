BEGIN TRANSACTION;

-- 1 - goto
INSERT INTO projects (title, description, github_link, point_one, point_two, point_three)
VALUES ('goto', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus.', 'https://github.com/gabrielg2020', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'HERE', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.');

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (1, 1, 1); -- go

-- 2 - chess-api
INSERT INTO projects (title, description, github_link, point_one, point_two, point_three)
VALUES ('ChessAPI','Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus.', 'https://github.com/gabrielg2020', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.');

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (2, 1, 1); -- go

-- 3 - maze-visualiser
INSERT INTO projects (title, description, github_link, point_one, point_two, point_three)
VALUES ('Maze Visualiser','Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus.', 'https://github.com/gabrielg2020', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.');

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (3, 3, 1); -- vbnet

-- 4- portfolio
INSERT INTO projects (title, description, github_link, point_one, point_two, point_three)
VALUES ('Portfolio','Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus.', 'https://github.com/gabrielg2020', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.');

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (4, 1, 1), -- go
    (4, 2, 2); -- typescript

INSERT INTO project_technologies (project_id, technology_id, position)
VALUES 
    (4, 1, 1), -- svelte
    (4, 2, 2), -- docker
    (4, 3, 3); -- sqlite

COMMIT;