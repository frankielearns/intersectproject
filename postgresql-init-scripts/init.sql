GRANT CONNECT ON DATABASE stock TO stockuser;
GRANT USAGE ON SCHEMA public TO stockuser;
CREATE TABLE stockhistory (
  id SERIAL PRIMARY KEY,
  stocktick VARCHAR(255),
  price numeric,
  username VARCHAR(255)
);

GRANT SELECT, INSERT, UPDATE ON stockhistory TO stockuser;
grant all on sequence stockhistory_id_seq to stockuser;