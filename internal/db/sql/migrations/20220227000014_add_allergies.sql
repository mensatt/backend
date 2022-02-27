-- migrate:up
INSERT INTO allergy (abbreviation, name) VALUES ('Wz', 'Weizen (Gluten)');
INSERT INTO allergy (abbreviation, name) VALUES ('Ro', 'Roggen (Gluten)');
INSERT INTO allergy (abbreviation, name) VALUES ('Ge', 'Gerste (Gluten)');
INSERT INTO allergy (abbreviation, name) VALUES ('Hf', 'Hafer (Gluten)');
INSERT INTO allergy (abbreviation, name) VALUES ('Kr', 'Krebstiere');
INSERT INTO allergy (abbreviation, name) VALUES ('Ei', 'Eier');
INSERT INTO allergy (abbreviation, name) VALUES ('Fi', 'Fisch');
INSERT INTO allergy (abbreviation, name) VALUES ('Er', 'Erdnüsse');
INSERT INTO allergy (abbreviation, name) VALUES ('So', 'Soja');
INSERT INTO allergy (abbreviation, name) VALUES ('Mi', 'Milch/Laktose');
INSERT INTO allergy (abbreviation, name) VALUES ('Man', 'Mandeln');
INSERT INTO allergy (abbreviation, name) VALUES ('Hs', 'Haselnüsse');
INSERT INTO allergy (abbreviation, name) VALUES ('Wa', 'Walnüsse');
INSERT INTO allergy (abbreviation, name) VALUES ('Ka', 'Cashewnüsse');
INSERT INTO allergy (abbreviation, name) VALUES ('Pe', 'Pekanüsse');
INSERT INTO allergy (abbreviation, name) VALUES ('Pa', 'Paranüsse');
INSERT INTO allergy (abbreviation, name) VALUES ('Pi', 'Pistazien');
INSERT INTO allergy (abbreviation, name) VALUES ('Mac', 'Macadamianüsse');
INSERT INTO allergy (abbreviation, name) VALUES ('Sel', 'Sellerie');
INSERT INTO allergy (abbreviation, name) VALUES ('Sen', 'Senf');
INSERT INTO allergy (abbreviation, name) VALUES ('Ses', 'Sesam');
INSERT INTO allergy (abbreviation, name) VALUES ('Su', 'Schwefeloxid/Sulfite');
INSERT INTO allergy (abbreviation, name) VALUES ('Lu', 'Lupinen');
INSERT INTO allergy (abbreviation, name) VALUES ('We', 'Weichtiere');

-- migrate:down
TRUNCATE TABLE allergy;
