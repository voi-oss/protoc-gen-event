{ buildGoModule, config, lib, pkgs, installShellFiles, date, commit }:

buildGoModule rec {
  pname = "protoc-gen-event";

  src = lib.cleanSource ./.;
  ldflags = [
    "-s"
    "-w"
  ];

  meta = with lib; {
    description = "protoc-gen-event is a tool that allows to generate events structs from protobuf interfaces";
  };

  doCheck = false; # it takes ages to run the tests
}
