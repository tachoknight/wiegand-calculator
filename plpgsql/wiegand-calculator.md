# wiegand-calculator
An encode/decode calculator for Wiegand-based cards and tags, written in [PL/pgSQL](https://en.wikipedia.org/wiki/PL/pgSQL).

This is basically the same version as the Go one, but ported to plpgsql so the encoding/decodibng can be done as entries into a Postgres database. The functions can be used like:

`select converttowiegand(10978235); -- convert the tag to the wiegand value`

`select convertfromwiegand(16733723); -- convert the weigand value (the converted number from above) back to the original tag number`


To make sure the tag is converted correctly, you can double up the functions like so:

`select convertfromwiegand(converttowiegand(10978235));`