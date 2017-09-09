package col

// (c) Christian Maurer   v. 140101 - license see murus.go

var ( // RAL-Farben
  Grünbeige =        colour3 (214, 199, 148)
  Beige =            colour3 (217, 186, 140)
  Sandgelb =         colour3 (214, 176, 117)
  Signalgelb =       colour3 (252, 163,  41)
  Goldgelb =         colour3 (227, 150,  36)
  Honiggelb =        colour3 (201, 135,  33)
  Maisgelb =         colour3 (224, 130,  31)
  Narzissengelb =    colour3 (227, 122,  31)
  Braunbeige =       colour3 (173, 122,  79)
  Zitronengelb =     colour3 (227, 184,  56)
  Perlweiß =         colour3 (255, 245, 227)
  Elfenbein =        colour3 (240, 214, 171)
  Hellelfenbein =    colour3 (252, 235, 204)
  Schwefelgelb =     colour3 (255, 245,  66)
  Safrangelb =       colour3 (255, 171,  89)
  Zinkgelb =         colour3 (255, 214,  77)
  Graubeige =        colour3 (163, 140, 122)
  Olivgelb =         colour3 (156, 143,  97)
  Rapsgelb =         colour3 (252, 189,  31)
  Verkehrsgelb =     colour3 (252, 184,  33)
  Ockergelb =        colour3 (181, 140,  79)
  Leuchtgelb =       colour3 (255, 255,  10)
  Currygelb =        colour3 (153, 117,  33)
  Melonengelb =      colour3 (255, 140,  26)
  Ginstergelb =      colour3 (227, 163,  41)
  Dahliengelb =      colour3 (255, 148,  54)
  Pastelgelb =       colour3 (247, 153,  92)

  Gelborange =       colour3 (224,  94,  31)
  Rotorange =        colour3 (186,  46,  33)
  Blutorange =       colour3 (204,  36,  28)
  Pastellorange =    colour3 (255,  99,  54)
  Reinorange =       colour3 (242,  59,  28)
  Leuchtorange =     colour3 (252,  28,  20)
  Leuchthellorange = colour3 (255, 117,  33)
  Hellrotorange =    colour3 (250,  79,  41)
  Verkehrsorange =   colour3 (235,  59,  28)
  Signalorange =     colour3 (212,  69,  41)
  Tieforange =       colour3 (237,  92,  41)
  Lachsorange =      colour3 (222,  82,  71)

  Feuerrot =         colour3 (171,  31,  28)
  Signalrot =        colour3 (163,  23,  26)
  Karminrot =        colour3 (163,  26,  26)
  Rubinrot =         colour3 (138,  18,  20)
  Purpurrot =        colour3 (105,  15,  20)
  Weinrot =          colour3 ( 79,  18,  26)
  Schwarzrot =       colour3 ( 46,  18,  26)
  Oxidrot =          colour3 ( 94,  33,  33)
  Braunrot =         colour3 (120,  20,  23)
  Beigerot =         colour3 (204, 130, 115)
  Tomatenrot =       colour3 (150,  31,  28)
  Altrosa =          colour3 (217, 102, 117)
  Hellrosa =         colour3 (232, 156, 181)
  Korallenrot =      colour3 (166,  36,  38)
  Rose =             colour3 (209,  54,  84)
  Erdbeerrot =       colour3 (207,  41,  66)
  Verkehrsrot =      colour3 (199,  23,  18)
  Lachsrot =         colour3 (217,  89,  79)
  Leuchtrot =        colour3 (252,  10,  28)
  Leuchthellrot =    colour3 (252,  20,  20)
  Himbeerrot =       colour3 (181,  18,  51)
  Orientrot =        colour3 (166,  28,  46)

  Rotlila =          colour3 (130,  64, 128)
  Rotmagenta =       colour3 (143,  38,  64)
  Erikamagenta =     colour3 (201,  56, 140)
  Bordeauxmagenta =  colour3 ( 92,   8,  43)
  Blaulila =         colour3 ( 99,  61, 156)
  Verkehrspurpur =   colour3 (145,  15, 102)
  Purpurmagenta =    colour3 ( 56,  10,  46)
  Signalviolett =    colour3 (125,  31, 122)
  Pastelviolett =    colour3 (158, 115, 148)
  Telemagenta =      colour3 (191,  23, 115)

  Violettblau =      colour3 ( 23,  51, 107)
  Grünblau =         colour3 ( 10,  51,  84)
  Ultramarinblau =   colour3 (  0,  15, 117)
  Saphirblau =       colour3 (  0,  23,  69)
  Schwarzblau =      colour3 (  3,  13,  31)
  Signalblau =       colour3 (  0,  46, 122)
  Brillantblau =     colour3 ( 38,  79, 135)
  Graublau =         colour3 ( 26,  41,  56)
  Azurblau =         colour3 ( 23,  69, 112)
  Enzianblau =       colour3 (  0,  43, 112)
  Stahlblau =        colour3 (  3,  20,  46)
  Lichtblau =        colour3 ( 41, 115, 184)
  Kobaltblau =       colour3 (  0,  18,  69)
  Taubenblau =       colour3 ( 77, 105, 153)
  Himmelblau =       colour3 ( 23,  97, 171)
  Verkehrsblau =     colour3 (  0,  59, 128)
  Türkisblau =       colour3 ( 56, 148, 130)
  Capriblau =        colour3 ( 10,  66, 120)
  Ozeanblau =        colour3 (  5,  51,  51)
  Wasserblau =       colour3 ( 26, 122,  99)
  Nachtblau =        colour3 (  0,   8,  79)
  Fernblau =         colour3 ( 46,  82, 143)
  Pastellblau =      colour3 ( 87, 140, 181)

  Patinagrün =       colour3 ( 51, 120,  84)
  Smaragdgrün =      colour3 ( 38, 102,  41)
  Laubgrün =         colour3 ( 38,  87,  33)
  Olivgrün =         colour3 ( 61,  69,  46)
  Blaugrün =         colour3 ( 13,  59,  46)
  Moosgrün =         colour3 ( 10,  56,  31)
  Grauoliv =         colour3 ( 41,  43,  36)
  Flaschengrün =     colour3 ( 28,  38,  23)
  Braungrün =        colour3 ( 33,  33,  26)
  Tannengrün =       colour3 ( 23,  41,  28)
  Grasgrün =         colour3 ( 54, 105,  38)
  Resedagrün =       colour3 ( 94, 125,  79)
  Schwarzgrün =      colour3 ( 31,  46,  43)
  Schilfgrün =       colour3 (117, 115,  79)
  Gelboliv =         colour3 ( 51,  48,  38)
  Schwarzoliv =      colour3 ( 41,  43,  38)
  Cyangrün =         colour3 ( 15, 112,  51)
  Maigrün =          colour3 ( 64, 130,  54)
  Gelbgrün =         colour3 ( 79, 168,  51)
  Weißgrün =         colour3 (191, 227, 186)
  Chromoxidgrün =    colour3 ( 38,  56,  41)
  Blassgrün =        colour3 (133, 166, 122)
  Braunoliv =        colour3 ( 43,  38,  28)
  Verkehrsgrün =     colour3 ( 36, 145,  64)
  Farngrün =         colour3 ( 74, 110,  51)
  Opalgrün =         colour3 ( 10,  92,  51)
  Lichtgrün =        colour3 (125, 204, 189)
  Kieferngrün =      colour3 ( 38,  74,  51)
  Minzgrün =         colour3 ( 18, 120,  38)
  Signalgrün =       colour3 ( 41, 138,  64)
  Minttürkis =       colour3 ( 66, 140, 120)
  Pasteltürkis =     colour3 (125, 189, 181)

  Fehgrau =          colour3 (115, 133, 145)
  Silbergrau =       colour3 (135, 148, 166)
  Olivgrau =         colour3 (122, 117,  97)
  Moosgrau =         colour3 (112, 112,  97)
  Signalgrau =       colour3 (156, 156, 166)
  Mausgrau =         colour3 ( 97, 105, 105)
  Beigegrau =        colour3 (107,  97,  87)
  Khakigrau =        colour3 (105,  84,  56)
  Grüngrau =         colour3 ( 77,  82,  74)
  Zeltgrau =         colour3 ( 74,  79,  74)
  Eisengrau =        colour3 ( 64,  74,  84)
  Basaltgrau =       colour3 ( 74,  84,  89)
  Braungrau =        colour3 ( 71,  66,  56)
  Schiefergrau =     colour3 ( 61,  66,  82)
  Anthrazitgrau =    colour3 ( 38,  46,  56)
  Schwarzgrau =      colour3 ( 26,  33,  41)
  Umbragrau =        colour3 ( 61,  61,  59)
  Betongrau =        colour3 (122, 125, 117)
  Graphitgrau =      colour3 ( 48,  56,  69)
  Granitgrau =       colour3 ( 38,  51,  56)
  Steingrau =        colour3 (145, 143, 135)
  Blaugrau =         colour3 ( 77,  92, 107)
  Kieselgrau =       colour3 (189, 186, 171)
  Zementgrau =       colour3 (122, 130, 117)
  Gelbgrau =         colour3 (143, 135, 112)
  Lichtgrau =        colour3 (212, 217, 219)
  Platingrau =       colour3 (158, 150, 156)
  Staubgrau =        colour3 (122, 125, 128)
  Achatgrau =        colour3 (186, 189, 186)
  Quarzgrau =        colour3 ( 97,  94,  89)
  Fenstergrau =      colour3 (158, 163, 176)
  VerkehrsgrauA =    colour3 (143, 150, 153)
  VerkehrsgrauB =    colour3 ( 64,  69,  69)
  Seidengrau =       colour3 (194, 191, 184)
  Telegrau1 =        colour3 (143, 148, 158)
  Telegrau2 =        colour3 (120, 130, 140)
  Telegrau4 =        colour3 (217, 214, 219)

  Grünbraun =        colour3 (125,  92,  56)
  Ockerbraun =       colour3 (145,  82,  46)
  Signalbraun =      colour3 (110,  59,  48)
  Lehmbraun =        colour3 (115,  59,  36)
  Kupferbraun =      colour3 (133,  56,  43)
  Rehbraun =         colour3 ( 94,  51,  31)
  Olivbraun =        colour3 ( 99,  61,  36)
  Nussbraun =        colour3 ( 71,  38,  28)
  Rotbraun =         colour3 ( 84,  31,  31)
  Sepiabraun =       colour3 ( 56,  38,  28)
  Kastanienbraun =   colour3 ( 77,  31,  28)
  Mahagonibraun =    colour3 ( 61,  31,  28)
  Schokoladenbraun = colour3 ( 46,  28,  28)
  Graubraun =        colour3 ( 43,  38,  41)
  Schwarzbraun =     colour3 ( 13,   8,  13)
  Orangebraun =      colour3 (156,  69,  41)
  Beigebraun =       colour3 (110,  64,  48)
  Blassbraun =       colour3 (102,  74,  61)
  Terrabraun =       colour3 ( 64,  46,  33)

  Cremeweiß =        colour3 (255, 252, 240)
  Grauweiß =         colour3 (240, 237, 230)
  Signalweiß =       colour3 (255, 255, 255)
  Signalschwarz =    colour3 ( 28,  28,  33)
  Tiefschwarz =      colour3 (  3,   5,  10)
  Aluminiumweiß =    colour3 (166, 171, 181)
  Aluminiumgrau =    colour3 (125, 122, 120)
  Reinweiß =         colour3 (250, 255, 255)
  Graphitschwarz =   colour3 ( 13,  18,  26)
  Verkehrweiß =      colour3 (252, 255, 255)
  Verkehrschwarz =   colour3 ( 20,  23,  28)
  Papyrusweiß =      colour3 (219, 227, 222)
)
