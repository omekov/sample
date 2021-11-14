package domain

import sq "github.com/Masterminds/squirrel"

func (User) tableName() string {
	return "users"
}

func (u *User) InsertScript() (string, []interface{}, error) {
	return sq.Insert(u.tableName()).
		Columns(
			"username",
			"password",
			"firstname",
			"lastname",
			"birth_date",
			"gender",
		).
		Values(
			u.Name,
			u.Password,
			u.FirstName,
			u.LastName,
			u.BirthDate,
			u.GenderID,
		).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
}

func (u *User) SelectOneScript() (string, []interface{}, error) {
	return sq.Select("name", "firstname", "lastname", "birth_date").
		From(u.tableName()).
		Where(sq.Eq{"id": u.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
}

func (u *User) SelectOneByNameScript() (string, []interface{}, error) {
	return sq.Select("name", "firstname", "lastname", "birth_date").
		From(u.tableName()).
		Where(sq.Eq{"name": u.Name}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
}

func (u *User) SelectScript(page int) (string, []interface{}, error) {
	limit, offset := getOffsetLimitTen(page)
	return sq.Select("name", "firstname", "lastname", "birth_date").
		From(u.tableName()).
		PlaceholderFormat(sq.Dollar).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
}
