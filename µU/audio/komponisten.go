package audio

import
  "µU/text"
var
  k = []string {"Albinoni, Tomaso",
                "Bach, Johann Sebastian",
                "Bach, Philip Emanuel",
                "Bartok, Bela",
                "Beethoven, Ludwig van",
                "Berg, Alban",
                "Berlioz, Hector",
                "Bizet, Georges",
                "Blacher, Boris",
                "Boccherini, Luigi",
                "Borodin, Alexander",
                "Boulez, Pierre",
                "Brahms, Johannes",
                "Britten, Benjamin",
                "Bruckner, Anton",
                "Buxtehude, Dietrich",
                "Caldara, Antonio",
                "Chatschaturian, Aram",
                "Cherubini, Luigi",
                "Chopin, Frederic",
                "Corelli, Arcangelo",
                "Couperin, Francois",
                "Debussy, Claude",
                "Donizetti, Gaetano",
                "Dvorak, Antonin",
                "Falla, Manuel de",
                "Fortner, Wolfgang",
                "Franck, Cesar",
                "Frescobaldi, Girolamo",
                "Furtwängler, Wilhelm",
                "Geminiani, Francesco",
                "Glinka, Michael",
                "Gluck, Christoph Willibald",
                "Grieg, Edward",
                "Händel, Georg Friedrich",
                "Haydn, Joseph",
                "Hindemith, Paul",
                "Honegger, Arthur",
                "Hummel, Johann Nepomuk",
                "Janacek, Leos",
                "Lalo, Edouard",
                "Liszt, Franz",
                "Locatelli, Pietro",
                "Lortzing, Albert",
                "Lully, Jean Baptiste",
                "Mahler, Gustav",
                "Mendelssohn-Bartholdy, Felix",
                "Monteverdi, Claudio",
                "Mozart, Wolfgang Amadeus",
                "Mussorgski, Modest",
                "Orff, Carl",
                "Paganini, Niccolo",
                "Palestrina, Claudio da",
                "Pergolesi, Giovanni Battista",
                "Pfitzner, Hans",
                "Prokofieff, Serge",
                "Purcell, Henry",
                "Rachmaninow, Sergej",
                "Rameau, Jean-Philippe",
                "Ravel, Maurice",
                "Reger, Max",
                "Rimskij-Korssakow, Nikolai",
                "Rossini, Gioacchino",
                "Saint-Saens, Camille",
                "Scarlatti, Alessandro",
                "Scarlatti, Domenico",
                "Schönberg, Arnold",
                "Schostakowitsch, Dimitri",
                "Schubert, Franz",
                "Schumann, Robert",
                "Schütz, Heinrich",
                "Scriabin, Alexander",
                "Sibelius, Jean",
                "Smetana, Friedrich",
                "Spohr, Ludwig",
                "Strauss, Johann",
                "Strauß, Richard",
                "Strawinsky, Igor",
                "Telemann, Georg Philipp",
                "Tschaikowskij, Peter",
                "Verdi, Giuseppe",
                "Vivaldi, Antonio",
                "Wagner, Richard",
                "Weber, Carl Maria von",
                "Webern, Anton von"}
var
  komponist []text.Text

func init() {
  n := len(k)
  komponist = make([]text.Text, n)
  for i := 0; i < n; i++ {
    komponist[i] = text.New (len0)
    komponist[i].Defined (k[i])
  }
}
