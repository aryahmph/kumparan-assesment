CREATE INDEX text_search_idx ON articles USING GIN (text_search);