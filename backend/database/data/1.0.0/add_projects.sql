BEGIN TRANSACTION;

-- 1 - maze-visualiser
INSERT INTO projects (title, description, github_url, live_url, display_order)
VALUES ('Maze Visualiser','The Maze Visualizer is an academic desktop application that demonstrates various maze generation and solving algorithms in real-time with customizable parameters including entry/exit points and color schemes.', 'https://github.com/gabrielg2020/maze-visualiser', 'https://www.youtube.com/playlist?list=PL3YAEsPABRrxRmNFIRFjes7XDmWsjn9kN', 2);

INSERT INTO project_languages (project_id, language_id)
VALUES 
    (1, 3); -- vbnet

-- 2 - portfolio
INSERT INTO projects (title, description, github_url, live_url, display_order)
VALUES ('Portfolio','I designed and developed a comprehensive personal portfolio website to showcase my professional projects, technical skills, and relevant work experience, providing visitors with an interactive platform to explore my capabilities and accomplishments.', 'https://github.com/gabrielg2020', 'https://github.com/gabrielg2020/portfolio', 1);

INSERT INTO project_languages (project_id, language_id)
VALUES 
    (2, 1), -- go
    (2, 2); -- typescript

INSERT INTO project_technologies (project_id, technology_id)
VALUES 
    (2, 1), -- svelte
    (2, 2), -- docker
    (2, 3); -- sqlite

COMMIT;