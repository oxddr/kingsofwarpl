PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE Events (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, city TEXT NOT NULL, series TEXT REFERENCES Series (name), tabletop_url TEXT UNIQUE ON CONFLICT FAIL, location TEXT, format INTEGER, date DATE, website TEXT);
INSERT INTO Events VALUES(1,'Inauguracyjny turniej polskiej ligii Kings of War, sezon 2022','Katowice','Liga 2022','https://tabletop.to/inauguracyjny-turniej-polskiej-ligii-kings-of-war-sezon-2022','Inny Wymiar, Katowice, ul. Bożogrobców 20',1995,'2022-01-22','https://www.facebook.com/events/617366889679424');
INSERT INTO Events VALUES(2,'II turniej Ligii King of War - Kraków','Kraków','Liga 2022','https://tabletop.to/ii-turniej-ligii-kings-of-war-2022-krakow','MDK Grunwaldzka, ul. Grunwaldzka 5 Kraków',1880,'2022-02-26','https://www.facebook.com/events/3008196419443442');
INSERT INTO Events VALUES(3,'Krawaty 13','Warszawa','Krawaty','https://tabletop.to/krawaty-13','Klub Adeptus Mechanicus al. Krakowska 114, Warszawa',1995,'2022-01-19','https://forum.kingsofwar.pl/showthread.php?tid=72');
INSERT INTO Events VALUES(4,'Krawaty 15','Warszawa','Krawaty','https://tabletop.to/krawaty-15','Klub Adeptus Mechanicus al. Krakowska 114, Warszawa',1995,'2022-03-05','https://forum.kingsofwar.pl/showthread.php?tid=78');
INSERT INTO Events VALUES(5,'Krawaty 14','Warszawa','Krawaty','https://tabletop.to/krawaty-14','Klub Adeptus Mechanicus al. Krakowska 114, Warszawa',1880,'2022-02-14','https://forum.kingsofwar.pl/showthread.php?tid=77');
INSERT INTO Events VALUES(6,'Gigantyczny problem z Zielonymi','Wrocław',NULL,NULL,'Bolter, Podwale 6, 50-043 Wrocław',2300,'2022-02-27','https://www.facebook.com/events/360306015602256');
INSERT INTO Events VALUES(7,'4 Turniej Mistrzostw Polski Kings of War 2021-BIAŁOSTOCKI DEBIUT','Białystok','Liga 2021','https://tabletop.to/4-turniej-mistrzostw-polski-kings-of-war-2021biaostocki-debiut','Klub Grota, Jana Henryka Dąbrowskiego 20, Białystok ',2300,'2021-09-18',NULL);
INSERT INTO Events VALUES(8,'II turniej cyklu MISTRZOSTW POLSKI Kings of War, 2021','Katowice','Liga 2021','https://tabletop.to/ii-turniej-cyklu-mistrzostw-polski-kings-of-war-2021','Inny Wymiar, Katowice, ul. Bożogrobców 20',2300,'2021-02-27',NULL);
INSERT INTO Events VALUES(9,'Pierwszy turniej ligi 2021 Warszawa','Warszawa','Liga 2021','https://tabletop.to/1-turniej-ligi-kings-of-war-2021-warszawa','Klub Adeptus Mechanicus al. Krakowska 114, Warszawa',2100,'2021-01-24',NULL);
INSERT INTO Events VALUES(10,'III Turniej Ligi Kings of War 2021','Warszawa','Liga 2021','https://tabletop.to/iii-turniej-ligi-kings-of-war-polska-2021','Klub Adeptus Mechanicus al. Krakowska 114, Warszawa',1900,'2021-04-18',NULL);
CREATE TABLE Series (name TEXT PRIMARY KEY NOT NULL);
INSERT INTO Series VALUES('Liga 2022');
INSERT INTO Series VALUES('Liga 2021');
INSERT INTO Series VALUES('Krawaty');
CREATE TABLE Results (player TEXT REFERENCES Players (tabletop_id), event INTEGER REFERENCES Events (id), faction TEXT REFERENCES Factions (faction), tp INTEGER, bonus_tp INTEGER, attrition_points INTEGER, UNIQUE (player, event) ON CONFLICT REPLACE);
INSERT INTO Results VALUES('misio',1,'The Trident Realm of Neritica',48,5,4790);
INSERT INTO Results VALUES('ddr',1,'Night-stalkers',53,0,4030);
INSERT INTO Results VALUES('lires',1,'Goblins',39,5,4510);
INSERT INTO Results VALUES('bart3',1,'Forces of the Abyss',37,5,3725);
INSERT INTO Results VALUES('nieumary-plankton',1,'Undead',38,0,3800);
INSERT INTO Results VALUES('stanisaw-lajon-urbaniak',1,'The Herd',37,0,3550);
INSERT INTO Results VALUES('adam-athard-cychner',1,'Elves',31,5,3320);
INSERT INTO Results VALUES('teo-tomrider',1,'Basilea',28,5,3455);
INSERT INTO Results VALUES('myter',1,'Night-stalkers',24,5,3205);
INSERT INTO Results VALUES('oldwin',1,'Forces of the Abyss',21,5,3505);
INSERT INTO Results VALUES('eledan',1,'Basilea',21,5,2525);
INSERT INTO Results VALUES('rafal-troll',1,'Dwarfs',19,5,3480);
INSERT INTO Results VALUES('grzegorz-gregx',1,'Dwarfs',19,5,2115);
INSERT INTO Results VALUES('gobos2',1,'Undead',5,0,1565);
INSERT INTO Results VALUES('makras',2,'Ogres',48,5,4515);
INSERT INTO Results VALUES('ddr',2,'Night-stalkers',48,5,3125);
INSERT INTO Results VALUES('lires',2,'Goblins',44,5,4580);
INSERT INTO Results VALUES('adam-athard-cychner',2,'Elves',43,5,3970);
INSERT INTO Results VALUES('bart3',2,'Forces of the Abyss',41,5,4950);
INSERT INTO Results VALUES('misio',2,'The Trident Realm of Neritica',38,5,4690);
INSERT INTO Results VALUES('nieumary-plankton',2,'Undead',38,5,3845);
INSERT INTO Results VALUES('stanisaw-lajon-urbaniak',2,'The Empire of Dust',38,5,3645);
INSERT INTO Results VALUES('bartoszwaw',2,'Kingdoms of Men',31,5,3095);
INSERT INTO Results VALUES('eledan',2,'Basilea',30,5,3520);
INSERT INTO Results VALUES('gobos2',2,'Northern Alliance',32,0,2565);
INSERT INTO Results VALUES('kacper-alien-szczytowski',2,'Undead',25,5,3050);
INSERT INTO Results VALUES('lukasz-jarochowski',2,'Undead',25,5,2450);
INSERT INTO Results VALUES('rafal-troll',2,'Dwarfs',22,5,2980);
INSERT INTO Results VALUES('maciejufo-szczytowski',2,'Abyssal Dwarfs',20,5,3010);
INSERT INTO Results VALUES('grzegorz-gregx',2,'Free Dwarfs',20,5,2835);
INSERT INTO Results VALUES('teo-tomrider',2,'Basilea',19,5,3450);
INSERT INTO Results VALUES('seba',2,'The Varangur',17,5,2470);
INSERT INTO Results VALUES('berdysz',2,'Kingdoms of Men',15,5,2385);
INSERT INTO Results VALUES('oldwin',2,'Abyssal Dwarfs',6,5,2210);
INSERT INTO Results VALUES('makras',9,'Ogres',48,5,4620);
INSERT INTO Results VALUES('misio',9,'The Trident Realm of Neritica',44,5,3490);
INSERT INTO Results VALUES('pawe-pawluczuk',9,'Forces of the Abyss',42,5,5590);
INSERT INTO Results VALUES('kazan2',9,'Undead',37,5,4545);
INSERT INTO Results VALUES('ddr',9,'Night-stalkers',36,5,3480);
INSERT INTO Results VALUES('bartoszwaw',9,'The League of Rhordia',34,5,2980);
INSERT INTO Results VALUES('adam-athard-cychner',9,'Elves',30,5,3065);
INSERT INTO Results VALUES('lires',9,'Elves',26,5,4150);
INSERT INTO Results VALUES('stanisaw-lajon-urbaniak',9,'The Empire of Dust',25,5,3415);
INSERT INTO Results VALUES('damian-fakapu-fiedoruk',9,'The Empire of Dust',25,0,3360);
INSERT INTO Results VALUES('kamil-wichan',9,'Dwarfs',19,5,3070);
INSERT INTO Results VALUES('maciejufo-szczytowski',9,'Abyssal Dwarfs',22,0,2160);
INSERT INTO Results VALUES('kacper-rasinski',9,'Forces of the Abyss',23,-5,3470);
INSERT INTO Results VALUES('nieumary-plankton',9,'Undead',9,0,2675);
INSERT INTO Results VALUES('lires',10,'Goblins',49,0,4330);
INSERT INTO Results VALUES('kazan2',10,'Undead',39,0,4775);
INSERT INTO Results VALUES('kamil-wichan',10,'Dwarfs',37,0,4460);
INSERT INTO Results VALUES('makras',10,'Ogres',36,0,4130);
INSERT INTO Results VALUES('lukasz-jarochowski',10,'Undead',33,0,1910);
INSERT INTO Results VALUES('ddr',10,'Night-stalkers',25,0,3325);
INSERT INTO Results VALUES('stanisaw-lajon-urbaniak',10,'The Empire of Dust',25,-5,2920);
INSERT INTO Results VALUES('kacper-alien-szczytowski',10,'Undead',24,-5,3760);
INSERT INTO Results VALUES('bartoszwaw',10,'Order of the Green Lady',17,0,2345);
INSERT INTO Results VALUES('maciejufo-szczytowski',10,'Abyssal Dwarfs',15,0,2765);
INSERT INTO Results VALUES('misio',8,'',47,5,4365);
INSERT INTO Results VALUES('makras',8,'',43,5,4185);
INSERT INTO Results VALUES('lires',8,'',36,5,4395);
INSERT INTO Results VALUES('ddr',8,'',33,5,4095);
INSERT INTO Results VALUES('rafal-troll',8,'',25,5,3795);
INSERT INTO Results VALUES('maciejufo-szczytowski',8,'',26,0,4765);
INSERT INTO Results VALUES('grzegorz-gregx',8,'',18,5,3770);
INSERT INTO Results VALUES('makras',7,'',50,0,5700);
INSERT INTO Results VALUES('lires',7,'',37,0,4790);
INSERT INTO Results VALUES('ddr',7,'',41,-5,5610);
INSERT INTO Results VALUES('bartoszwaw',7,'',35,0,4490);
INSERT INTO Results VALUES('lukasz-jarochowski',7,'',27,0,3280);
INSERT INTO Results VALUES('rafastrongestavenger-swierkot',7,'',31,-5,4335);
INSERT INTO Results VALUES('damian-fakapu-fiedoruk',7,'',26,0,3980);
INSERT INTO Results VALUES('kacper-alien-szczytowski',7,'',22,0,3815);
INSERT INTO Results VALUES('kacper-rasinski',7,'',21,-5,4175);
INSERT INTO Results VALUES('maciejufo-szczytowski',7,'',10,0,3079);
CREATE TABLE Factions (faction TEXT PRIMARY KEY UNIQUE NOT NULL) WITHOUT ROWID;
INSERT INTO Factions VALUES('Abyssal Dwarfs');
INSERT INTO Factions VALUES('Basilea');
INSERT INTO Factions VALUES('Dwarfs');
INSERT INTO Factions VALUES('Elves');
INSERT INTO Factions VALUES('Empire of Dust');
INSERT INTO Factions VALUES('Forces of Nature');
INSERT INTO Factions VALUES('Forces of the Abyss');
INSERT INTO Factions VALUES('Free Dwarfs');
INSERT INTO Factions VALUES('Goblins');
INSERT INTO Factions VALUES('Halflings');
INSERT INTO Factions VALUES('Kingdoms of Men');
INSERT INTO Factions VALUES('League of Rhordia');
INSERT INTO Factions VALUES('Night-stalkers');
INSERT INTO Factions VALUES('Northern Alliance');
INSERT INTO Factions VALUES('Ogres');
INSERT INTO Factions VALUES('Orcs');
INSERT INTO Factions VALUES('Order of the Brothermark');
INSERT INTO Factions VALUES('Order of the Green Lady');
INSERT INTO Factions VALUES('Ratkin');
INSERT INTO Factions VALUES('Ratkin Slaves');
INSERT INTO Factions VALUES('Riftforged Orcs');
INSERT INTO Factions VALUES('Salamanders');
INSERT INTO Factions VALUES('Sylvan Kin');
INSERT INTO Factions VALUES('The Herd');
INSERT INTO Factions VALUES('Trident Realm of Neritica');
INSERT INTO Factions VALUES('Twilight Kin');
INSERT INTO Factions VALUES('Undead');
INSERT INTO Factions VALUES('Varangur');
CREATE TABLE Players (name TEXT, tabletop_id TEXT PRIMARY KEY ON CONFLICT REPLACE);
INSERT INTO Players VALUES('Bart','bart3');
INSERT INTO Players VALUES('Eledan','eledan');
INSERT INTO Players VALUES('Gobos','gobos2');
INSERT INTO Players VALUES('Teo TomRider','teo-tomrider');
INSERT INTO Players VALUES('Seba','seba');
INSERT INTO Players VALUES('Berdysz','berdysz');
INSERT INTO Players VALUES('Oldwin','oldwin');
INSERT INTO Players VALUES('Myter','myter');
INSERT INTO Players VALUES('Paweł Pawluczuk','pawe-pawluczuk');
INSERT INTO Players VALUES('Adam "Athard" Cychner','adam-athard-cychner');
INSERT INTO Players VALUES('Nieumarły Pla(nk)ton','nieumary-plankton');
INSERT INTO Players VALUES('Kazan','kazan2');
INSERT INTO Players VALUES('Kamil Wichan','kamil-wichan');
INSERT INTO Players VALUES('Stanisław "Lajon" Urbaniak','stanisaw-lajon-urbaniak');
INSERT INTO Players VALUES('MiSiO','misio');
INSERT INTO Players VALUES('Rafal ''Troll''','rafal-troll');
INSERT INTO Players VALUES('Grzegorz Gregx','grzegorz-gregx');
INSERT INTO Players VALUES('Makras','makras');
INSERT INTO Players VALUES('Lires','lires');
INSERT INTO Players VALUES('ddr','ddr');
INSERT INTO Players VALUES('BartoszWaw','bartoszwaw');
INSERT INTO Players VALUES('Lukasz Jarochowski','lukasz-jarochowski');
INSERT INTO Players VALUES('Rafał"StrongestAvenger" Świerkot','rafastrongestavenger-swierkot');
INSERT INTO Players VALUES('Damian "Fakapu" Fiedoruk','damian-fakapu-fiedoruk');
INSERT INTO Players VALUES('Kacper "Alien" Szczytowski','kacper-alien-szczytowski');
INSERT INTO Players VALUES('Kacper Rasiński','kacper-rasinski');
INSERT INTO Players VALUES('Maciej"UFO" Szczytowski','maciejufo-szczytowski');
DELETE FROM sqlite_sequence;
INSERT INTO sqlite_sequence VALUES('Events',10);
COMMIT;