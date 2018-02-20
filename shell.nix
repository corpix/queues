with import <nixpkgs> {};
stdenv.mkDerivation {
  name = "nix-cage-shell";
  buildInputs = [
    influxdb
    go
    gocode
    go-bindata
    glide
    godef
  ];
  shellHook = ''
    export GOPATH=~/projects
  '';
}
