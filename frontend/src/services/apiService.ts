export interface Project {
  title: string,
  smallDescription: string,
  largeDescription: string[],
  githubLink: string,
  languages: string[],
  technologies: string[]
}

export interface Experience extends Project {
  company: string,
  yearStart: string,
  yearEnd: string
}

export const GetProjects = async (): Promise<Project[]> => {
  let projects: Project[] = []
  try {
    const response = await fetch("/api/projects")
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const data = await response.json()
    projects = data.projects
  } catch (err) {
    console.error("Failed to fetch projects:", err);
  }

  return projects
}

export const GetExperiences = async (): Promise<Experience[]> => {
  let experiences: Experience[] = []
  try {
    const response = await fetch("/api/experiences")
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const data = await response.json()
    experiences = data.experiences
  } catch (err) {
    console.error("Failed to fetch experiences:", err);
  }

  return experiences
}