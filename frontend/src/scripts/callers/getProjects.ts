export type Project = {
  title: string,
  description: string,
  githubUrl: string,
  liveUrl: string,
  languages: string[],
  technologies: string[]
}

export async function GetProjects(): Promise<Project[]> {
  let projects: Project[] = []

  try {
    const response = await fetch("/api/projects")
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const data = await response.json();
    projects = data.projects;
  } catch (err) {
    console.error("Failed to fetch projects:", err);
    throw err; // Re-throw the error so it propagates to the caller
  }

  return projects
}
