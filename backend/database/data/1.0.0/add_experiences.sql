BEGIN TRANSACTION;

-- 1 - DVSA
INSERT INTO experiences (organisation, role, start_year, end_year, description )
VALUES ('DVSA', 'Software Engineer', '2024', 'Current', 'Developed and maintained the Vehicle Operator Licence service with modern development approaches that enhanced user experience for over 20,000 UK businesses, improved system security through containerisation. ');

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
INSERT INTO experiences (organisation, role, start_year, end_year, description )
VALUES ('University of Roehampton', 'B.Sc. Digital Technology Solutions', '2024', 'Current', 'Designed and developed business-focused applications and systems architecture like database functionallity and developer tools, while demonstrating strong theoretical understanding through academic writing on enterprise systems, maintaining performance consistent with upper second-class honors.');

INSERT INTO experience_languages (experience_id, language_id, position)
VALUES 
    (2, 5, 1), -- python
    (2, 1, 2); -- go

INSERT INTO experience_technologies (experience_id, technology_id, position)
VALUES 
    (2, 5, 1), -- aws
    (2, 6, 2); -- mysql

-- 3 - LEGO
INSERT INTO experiences (organisation, role, start_year, end_year, description )
VALUES ('LEGO', 'Retail Marketing Intern', '2022', '2022', 'Collaborated with a team of six to develop an innovative in-store retail activation concept for LEGO, gathering strategic insights through specialist consultations and delivering a comprehensive presentation that demonstrated seamless integration of data capture solutions within existing infrastructure, complete with defined roles and implementation strategies.');

COMMIT;