{
  description = "Research Loop - Agent OS for scientific researchers";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go_1_26
            gopls
            gotools
            go-tools
            golangci-lint
          ];

          shellHook = ''
            echo "Research Loop development environment"
            echo "Go version: $(go version)"
          '';
        };
      }
    );
}
