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
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('S', 'Schwein', 'Schweinefleisch', '🐷', 'HIGHER', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('R', 'Rind', 'Rindfleisch', '🐮', 'HIGHER', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('G', 'Geflügel', 'Geflügel', '🐔', 'HIGHER', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('L', 'Lamm', 'Lamm', '🐑', 'HIGHER', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('W', 'Wild', 'Wildfleisch', '🐈', 'HIGHER', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('F', 'Fisch', 'Fisch', '🐟', 'HIGHER', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('V', 'Vegetarisch', 'Vegetarisch', '🥕', 'HIGHER', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('Veg', 'Vegan', 'Vegan', '🌱', 'HIGHER', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('MSC', 'MSC Fish', 'MSC Fish', '🐟', 'HIGHER', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Gf', 'Glutenfrei', 'Glutenfrei', 'Glutenfrei', 'MEDIUM', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('CO2', 'CO2-Neutral', 'CO2-Neutral', 'CO2-Neutral', 'MEDIUM', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('B', 'Bio', 'Bio', 'Bio', 'HIGHER', FALSE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden) VALUES ('MV', 'MensaVital', 'MensaVital', 'MensaVital', 'MEDIUM', FALSE);
-- missing: alcohol icon

-- Allergies
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Wz', 'Weizen', 'Weizen (Gluten)', '🌾', 'MEDIUM', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Ro', 'Roggen', 'Roggen (Gluten)', '', 'MEDIUM', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Ge', 'Gerste', 'Gerste (Gluten)', '', 'MEDIUM', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Hf', 'Hafer', 'Hafer (Gluten)', '', 'MEDIUM', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Kr', 'Krebstiere', 'Krebstiere', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Ei', 'Eier', 'Eier', '🥚', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Fi', 'Fisch', 'Fisch', '🐟', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Er', 'Erdnüsse', 'Erdnüsse', '🥜', 'MEDIUM', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('So', 'Soja', 'Soja', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Mi', 'Milch/Laktose', 'Milch/Laktose', '', 'MEDIUM', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Man', 'Mandeln', 'Mandeln', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Hs', 'Haselnüsse', 'Haselnüsse', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Wa', 'Walnüsse', 'Walnüsse', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Ka', 'Cashewnüsse', 'Cashewnüsse', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Pe', 'Pekanüsse', 'Pekanüsse', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Pa', 'Paranüsse', 'Paranüsse', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Pi', 'Pistazien', 'Pistazien', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Mac', 'Macadamianüsse', 'Macadamianüsse', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Sel', 'Sellerie', 'Sellerie', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Sen', 'Senf', 'Senf', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Ses', 'Sesam', 'Sesam', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Su', 'Schwefeloxid/Sulfite', 'Schwefeloxid/Sulfite', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('Lu', 'Lupinen', 'Lupinen', '', 'LOWER', FALSE, TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_hidden, is_allergy) VALUES ('We', 'Weichtiere', 'Weichtiere', '', 'LOWER', FALSE, TRUE);
