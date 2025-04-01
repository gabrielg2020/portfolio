export type Experience = {
  type: string,
  organisation: string,
  role: string,
  startYear: string,
  endYear: string,
  description: string,
  languages: string[],
  technologies: string[]
}

export async function GetExperiece(): Promise<Experience[]> {
  let experiences: Experience[] = []

  try {
    const response = await fetch("/api/experience")
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const data = await response.json();
    experiences = data.experiences;
  } catch (err) {
    console.error("Failed to fetch experiences:", err);
    throw err; // Re-throw the error so it propagates to the caller
  }

  return experiences
}
