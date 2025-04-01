BEGIN TRANSACTION;

-- 1 - goto
INSERT INTO projects (title, description, github_url, live_url)
VALUES ('goto', 'A lightweight terminal tool that simplifies directory navigation.', 'https://github.com/gabrielg202/goto', NULL);

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (1, 1, 1); -- go

-- 2 - maze-visualiser
INSERT INTO projects (title, description, github_url, live_url)
VALUES ('Maze Visualiser','The Maze Visualizer is an academic desktop application that demonstrates various maze generation and solving algorithms in real-time with customizable parameters including entry/exit points and color schemes.', 'https://github.com/gabrielg2020/maze-visualiser', 'https://www.youtube.com/playlist?list=PL3YAEsPABRrxRmNFIRFjes7XDmWsjn9kN');

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (2, 3, 1); -- vbnet

-- 3- portfolio
INSERT INTO projects (title, description, github_url, live_url)
VALUES ('Portfolio','I designed and developed a comprehensive personal portfolio website to showcase my professional projects, technical skills, and relevant work experience, providing visitors with an interactive platform to explore my capabilities and accomplishments.', 'https://github.com/gabrielg2020', 'https://github.com/gabrielg2020/portfolio');

INSERT INTO project_languages (project_id, language_id, position)
VALUES 
    (3, 1, 1), -- go
    (3, 2, 2); -- typescript

INSERT INTO project_technologies (project_id, technology_id, position)
VALUES 
    (3, 1, 1), -- svelte
    (3, 2, 2), -- docker
    (3, 3, 3); -- sqlite

COMMIT;