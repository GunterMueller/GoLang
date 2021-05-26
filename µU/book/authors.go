package book

// (c) Christian Maurer   v. 210512 - license see µU.go

import
  "µU/text"

var
  a = []string {"Andersch, Alfred              ",
                "Bachmann, Ingeborg            ",
                "Böll, Heinrich                ",
                "Brecht, Berthold              ",
                "Büchner, Georg                ",
                "Camilleri, Andrea             ",
                "Carofiglio, Gianrico          ",
                "Chamisso, Adelbert von        ",
                "Claudius, Matthias            ",
                "Davis, Lindsey                ",
                "De Cataldo, Giancarlo         ",
                "Dibdin, Michael               ",
                "Donna Leon                    ",
                "Droste-Hülshoff, Annette von  ",
                "Dürrenmatt, Friedrich         ",
                "Eco, Umberto                  ",
                "Eichendorff, Joseph von       ",
                "Faletti, Giorgio              ",
                "Feuchtwanger, Lion            ",
                "Filasto, Nino                 ",
                "Fontane, Theodor              ",
                "Frisch, Max                   ",
                "Fruttero, Carlo               ",
                "Gellert, Christian Fürchtegott",
                "Giuttari, Michele             ",
                "Goethe, Johann Wolfgang       ",
                "Grabbe, Christian Dietrich    ",
                "Grass, Günter                 ",
                "Grimaldi, Laura               ",
                "Grimmelshausen, Hans Jacob Ch.",
                "Harris, Robert                ",
                "Hauptmann, Gerhart            ",
                "Hausmann, Manfred             ",
                "Hoffmann, E. T. A.            ",
                "Hölderlin, Friedrich          ",
                "Isari, Andrea (Richter)       ",
                "Kafka, Franz                  ",
                "Kästner, Erich                ",
                "Keller, Gottfried             ",
                "Kleist, Heinrich von          ",
                "Lama, Diana Fiametta          ",
                "Lasker-Schüler, Else          ",
                "Lem, Stanislaw                ",
                "Lenz, Jakob Michael Reinhold  ",
                "Lenz, Siegfried               ",
                "Lessing, Gotthold Ephraim     ",
                "Lucarelli, Carlo              ",
                "Lucentini, Franco             ",
                "Macchiavelli, Loriano         ",
                "Malvaldi                      ",
                "Mann, Heinrich                ",
                "Mann, Thomas                  ",
                "Mayall, Felicitas (Veit, B.)  ",
                "McCullough, Coleen            ",
                "Mistretta, Roberto            ",
                "Mongardi, Fabio               ",
                "Montanari, Danila Comastri    ",
                "Morgenstern, Christian        ",
                "Mörike, Eduard                ",
                "Paglieri, Claudio             ",
                "Paul, Jean                    ",
                "Piazzese, Santo               ",
                "Pittorru, Fabio               ",
                "Prior, Lily                   ",
                "Remin, Nicolas                ",
                "Rilke, Rainer Maria           ",
                "Ringelnatz, Joachim           ",
                "Roberts, John Maddox          ",
                "Saylor, Steven                ",
                "Scarrow, Simon                ",
                "Scerbanenco, Giorgio          ",
                "Schiller, Friedrich           ",
                "Simi, Giampaolo               ",
                "Soria, Piero                  ",
                "Stifter, Adalbert             ",
                "Storm, Theodor                ",
                "Stöver, Hans Dieter           ",
                "Tucholsky, Kurt               ",
                "Varesi, Valerio               ",
                "Verasani, Giorgio             ",
                "Vichi, Marco                  ",
                "Walser, Martin                ",
                "Wedekind, Frank               ",
                "Weiss, Peter                  ",
                "Werfel, Franz                 ",
                "Wishart, David                ",
                "Zilahy, Mirko                 ",
                "Zweig, Stefan                 "}
var
  author []text.Text

func init() {
  n := len(a)
  author = make([]text.Text, n)
  for i := 0; i < n; i++ {
    author[i] = text.New (len0)
    author[i].Defined (a[i])
  }
}
