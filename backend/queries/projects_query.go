package queries

import "fmt"

func GetProjects(limit int) string {
	languagesSubquery := `
(
  SELECT GROUP_CONCAT(l.name, ', ')
  FROM project_languages pl
    JOIN languages l ON pl.language_id = l.language_id
  WHERE pl.project_id = p.project_id
  ORDER BY pl.position ASC
) AS languages,
	`

	technologiesSubquery := `
(
  SELECT GROUP_CONCAT(t.name, ', ')
  FROM project_technologies pt
    JOIN technologies t ON pt.technology_id = t.technology_id
  WHERE pt.project_id = p.project_id
  ORDER BY pt.position ASC
) AS technologies
	`

	limitSubString := ";"
	if limit != 0 {
		limitSubString = fmt.Sprintf(`LIMIT
	%d;`, limit)
	}

	// Build full query
	return fmt.Sprintf(`SELECT
    p.project_id,
    p.title,
    p.description,
    p.github_url,
    p.live_url,
		%s
		%s
FROM
	projects p
ORDER BY
	p.display_order ASC,
	p.created_at DESC
%s
`, languagesSubquery, technologiesSubquery, limitSubString)
}
