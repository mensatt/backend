-- migrate:up
INSERT INTO allergy (abbreviation, name) VALUES ('1', 'mit Farbstoff');
INSERT INTO allergy (abbreviation, name) VALUES ('2', 'mit Koffein');
INSERT INTO allergy (abbreviation, name) VALUES ('4', 'mit Konservierungsstoffen');
INSERT INTO allergy (abbreviation, name) VALUES ('5', 'mit Süßungsmittel');
INSERT INTO allergy (abbreviation, name) VALUES ('7', 'mit Antioxidationsmittel');
INSERT INTO allergy (abbreviation, name) VALUES ('8', 'mit Geschmacksverstärker');
INSERT INTO allergy (abbreviation, name) VALUES ('9', 'geschwefelt');
INSERT INTO allergy (abbreviation, name) VALUES ('10', 'geschwärzt');
INSERT INTO allergy (abbreviation, name) VALUES ('11', 'gewachst');
INSERT INTO allergy (abbreviation, name) VALUES ('12', 'mit Phosphat');
INSERT INTO allergy (abbreviation, name) VALUES ('13', 'mit einer Phenylalaninquelle');
INSERT INTO allergy (abbreviation, name) VALUES ('30', 'mit Fettglasur');

-- migrate:down
TRUNCATE TABLE tag;
