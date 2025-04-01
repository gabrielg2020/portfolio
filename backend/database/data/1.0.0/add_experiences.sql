BEGIN TRANSACTION;

-- 1 - DVSA
INSERT INTO experiences (type, organisation, role, start_year, end_year, description)
VALUES ('work', 'DVSA', 'Software Engineer', '2024', 'Current', 'Developed and maintained the Vehicle Operator Licence service with modern development approaches that enhanced user experience for over 20,000 UK businesses, improved system security through containerisation.');

-- No need to specify positions anymore
INSERT INTO experience_languages (experience_id, language_id)
VALUES 
    (1, 4), -- php
    (1, 2); -- typescript

INSERT INTO experience_technologies (experience_id, technology_id)
VALUES 
    (1, 2), -- docker
    (1, 4), -- laminas
    (1, 5), -- aws
    (1, 6); -- mysql

-- 2 - University of Roehampton
INSERT INTO experiences (type, organisation, role, start_year, end_year, description)
VALUES ('education', 'University of Roehampton', 'BSc Digital Technology Solutions', '2024', 'Current', 'Designed and developed business-focused applications and systems architecture like database functionality and developer tools, while demonstrating strong theoretical understanding through academic writing on enterprise systems, maintaining performance consistent with upper second-class honors.');

INSERT INTO experience_languages (experience_id, language_id)
VALUES 
    (2, 5), -- python
    (2, 1); -- go

INSERT INTO experience_technologies (experience_id, technology_id)
VALUES 
    (2, 5), -- aws
    (2, 6); -- mysql

-- 3 - LEGO
INSERT INTO experiences (type, organisation, role, start_year, end_year, description)
VALUES ('work', 'LEGO', 'Retail Marketing Intern', '2022', '2022', 'Collaborated with a team of six to develop an innovative in-store retail activation concept for LEGO, gathering strategic insights through specialist consultations and delivering a comprehensive presentation that demonstrated seamless integration of data capture solutions within existing infrastructure, complete with defined roles and implementation strategies.');

-- 4 - Fullbrook 6th Form
INSERT INTO experiences (type, organisation, role, start_year, end_year, description)
VALUES ('education', 'Fullbrook 6th Form', 'A Level Computer Science, Maths, Physics', '2021', '2023', 'Studied foundational topics like calculus, Newtonian mechanics, and data structures and algorithms, achieving a C grade average overall while excelling with a perfect score of 100% in computer science coursework assignment.');

INSERT INTO experience_languages (experience_id, language_id)
VALUES 
    (4, 5), -- python
    (4, 3); -- vbnet

INSERT INTO experience_technologies (experience_id, technology_id)
VALUES 
    (4, 6); -- mysql

COMMIT;