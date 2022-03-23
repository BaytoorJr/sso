package postgres

var migrations = []string{
	"create table if not exists $1.users_test (" +
		"id uuid not null, " +
		"login varchar(255), " +
		"password varchar(255), " +
		"created_at timestamp without time zone, " +
		"updated_at timestamp without time zone, " +
		"primary key(id));",

	"create table if not exists $1.users_fields_test (" +
		"id uuid not null, " +
		"field_name varchar(255), " +
		"field_value varchar(255), " +
		"foreign key(id) references users_test(id) on update cascade on delete cascade);",
}
