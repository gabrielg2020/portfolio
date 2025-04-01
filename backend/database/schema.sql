DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS experiences;
DROP TABLE IF EXISTS languages;
DROP TABLE IF EXISTS technologies;
DROP TABLE IF EXISTS project_languages;
DROP TABLE IF EXISTS project_technologies;
DROP TABLE IF EXISTS experience_languages;
DROP TABLE IF EXISTS experience_technologies;

CREATE TABLE projects (
  project_id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  github_url TEXT,
  live_url TEXT,
  display_order INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE experiences (
  experience_id INTEGER PRIMARY KEY AUTOINCREMENT,
  type TEXT NOT NULL,
  organisation TEXT NOT NULL,
  role TEXT NOT NULL,
  start_year TEXT NOT NULL,
  end_year TEXT NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE languages (
  language_id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE technologies (
  technology_id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE project_languages (
  project_id INTEGER,
  language_id INTEGER,
  position INTEGER CHECK (position BETWEEN 1 AND 4),
  PRIMARY KEY (project_id, position),
  UNIQUE (project_id, language_id),
  FOREIGN KEY (project_id) REFERENCES projects(project_id) ON DELETE CASCADE,
  FOREIGN KEY (language_id) REFERENCES languages(language_id) ON DELETE RESTRICT
);

CREATE TABLE project_technologies (
  project_id INTEGER,
  technology_id INTEGER,
  position INTEGER CHECK (position BETWEEN 1 AND 4),
  PRIMARY KEY (project_id, position),
  UNIQUE (project_id, technology_id),
  FOREIGN KEY (project_id) REFERENCES projects(project_id) ON DELETE CASCADE,
  FOREIGN KEY (technology_id) REFERENCES technologies(technology_id) ON DELETE RESTRICT
);

CREATE TABLE experience_languages (
  experience_id INTEGER,
  language_id INTEGER,
  position INTEGER CHECK (position BETWEEN 1 AND 4),
  PRIMARY KEY (experience_id, position),
  UNIQUE (experience_id, language_id),
  FOREIGN KEY (experience_id) REFERENCES experiences(experience_id) ON DELETE CASCADE,
  FOREIGN KEY (language_id) REFERENCES languages(language_id) ON DELETE RESTRICT
);

CREATE TABLE experience_technologies (
  experience_id INTEGER,
  technology_id INTEGER,
  position INTEGER CHECK (position BETWEEN 1 AND 4),
  PRIMARY KEY (experience_id, position),
  UNIQUE (experience_id, technology_id),
  FOREIGN KEY (experience_id) REFERENCES experiences(experience_id) ON DELETE CASCADE,
  FOREIGN KEY (technology_id) REFERENCES technologies(technology_id) ON DELETE RESTRICT
);

-- Auto-update display_order for new projects if not specified
CREATE TRIGGER auto_display_order_projects
AFTER INSERT ON projects
WHEN NEW.display_order IS NULL
BEGIN
    UPDATE projects 
    SET display_order = (SELECT COALESCE(MAX(display_order), 0) + 1 FROM projects)
    WHERE project_id = NEW.project_id;
END;

-- Create triggers for project_languages
CREATE TRIGGER auto_position_project_languages
AFTER INSERT ON project_languages
WHEN NEW.position IS NULL
BEGIN
    UPDATE project_languages 
    SET position = (SELECT COALESCE(MAX(position), 0) + 1 
                   FROM project_languages 
                   WHERE project_id = NEW.project_id)
    WHERE project_id = NEW.project_id AND language_id = NEW.language_id;
END;

-- Create triggers for project_technologies
CREATE TRIGGER auto_position_project_technologies
AFTER INSERT ON project_technologies
WHEN NEW.position IS NULL
BEGIN
    UPDATE project_technologies 
    SET position = (SELECT COALESCE(MAX(position), 0) + 1 
                   FROM project_technologies 
                   WHERE project_id = NEW.project_id)
    WHERE project_id = NEW.project_id AND technology_id = NEW.technology_id;
END;

-- Create triggers for experience_languages
CREATE TRIGGER auto_position_experience_languages
AFTER INSERT ON experience_languages
WHEN NEW.position IS NULL
BEGIN
    UPDATE experience_languages 
    SET position = (SELECT COALESCE(MAX(position), 0) + 1 
                   FROM experience_languages 
                   WHERE experience_id = NEW.experience_id)
    WHERE experience_id = NEW.experience_id AND language_id = NEW.language_id;
END;

-- Create triggers for experience_technologies
CREATE TRIGGER auto_position_experience_technologies
AFTER INSERT ON experience_technologies
WHEN NEW.position IS NULL
BEGIN
    UPDATE experience_technologies 
    SET position = (SELECT COALESCE(MAX(position), 0) + 1 
                   FROM experience_technologies 
                   WHERE experience_id = NEW.experience_id)
    WHERE experience_id = NEW.experience_id AND technology_id = NEW.technology_id;
END;

-- Helper triggers to re-order positions after deletion
CREATE TRIGGER reorder_project_languages_after_delete
AFTER DELETE ON project_languages
BEGIN
    UPDATE project_languages
    SET position = (
        SELECT COUNT(*) 
        FROM project_languages AS pl 
        WHERE pl.project_id = OLD.project_id 
        AND pl.position < project_languages.position
    ) + 1
    WHERE project_id = OLD.project_id AND position > OLD.position;
END;

CREATE TRIGGER reorder_project_technologies_after_delete
AFTER DELETE ON project_technologies
BEGIN
    UPDATE project_technologies
    SET position = (
        SELECT COUNT(*) 
        FROM project_technologies AS pt 
        WHERE pt.project_id = OLD.project_id 
        AND pt.position < project_technologies.position
    ) + 1
    WHERE project_id = OLD.project_id AND position > OLD.position;
END;

CREATE TRIGGER reorder_experience_languages_after_delete
AFTER DELETE ON experience_languages
BEGIN
    UPDATE experience_languages
    SET position = (
        SELECT COUNT(*) 
        FROM experience_languages AS el 
        WHERE el.experience_id = OLD.experience_id 
        AND el.position < experience_languages.position
    ) + 1
    WHERE experience_id = OLD.experience_id AND position > OLD.position;
END;

CREATE TRIGGER reorder_experience_technologies_after_delete
AFTER DELETE ON experience_technologies
BEGIN
    UPDATE experience_technologies
    SET position = (
        SELECT COUNT(*) 
        FROM experience_technologies AS et 
        WHERE et.experience_id = OLD.experience_id 
        AND et.position < experience_technologies.position
    ) + 1
    WHERE experience_id = OLD.experience_id AND position > OLD.position;
END;