
package queries

import "fmt"

func GetExperiences(limit int) string {
	languagesSubquery := `
(
  SELECT GROUP_CONCAT(l.name, ', ')
  FROM experience_languages el
    JOIN languages l ON el.language_id = l.language_id
  WHERE el.experience_id = e.experience_id
  ORDER BY el.position ASC
) AS languages,
	`

	technologiesSubquery := `
(
  SELECT GROUP_CONCAT(t.name, ', ')
  FROM experience_technologies et
    JOIN technologies t ON et.technology_id = t.technology_id
  WHERE et.experience_id = e.experience_id
  ORDER BY et.position ASC
) AS technologies
	`

	limitSubString := ";"
	if limit != 0 {
		limitSubString = fmt.Sprintf(`LIMIT
	%d;`, limit)
	}

	// Build full query
	return fmt.Sprintf(`SELECT
    e.experience_id,
    e.organisation,
    e.role,
    e.start_year,
    e.end_year,
    e.description,
		%s
		%s
FROM
	experiences e
ORDER BY
	e.created_at DESC
%s
`, languagesSubquery, technologiesSubquery, limitSubString)
}
