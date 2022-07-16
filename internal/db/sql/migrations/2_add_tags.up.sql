-- migrate:up

-- Tags
INSERT INTO tag (key, name, description) VALUES ('1', 'Farbstoff', 'mit Farbstoff');
INSERT INTO tag (key, name, description) VALUES ('2', 'Koffein', 'mit Koffein');
INSERT INTO tag (key, name, description) VALUES ('4', 'Konservierungsstoffen', 'mit Konservierungsstoffen');
INSERT INTO tag (key, name, description) VALUES ('5', 'Süßungsmittel', 'mit Süßungsmittel');
INSERT INTO tag (key, name, description) VALUES ('7', 'Antioxidationsmittel', 'mit Antioxidationsmittel');
INSERT INTO tag (key, name, description) VALUES ('8', 'Geschmacksverstärker', 'mit Geschmacksverstärker');
INSERT INTO tag (key, name, description) VALUES ('9', 'geschwefelt', 'geschwefelt');
INSERT INTO tag (key, name, description) VALUES ('10', 'geschwärzt', 'geschwärzt');
INSERT INTO tag (key, name, description) VALUES ('11', 'gewachst', 'gewachst');
INSERT INTO tag (key, name, description) VALUES ('12', 'Phosphat', 'mit Phosphat');
INSERT INTO tag (key, name, description) VALUES ('13', 'Phenylalaninquelle', 'mit einer Phenylalaninquelle');
INSERT INTO tag (key, name, description) VALUES ('30', 'Fettglasur', 'mit Fettglasur');

-- Icon-Tags
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('S', 'Schwein', 'Schweinefleisch', '🐷', 'HIGHER', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('R', 'Rind', 'Rindfleisch', '🐮', 'HIGHER', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('G', 'Geflügel', 'Geflügel', '🐔', 'HIGHER', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('L', 'Lamm', 'Lamm', '🐑', 'HIGHER', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('W', 'Wild', 'Wildfleisch', '🐈', 'HIGHER', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('F', 'Fisch', 'Fisch', '🐟', 'HIGHER', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('V', 'Vegetarisch', 'Vegetarisch', '🥕', 'HIGHER', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('Veg', 'Vegan', 'Vegan', '🌱', 'HIGHER', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('MSC', 'MSC Fish', 'MSC Fish', '🐟', 'HIGHER', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Gf', 'Glutenfrei', 'Glutenfrei', 'Glutenfrei', 'MEDIUM', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('CO2', 'CO2-Neutral', 'CO2-Neutral', 'CO2-Neutral', 'MEDIUM', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('B', 'Bio', 'Bio', 'Bio', 'HIGHER', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('MV', 'MensaVital', 'MensaVital', 'MensaVital', 'MEDIUM', TRUE);
-- missing: alcohol icon

-- Allergies
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Wz', 'Weizen', 'Weizen (Gluten)', '🌾', 'MEDIUM', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Ro', 'Roggen', 'Roggen (Gluten)', '', 'MEDIUM', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Ge', 'Gerste', 'Gerste (Gluten)', '', 'MEDIUM', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Hf', 'Hafer', 'Hafer (Gluten)', '', 'MEDIUM', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Kr', 'Krebstiere', 'Krebstiere', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Ei', 'Eier', 'Eier', '🥚', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Fi', 'Fisch', 'Fisch', '🐟', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Er', 'Erdnüsse', 'Erdnüsse', '🥜', 'MEDIUM', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('So', 'Soja', 'Soja', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Mi', 'Milch/Laktose', 'Milch/Laktose', '', 'MEDIUM', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Man', 'Mandeln', 'Mandeln', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Hs', 'Haselnüsse', 'Haselnüsse', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Wa', 'Walnüsse', 'Walnüsse', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Ka', 'Cashewnüsse', 'Cashewnüsse', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Pe', 'Pekanüsse', 'Pekanüsse', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Pa', 'Paranüsse', 'Paranüsse', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Pi', 'Pistazien', 'Pistazien', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Mac', 'Macadamianüsse', 'Macadamianüsse', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Sel', 'Sellerie', 'Sellerie', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Sen', 'Senf', 'Senf', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Ses', 'Sesam', 'Sesam', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Su', 'Schwefeloxid/Sulfite', 'Schwefeloxid/Sulfite', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Lu', 'Lupinen', 'Lupinen', '', 'LOWER', TRUE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('We', 'Weichtiere', 'Weichtiere', '', 'LOWER', TRUE, TRUE);
