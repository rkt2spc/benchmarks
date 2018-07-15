/* uuid_primary_key table */
CREATE TABLE uuid_primary_key (
  id UUID PRIMARY KEY,

  value BIGINT,
  indexed_value BIGINT
);
CREATE INDEX uuid_primary_key_idx ON uuid_primary_key (indexed_value);

/* integer_primary_key table */
CREATE TABLE integer_primary_key (
  id BIGSERIAL,

  value BIGINT,
  indexed_value BIGINT
);
CREATE INDEX integer_primary_key_idx ON integer_primary_key (indexed_value);
