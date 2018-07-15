/* monolith table */
CREATE TABLE monolith (
  id BIGSERIAL PRIMARY KEY,

  value BIGINT,
  indexed_value BIGINT
);
CREATE INDEX monolith_indexed_value_idx ON monolith (indexed_value);

/* partition table */
CREATE TABLE partition (
  id BIGSERIAL,
  partition_key INTEGER,

  value BIGINT,
  indexed_value BIGINT
) PARTITION BY LIST(partition_key);

/* partition 0 */
CREATE TABLE partition_0
PARTITION OF partition
FOR VALUES IN (0);
ALTER TABLE partition_0 ADD PRIMARY KEY (id);
CREATE INDEX partition_0_indexed_value_idx ON partition_0 (indexed_value);

/* partition 1 */
CREATE TABLE partition_1
PARTITION OF partition
FOR VALUES IN (1);
ALTER TABLE partition_1 ADD PRIMARY KEY (id);
CREATE INDEX partition_1_indexed_value_idx ON partition_1 (indexed_value);

/* partition 2 */
CREATE TABLE partition_2
PARTITION OF partition
FOR VALUES IN (2);
ALTER TABLE partition_2 ADD PRIMARY KEY (id);
CREATE INDEX partition_2_indexed_value_idx ON partition_2 (indexed_value);

/* partition 3 */
CREATE TABLE partition_3
PARTITION OF partition
FOR VALUES IN (3);
ALTER TABLE partition_3 ADD PRIMARY KEY (id);
CREATE INDEX partition_3_indexed_value_idx ON partition_3 (indexed_value);

/* partition 4 */
CREATE TABLE partition_4
PARTITION OF partition
FOR VALUES IN (4);
ALTER TABLE partition_4 ADD PRIMARY KEY (id);
CREATE INDEX partition_4_indexed_value_idx ON partition_4 (indexed_value);