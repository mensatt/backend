-- dishes
INSERT INTO dish (id, name) VALUES ('4328c2c1-242d-4bb3-9797-05d3ed2557bb', 'MSC - Alaska-Seelachsfilet gebacken mit Zitrone, Remouladensoße und Kartoffeln');
INSERT INTO dish (id, name) VALUES ('4328c2c1-242d-4bb3-9797-05d3ed2557ba', 'Pommes 🍟');

-- occurrences
INSERT INTO occurrence (id, dish, date, price_student, price_staff, price_guest) VALUES ('8e492a06-fda2-4a24-b899-c1ec169a509b', '4328c2c1-242d-4bb3-9797-05d3ed2557bb', '2022-02-28', 350, 450, 700);
INSERT INTO occurrence (id, dish, date, price_student, price_staff, price_guest) VALUES ('8e492a06-fda2-4a24-b899-c1ec169a509c', '4328c2c1-242d-4bb3-9797-05d3ed2557ba', '2022-02-28', 200, 300, 400);
