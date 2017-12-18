with import <nixpkgs> {};
stdenv.mkDerivation {
  name = "gluttony-shell";
  buildInputs = [
    influxdb
    ncurses
    go
    gocode
    go-bindata
    glide
    godef
    bison
  ];
}
