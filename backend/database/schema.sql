DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS languages;
DROP TABLE IF EXISTS technologies;
DROP TABLE IF EXISTS project_languages;
DROP TABLE IF EXISTS project_technologies;

CREATE TABLE projects (
  project_id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  github_link TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE languages (
  language_id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE technologies (
  technolgoy_id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE project_languages (
  project_id INTEGER,
  language_id INTEGER,
  position INTEGER NOT NULL CHECK (position BETWEEN 1 AND 4),
  PRIMARY KEY (project_id, position),
  UNIQUE (project_id, language_id),
  FOREIGN KEY (project_id) REFERENCES projects(project_id) ON DELETE CASCADE,
  FOREIGN KEY (language_id) REFERENCES languages(language_id) ON DELETE RESTRICT
);

CREATE TABLE project_technologies (
  project_id INTEGER,
  technology_id INTEGER,
  position INTEGER NOT NULL CHECK (position BETWEEN 1 AND 4),
  PRIMARY KEY (project_id, position),
  UNIQUE (project_id, technology_id),
  FOREIGN KEY (project_id) REFERENCES projects(project_id) ON DELETE CASCADE,
  FOREIGN KEY (technology_id) REFERENCES technologies(technology_id) ON DELETE RESTRICT
);