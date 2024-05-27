{
  description = "protoc-gen-event is a tool that allows to generate events structs from protobuf interfaces";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/master";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    let
      mkPackage = pkgs: pkgs.callPackage ./default.nix {
        #version = self.shortRev or "dirty";
        date = self.lastModifiedDate;
        commit = self.shortRev or "dirty";
      };
    in
    (flake-utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
    in {
        packages.protoc-gen-event = mkPackage pkgs;
        defaultPackage = self.packages."${system}".protoc-gen-event;
        apps.protoc-gen-event = flake-utils.lib.mkApp {
          drv = self.packages."${system}".protoc-gen-event;
        };
        apps.default = self.apps."${system}".protoc-gen-event;
        # checks.default = self.packages."${system}".protoc-gen-event.overrideAttrs (prev: {
        #   doCheck = true;
        #   nativeBuildInputs = prev.nativeBuildInputs ++ (with pkgs; [python3]);
        # });
      })
    ) // {
      overlays.default = final: prev: {
        protoc-gen-event = mkPackage final;
      };
    }
  ;
}
