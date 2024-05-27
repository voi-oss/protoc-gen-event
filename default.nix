{ buildGoModule, config, lib, pkgs, installShellFiles, date, commit }:

buildGoModule rec {
  pname = "protoc-gen-event";
  version = "1.0.0";
  vendorHash = "sha256-nPf1IQG4aWlbVQMe9GMOGo/QZvmrz2TF7eMINqvzQOc=";

  src = lib.cleanSource ./.;
  ldflags = [
    "-s"
    "-w"
  ];

  meta = with lib; {
    description = "protoc-gen-event is a tool that allows to generate events structs from protobuf interfaces";
  };
  doCheck = false;
}
