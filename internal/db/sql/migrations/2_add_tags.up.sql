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
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('S', 'Schwein', 'Schweinefleisch', '🐷', 'HIGH');
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('R', 'Rind', 'Rindfleisch', '🐮', 'HIGH');
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('G', 'Geflügel', 'Geflügel', '🐔', 'HIGH');
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('L', 'Lamm', 'Lamm', '🐑', 'HIGH');
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('W', 'Wild', 'Wildfleisch', '🐈', 'HIGH');
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('F', 'Fisch', 'Fisch', '🐟', 'HIGH');
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('V', 'Vegetarisch', 'Vegetarisch', '🥕', 'HIGH');
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('Veg', 'Vegan', 'Vegan', '🌱', 'HIGH');
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('MSC', 'MSC Fish', 'MSC Fish', '🐟', 'HIGH');
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Gf', 'Glutenfrei', 'Glutenfrei', 'Glutenfrei', 'MEDIUM', TRUE);
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('CO2', 'CO2-Neutral', 'CO2-Neutral', 'CO2-Neutral', 'MEDIUM');
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('B', 'Bio', 'Bio', 'Bio', 'HIGH');
INSERT INTO tag (key, name, description, short_name, priority) VALUES ('MV', 'MensaVital', 'MensaVital', 'MensaVital', 'MEDIUM');
-- missing: alcohol icon

-- Allergies
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Wz', 'Weizen', 'Weizen (Gluten)', '🌾', 'MEDIUM', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Ro', 'Roggen', 'Roggen (Gluten)', '', 'MEDIUM', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Ge', 'Gerste', 'Gerste (Gluten)', '', 'MEDIUM', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Hafer', 'Hafer', 'Hafer (Gluten)', '', 'MEDIUM', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Kr', 'Krebstiere', 'Krebstiere', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Ei', 'Eier', 'Eier', '🥚', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Fi', 'Fisch', 'Fisch', '🐟', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Er', 'Erdnüsse', 'Erdnüsse', '🥜', 'MEDIUM', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('So', 'Soja', 'Soja', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Mi', 'Milch/Laktose', 'Milch/Laktose', '', 'MEDIUM', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Man', 'Mandeln', 'Mandeln', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Hs', 'Haselnüsse', 'Haselnüsse', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Wa', 'Walnüsse', 'Walnüsse', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Ka', 'Cashewnüsse', 'Cashewnüsse', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Pe', 'Pekanüsse', 'Pekanüsse', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Pa', 'Paranüsse', 'Paranüsse', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Pi', 'Pistazien', 'Pistazien', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Mac', 'Macadamianüsse', 'Macadamianüsse', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Sel', 'Sellerie', 'Sellerie', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Sen', 'Senf', 'Senf', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Ses', 'Sesam', 'Sesam', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Su', 'Schwefeloxid/Sulfite', 'Schwefeloxid/Sulfite', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('Lu', 'Lupinen', 'Lupinen', '', 'LOW', TRUE);
INSERT INTO tag (key, name, description, short_name, priority, is_allergy) VALUES ('We', 'Weichtiere', 'Weichtiere', '', 'LOW', TRUE);
