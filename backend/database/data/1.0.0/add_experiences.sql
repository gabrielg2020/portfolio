BEGIN TRANSACTION;

-- 1 - DVSA
INSERT INTO experiences (company, description, role, year_start, year_end, point_one, point_two, point_three )
VALUES ('DVSA', 'HERE!!! sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus', 'Software Engineer', '2024', 'Current', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.');

INSERT INTO experience_languages (experience_id, language_id, position)
VALUES 
    (1, 4, 1), -- php
    (1, 2, 2); -- typescript

INSERT INTO experience_technologies (experience_id, technology_id, position)
VALUES 
    (1, 2, 1), -- docker
    (1, 4, 2), -- laminas
    (1, 5, 3), -- aws
    (1, 6, 4); -- mysql

-- 2 - University of Roehampton
INSERT INTO experiences (company, description, role, year_start, year_end, point_one, point_two, point_three )
VALUES ('University of Roehampton', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus', 'B.Sc. Digital Technology Solutions', '2024', 'Current', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.');

INSERT INTO experience_languages (experience_id, language_id, position)
VALUES 
    (2, 5, 1), -- python
    (2, 1, 2); -- go

INSERT INTO experience_technologies (experience_id, technology_id, position)
VALUES 
    (2, 5, 1), -- aws
    (2, 6, 2); -- mysql

-- 3 - LEGO
INSERT INTO experiences (company, description, role, year_start, year_end, point_one, point_two, point_three )
VALUES ('LEGO', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus', 'Retail Marketing Intern', '2022', '2022', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.');

COMMIT;