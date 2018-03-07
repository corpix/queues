with import <nixpkgs> {};
stdenv.mkDerivation {
  name = "nix-cage-shell";
  buildInputs = [
    apacheKafka
    nsq
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
